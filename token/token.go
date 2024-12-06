package token

import (
	"fmt"
)

type Type uint8

const (
	_ Type = iota
	Whitespace
	True
	False
	Null
)

type Token struct {
	Type  Type
	Value string
}

func New(t Type, runes []rune) Token {
	return Token{Type: t, Value: string(runes)}
}

func Tokenize(data []rune) ([]Token, error) {
	tokens := []Token{}

	for i := 0; i < len(data); {
		switch {
		case isWhitespace(data[i]):
			t, n := tokenizeWhitespace(data[i:])
			i += n
			tokens = append(tokens, t)

		case maybeTrue(data[i:]):
			t, n, err := tokenizeTrue(data[i:])
			if err != nil {
				return nil, err
			}
			i += n
			tokens = append(tokens, t)

		case maybeFalse(data[i:]):
			t, n, err := tokenizeFalse(data[i:])
			if err != nil {
				return nil, err
			}
			i += n
			tokens = append(tokens, t)

		case maybeNull(data[i:]):
			t, n, err := tokenizeNull(data[i:])
			if err != nil {
				return nil, err
			}
			i += n
			tokens = append(tokens, t)

		default:
			return nil, fmt.Errorf("Unexpected token: '%s'", string(data[i]))
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
	return New(Whitespace, whitespace), len(whitespace)
}

func maybeTrue(data []rune) bool {
	return 4 <= len(data) && string(data[:4]) == "true"
}

func tokenizeTrue(data []rune) (Token, int, error) {
	if l := len(data); l < 4 || (l == 4 && string(data) != "true") {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data))
	}

	if 4 < len(data) && !isWhitespace(data[4]) {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data[4]))
	}

	return New(True, data[:4]), 4, nil
}

func maybeFalse(data []rune) bool {
	return 5 <= len(data) && string(data[:5]) == "false"
}

func tokenizeFalse(data []rune) (Token, int, error) {
	if l := len(data); l < 5 || (l == 5 && string(data) != "false") {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data))
	}

	if 5 < len(data) && !isWhitespace(data[5]) {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data[5]))
	}

	return New(False, data[:5]), 5, nil
}

func maybeNull(data []rune) bool {
	return 4 <= len(data) && string(data[:4]) == "null"
}

func tokenizeNull(data []rune) (Token, int, error) {
	if l := len(data); l < 4 || (l == 4 && string(data) != "null") {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data))
	}

	if 4 < len(data) && !isWhitespace(data[4]) {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data[4]))
	}

	return New(Null, data[:4]), 4, nil
}
