package token

import (
	"errors"
)

type Type uint8

const (
	_ Type = iota
	Object
	Array
	Value
	String
	Number
	Whitespace
	True
	False
	Null
)

type Token struct {
	Type  Type
	Value string
}

func Tokenize(data []rune) ([]Token, error) {
	tokens := []Token{}
	for i := 0; i < len(data); {
		switch {
		case isWhitespace(data[i]):
			t, n := tokenizeWhitespace(data[i:])
			i += n
			tokens = append(tokens, t)
		default:
			return nil, errors.New("Not implemented")
		}
	}

	return tokens, nil
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\r' || r == '\t'
}

func tokenizeWhitespace(data []rune) (Token, int) {
	whitespace := []rune{}
	for _, r := range data {
		if !isWhitespace(r) {
			break
		}
		whitespace = append(whitespace, r)
	}
	return Token{Type: Whitespace, Value: string(whitespace)}, len(whitespace)
}
