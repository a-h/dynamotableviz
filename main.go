package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	_ "embed"

	"github.com/a-h/dynamotableviz/value"
	"github.com/a-h/templ"
)

var pkFlag = flag.String("pk", "pk", "Name of the partition key attribute")
var skFlag = flag.String("sk", "", "Name of the sort key attribute")
var attrsFlag = flag.String("attrs", "gsi1,gsi2,gsi3,ttl", "Defines named attributes, which are then shown as a column")
var fileFlag = flag.String("file", "", "Load the data from the file instead of stdin.")

func main() {
	flag.Parse()

	// Check if data is being piped.
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("failed to stat stdin: %v", err)
	}
	var isPipe bool
	if (fi.Mode() & os.ModeNamedPipe) != 0 {
		isPipe = true
	}

	// If data is being piped, and a filename has been set, that's an error.
	if isPipe && *fileFlag != "" {
		fmt.Println("Cannot receive piped data when the file argument is set. Either set the file argument, or pipe data.")
		flag.Usage()
		os.Exit(1)
	}
	if !isPipe && *fileFlag == "" {
		flag.Usage()
		return
	}

	// Read the input rows in key/value pair format, see example.txt
	var keyValueText []byte
	if isPipe {
		keyValueText, err = io.ReadAll(os.Stdin)
	} else {
		keyValueText, err = os.ReadFile(*fileFlag)
	}
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	// Parse the rows.
	rows, err := value.ParseAll(string(keyValueText))
	if err != nil {
		log.Fatalf("failed to parse data: %v", err)
	}

	// Construct the view model for the table output.
	var d Data
	allAttributes := strings.Split(*attrsFlag, ",")
	if len(allAttributes) == 0 {
		log.Fatalf("attributes flag set with empty value, possibly not formatted properly?")
	}
	d.PK = *pkFlag
	d.SK = *skFlag
	// Only render named attributes that include some data.
	d.Attributes = getUsedAttributes(allAttributes, rows)
	// Configure maximum colspan values to ensure that the table has a consistent number of
	// columns rendered in each row.
	d.MaxColCount = getMaxColCount(rows)

	// pkToIndexMap maps the partition key value to the index in the list of
	// partitions. Each row is added to a partition.
	pkToIndexMap := map[string]int{}
	for rowIndex, r := range rows {
		var partitionRow Row
		partitionRow.Attributes = rows[rowIndex]
		for _, v := range r {
			if v.Key == d.PK {
				partitionRow.PK = v.Value
			}
			if v.Key == d.SK {
				partitionRow.SK = v.Value
			}
		}
		if partitionRow.PK == "" {
			log.Fatalf("row %d doesn't have a pk value (attribute named %q)", rowIndex, d.PK)
		}
		pkIndex, partitionAlreadyExists := pkToIndexMap[partitionRow.PK]
		if !partitionAlreadyExists {
			// Add a new partition to the slice, and keep track of its index in the map.
			d.Partitions = append(d.Partitions, Partition{PK: partitionRow.PK})
			pkToIndexMap[partitionRow.PK] = len(d.Partitions) - 1
			pkIndex = len(d.Partitions) - 1
		}
		// Add the row to the corresponding partition.
		d.Partitions[pkIndex].Rows = append(d.Partitions[pkIndex].Rows, partitionRow)
	}

	// Sort the partitions with the sort key.
	for _, p := range d.Partitions {
		sort.Slice(p.Rows, func(i, j int) bool {
			return strings.Compare(p.Rows[i].PK, p.Rows[j].PK) < 0
		})
	}

	// Render the HTML to stdout.
	table(d).Render(context.Background(), os.Stdout)
}

func getUsedAttributes(attrs []string, rows [][]value.Value) (filtered []string) {
	attributeToUsageMap := map[string]bool{}
	for _, r := range rows {
		for _, v := range r {
			attributeToUsageMap[v.Key] = true
		}
	}
	for _, attr := range attrs {
		if _, used := attributeToUsageMap[attr]; !used {
			continue
		}
		filtered = append(filtered, attr)
	}
	return filtered
}

func getMaxColCount(rows [][]value.Value) (count int) {
	for _, r := range rows {
		if len(r) > count {
			count = len(r)
		}
	}
	return
}

type Data struct {
	// PK is the name of the partition key attribute.
	PK string
	// SK is the name of the sort key attribute.
	SK string
	// Attributes are the attributes of a DynamoDB table that will be displayed with
	// a complete column, e.g. GSIs, and ttl attributes.
	Attributes []string
	// Partitions is the data in each partition of a DynamoDB table.
	Partitions []Partition
	// MaxColCount is the number of attributes in a single row.
	MaxColCount int
}

// IsNamedAttribute is used when drawing out the output table. Named attributes are
// displayed first (on the left), with the remaining attributes of a row displayed
// to the right. Since the named attributes are written out separetely, when we
// write out the rest of the attributes, we want to skip the ones we've already
// written.
func (d Data) IsNamedAttribute(key string) bool {
	if key == d.PK || key == d.SK {
		return true
	}
	for _, attr := range d.Attributes {
		if attr == key {
			return true
		}
	}
	return false
}

// GetAttributeCount gets the count of all non-key, non-named attributes in the row.
func (d Data) GetAttributeCount(row Row) (count int) {
	for _, v := range row.Attributes {
		if d.IsNamedAttribute(v.Value) {
			continue
		}
		count++
	}
	return count
}

type Partition struct {
	// PK is the value of the partition key.
	PK string
	// Rows are the named and un-named attributes.
	Rows []Row
}

type Row struct {
	PK string
	SK string
	// Attributes are the values, both named and unnamed.
	Attributes []value.Value
}

func getValueOrEmptyString(key string, r Row) string {
	for _, v := range r.Attributes {
		if v.Key == key {
			return v.Value
		}
	}
	return ""
}

func getRowClass(partitionIndex int) templ.CSSClasses {
	if partitionIndex%2 == 0 {
		return templ.CSSClasses{templ.ConstantCSSClass("jsontable-even")}
	}
	return templ.CSSClasses{templ.ConstantCSSClass("jsontable-odd")}
}
