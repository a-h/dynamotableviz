package value_test

import (
	"testing"

	"github.com/a-h/jsontable/value"
	"github.com/a-h/parse"
	"github.com/google/go-cmp/cmp"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]value.Value
	}{
		{
			name:  "values can be unquoted, if they don't contain quotes, equals signs, or commas",
			input: `a=this is a string,b=so is this`,
			expected: [][]value.Value{
				{
					value.Value{"a", "this is a string"},
					value.Value{"b", "so is this"},
				},
			},
		},
		{
			name:  "values can be quoted",
			input: `a="this is a string",b="so is this"`,
			expected: [][]value.Value{
				{
					value.Value{"a", "this is a string"},
					value.Value{"b", "so is this"},
				},
			},
		},
		{
			name:  "quoted strings can contain escaped quotes",
			input: `str="\"double quoted\"",num=456`,
			expected: [][]value.Value{
				{
					value.Value{"str", `"double quoted"`},
					value.Value{"num", "456"},
				},
			},
		},
		{
			name:  "quoted strings can contain commas",
			input: `str="\"double, quoted\"",num=456`,
			expected: [][]value.Value{
				{
					value.Value{"str", `"double, quoted"`},
					value.Value{"num", "456"},
				},
			},
		},
		{
			name:  "quoted strings can contain equals signs",
			input: `str="\"double, = quoted\"",num=456`,
			expected: [][]value.Value{
				{
					value.Value{"str", `"double, = quoted"`},
					value.Value{"num", "456"},
				},
			},
		},
		{
			name: "quoted strings can contain newlines",
			input: `str="new
line",num=456`,
			expected: [][]value.Value{
				{
					value.Value{"str", "new\nline"},
					value.Value{"num", "456"},
				},
			},
		},
		{
			name: "multiple rows can be returned",
			input: `str="\"double, = quoted\"",num=456
another=row,can=be,added=123`,
			expected: [][]value.Value{
				{
					value.Value{"str", `"double, = quoted"`},
					value.Value{"num", "456"},
				},
				{
					value.Value{"another", "row"},
					value.Value{"can", "be"},
					value.Value{"added", "123"},
				},
			},
		},
		{
			name:  "trailing newlines are permitted",
			input: "a=a,b=b\n\n\n",
			expected: [][]value.Value{
				{
					value.Value{"a", "a"},
					value.Value{"b", "b"},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rows, err := value.ParseAll(test.input)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
			}
			if diff := cmp.Diff(test.expected, rows); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestParserErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:     "there must be a key",
			input:    "b",
			expected: value.ErrParseKey{Position: parse.Position{Index: 1, Line: 0, Col: 1}},
		},
		{
			name:     "there must be a value",
			input:    "b=,c=",
			expected: value.ErrParseValue{Position: parse.Position{Index: 2, Line: 0, Col: 2}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := value.ParseAll(test.input)
			if diff := cmp.Diff(test.expected, err); diff != "" {
				t.Error(diff)
			}
		})
	}
}
