package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/a-h/dynamotableviz/value"
	"github.com/a-h/templ"
)

var pkFlag = flag.String("pk", "pk", "Name of the partition key attribute")
var skFlag = flag.String("sk", "", "Name of the sort key attribute")
var attrs = flag.String("attrs", "gsi1,gsi2,gsi3,ttl", "Defines named attributes, which are then shown as a column")

func main() {
	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	rows, err := value.ParseAll(string(data))
	if err != nil {
		log.Fatalf("failed to parse data: %v", err)
	}

	var d Data
	allAttributes := strings.Split(*attrs, ",")
	if len(allAttributes) == 0 {
		log.Fatalf("attributes flag set with empty value, possibly not formatted properly?")
	}
	// Only render named attributes that include some data.
	d.PK = *pkFlag
	d.SK = *skFlag
	d.Attributes = getUsedAttributes(allAttributes, rows)
	d.MaxColCount = getMaxColCount(rows)

	// pkToIndexMap maps the partition key value to the index in the list of
	// partitions. Each row is added to a partition.
	pkToIndexMap := map[string]int{}
	for rowIndex, r := range rows {
		var pk, sk string
		for _, v := range r {
			if v.Key == d.PK {
				pk = v.Value
			}
			if v.Key == d.SK {
				sk = v.Value
			}
		}
		if pk == "" {
			log.Fatalf("row %d doesn't have a pk value (attribute named %q)", rowIndex, d.PK)
		}
		pkIndex, partitionAlreadyExists := pkToIndexMap[pk]
		if !partitionAlreadyExists {
			// Add a new partition to the slice, and keep track of its index in the map.
			d.Partitions = append(d.Partitions, Partition{PK: pk, SK: sk})
			pkToIndexMap[pk] = len(d.Partitions) - 1
			pkIndex = len(d.Partitions) - 1
		}
		// Add the row to the corresponding partition.
		d.Partitions[pkIndex].Rows = append(d.Partitions[pkIndex].Rows, r)
	}

	//TODO: Sort each partition by its sort key.

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

// IsAttribute is used when drawing out the output table. Named attributes are
// displayed first (on the left), with the remaining attributes of a row displayed
// to the right. Since the named attributes are written out separetely, when we
// write out the rest of the attributes, we want to skip the ones we've already
// written.
func (d Data) IsAttribute(key string) bool {
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
func (d Data) GetAttributeCount(row []value.Value) (count int) {
	for _, v := range row {
		if d.IsAttribute(v.Value) {
			continue
		}
		count++
	}
	return count
}

type Partition struct {
	// PK is the value of the partition key.
	PK string
	// SK is the value of the sort key.
	SK string
	// Rows are the rest of the named and un-named attributes.
	Rows [][]value.Value
}

func getValueOrEmptyString(key string, r []value.Value) string {
	for _, v := range r {
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
