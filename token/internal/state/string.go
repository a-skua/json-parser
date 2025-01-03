package state

import (
	"fmt"

	"github.com/a-skua/json-parser/token/internal/runes"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=String
type String uint8

const (
	_ String = iota
	StringStart
	StringFirstQuote
	StringCodepoint
	StringEscapeSymbol
	StringEscape
	StringHexDigitSymbol
	StringHexDigit1
	StringHexDigit2
	StringHexDigit3
	StringHexDigit4
	StringLastQuote
	StringEnd
)

func NewString() String {
	return StringStart
}

func (s String) IsEnd() bool {
	return s == StringEnd
}

func (s String) Valid() bool {
	switch s {
	case StringLastQuote, StringEnd:
		return true
	default:
		return false
	}
}

func (s String) Next(r rune) (String, error) {
	switch s {
	case StringStart:
		return s.startNext(r)
	case StringFirstQuote, StringEscape, StringHexDigit4, StringCodepoint:
		return s.codepointNext(r)
	case StringEscapeSymbol:
		return s.escapeSymbolNext(r)
	case StringHexDigitSymbol, StringHexDigit1, StringHexDigit2, StringHexDigit3:
		return s.hexDigitNext(r)
	}

	return StringEnd, nil
}

func (s String) startNext(r rune) (String, error) {
	if r == '"' {
		return StringFirstQuote, nil
	}

	return 0, fmt.Errorf("Expected '\"' at the start of a string: %c", r)
}

func (s String) codepointNext(r rune) (String, error) {
	if r == '"' {
		return StringLastQuote, nil
	}

	if r == '\\' {
		return StringEscapeSymbol, nil
	}

	return StringCodepoint, nil
}

func (s String) escapeSymbolNext(r rune) (String, error) {
	switch r {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		return StringEscape, nil
	case 'u':
		return StringHexDigitSymbol, nil
	default:
		return 0, fmt.Errorf("Unexpected Escape String: '%c'", r)
	}
}

func (s String) hexDigitNext(r rune) (String, error) {
	if runes.IsHex(r) {
		switch s {
		case StringHexDigitSymbol:
			return StringHexDigit1, nil
		case StringHexDigit1:
			return StringHexDigit2, nil
		case StringHexDigit2:
			return StringHexDigit3, nil
		case StringHexDigit3:
			return StringHexDigit4, nil
		}
	}

	return 0, fmt.Errorf("Unexpected unicode: '%c'", r)
}
