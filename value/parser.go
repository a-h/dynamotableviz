package value

import (
	"fmt"
	"strings"

	"github.com/a-h/parse"
)

var keyValueDelimiter = parse.Rune('=')
var rowDelimiter = parse.NewLine
var colDelimiter = parse.Rune(',')
var doubleQuote = parse.Rune('"')
var escape = parse.Rune('\\')
var escapedQuote = parse.String(`\"`)
var stringUntilEscapedCharacterOrDoubleQuote = parse.StringUntil(parse.Any(escape, doubleQuote))

type Value struct {
	Key   string
	Value string
}

func (v Value) String() string {
	return fmt.Sprintf(`%v=%q`, v.Key, v.Value)
}

type quotedStringParser struct{}

func (p quotedStringParser) Parse(in parse.Input) (match string, ok bool, err error) {
	start := in.Index()
	// Start with a quote.
	_, ok, err = doubleQuote.Parse(in)
	if err != nil || !ok {
		// No match, so rewind.
		in.Seek(start)
		return
	}
	// Grab the contents.
	var sb strings.Builder
	for {
		// Try for an escaped quote.
		_, ok, err = escapedQuote.Parse(in)
		if err != nil {
			return
		}
		if ok {
			sb.WriteRune('"')
			continue
		}
		// Or a terminating quote.
		_, ok, err = doubleQuote.Parse(in)
		if err != nil {
			return
		}
		if ok {
			break
		}
		// Grab the runes.
		match, ok, err = stringUntilEscapedCharacterOrDoubleQuote.Parse(in)
		if err != nil {
			return
		}
		if ok {
			sb.WriteString(match)
			continue
		}
		// If we haven't gotten a match, we must have reached the end of the file.
		// Without closing the string.
		err = fmt.Errorf("unterminated quoted string from %v to %v", in.PositionAt(start), in.Position())
	}
	match = sb.String()
	return
}

type ErrParseKey struct {
	Position parse.Position
}

func (e ErrParseKey) Error() string {
	return fmt.Sprintf("parse error: expected 'key=' not found at line %d, col %d (index %d)", e.Position.Line, e.Position.Col, e.Position.Index)
}

type ErrParseDelimiter struct {
	Position parse.Position
}

func (e ErrParseDelimiter) Error() string {
	return fmt.Sprintf("parse error: expected '=' not found at line %d, col %d (index %d)", e.Position.Line, e.Position.Col, e.Position.Index)
}

type ErrParseValue struct {
	Position parse.Position
}

func (e ErrParseValue) Error() string {
	return fmt.Sprintf("parse error: expected value not found at line %d, col %d (index %d)", e.Position.Line, e.Position.Col, e.Position.Index)
}

var quotedString = parse.Parser[string](quotedStringParser{})
var unquotedString = parse.StringUntil(parse.Any(colDelimiter, rowDelimiter, parse.EOF[string]()))
var valueParser = parse.Func(func(in parse.Input) (match Value, ok bool, err error) {
	// Read the key.
	// key="value"
	// ^^^
	match.Key, ok, err = parse.StringUntil(keyValueDelimiter).Parse(in)
	if err != nil {
		return
	}
	if !ok {
		return match, ok, ErrParseKey{in.Position()}
	}
	// Chomp the key/value delimiter, which is definitely already there.
	// key="value"
	//    ^
	keyValueDelimiter.Parse(in)

	// Read the quoted/unquoted string.
	match.Value, ok, err = parse.Any(quotedString, unquotedString).Parse(in)
	if err != nil {
		return
	}
	if !ok || match.Value == "" {
		return match, ok, ErrParseValue{in.Position()}
	}

	// Chomp the optional col delimiter.
	colDelimiter.Parse(in)
	return
})

var row = parse.Func(func(in parse.Input) (match []Value, ok bool, err error) {
	match, ok, err = parse.UntilEOF(valueParser, rowDelimiter).Parse(in)
	// Chomp the row terminator.
	rowDelimiter.Parse(in)
	return
})

var rows = parse.UntilEOF(row, parse.Any(colDelimiter, parse.NewLine))

func ParseAll(s string) (v [][]Value, err error) {
	v, _, err = rows.Parse(parse.NewInput(s))
	return
}
