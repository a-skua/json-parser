package token

import (
	"fmt"

	"github.com/a-skua/json-parser/token/internal/runes"
	"github.com/a-skua/json-parser/token/internal/state"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=Type
type Type uint8

const (
	_ Type = iota
	Whitespace
	True
	False
	Null
	Number
	String
	LeftBrace
	RightBrace
	Colon
	Comma
	LeftBracket
	RightBracket
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
		case runes.IsWhitespace(data[i]):
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

		case maybeString(data[i:]):
			t, n, err := tokenizeString(data[i:])
			if err != nil {
				return nil, err
			}
			i += n
			tokens = append(tokens, t)

		case maybeNumber(data[i:]):
			t, n, err := tokenizeNumber(data[i:])
			if err != nil {
				return nil, err
			}
			i += n
			tokens = append(tokens, t)

		case runes.IsLeftBrace(data[i]):
			t, n := tokenizeLeftBrace(data[i:])
			i += n
			tokens = append(tokens, t)

		case runes.IsRightBrace(data[i]):
			t, n := tokenizeRightBrace(data[i:])
			i += n
			tokens = append(tokens, t)

		case runes.IsColon(data[i]):
			t, n := tokenizeColon(data[i:])
			i += n
			tokens = append(tokens, t)

		case runes.IsComma(data[i]):
			t, n := tokenizeComma(data[i:])
			i += n
			tokens = append(tokens, t)

		case runes.IsLeftBracket(data[i]):
			t, n := tokenizeLeftBracket(data[i:])
			i += n
			tokens = append(tokens, t)

		case runes.IsRightBracket(data[i]):
			t, n := tokenizeRightBracket(data[i:])
			i += n
			tokens = append(tokens, t)

		default:
			return nil, fmt.Errorf("Unexpected token: '%s'", string(data[i]))
		}
	}

	return tokens, nil
}

func tokenizeLeftBrace(data []rune) (Token, int) {
	return New(LeftBrace, data[:1]), 1
}

func tokenizeRightBrace(data []rune) (Token, int) {
	return New(RightBrace, data[:1]), 1
}

func tokenizeLeftBracket(data []rune) (Token, int) {
	return New(LeftBracket, data[:1]), 1
}

func tokenizeRightBracket(data []rune) (Token, int) {
	return New(RightBracket, data[:1]), 1
}

func tokenizeColon(data []rune) (Token, int) {
	return New(Colon, data[:1]), 1
}

func tokenizeComma(data []rune) (Token, int) {
	return New(Comma, data[:1]), 1
}

func tokenizeWhitespace(data []rune) (Token, int) {
	whitespace := []rune{}
	for _, r := range data {
		if !runes.IsWhitespace(r) {
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

	if 4 < len(data) && !runes.IsWhitespace(data[4]) {
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

	if 5 < len(data) && !runes.IsWhitespace(data[5]) {
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

	if 4 < len(data) && !runes.IsWhitespace(data[4]) {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data[4]))
	}

	return New(Null, data[:4]), 4, nil
}

func maybeNumber(data []rune) bool {
	return runes.IsDigit(data[0]) || data[0] == '-'
}

func tokenizeNumber(data []rune) (Token, int, error) {
	state := state.NewNumber()
	number := []rune{}

	for _, r := range data {
		var err error
		state, err = state.Next(r)
		if err != nil {
			return Token{}, 0, fmt.Errorf("Invalid number: %s'%s' (%w)", string(number), string(r), err)
		}
		if state.IsEnd() {
			break
		}
		number = append(number, r)
	}

	if !state.Valid() {
		return Token{}, 0, fmt.Errorf("Invalid number: '%s'", string(number))
	}

	return New(Number, number), len(number), nil
}

func isHex(r rune) bool {
	return runes.IsDigit(r) || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
}

func maybeString(data []rune) bool {
	return 2 <= len(data) && data[0] == '"'
}

func validStringEscapeSequence(data []rune) (int, error) {
	if len(data) < 2 {
		return 0, fmt.Errorf("Unexpected token: '%s'", string(data))
	}

	if data[0] != '\\' {
		return 0, fmt.Errorf("Unexpected token: '\\'")
	}

	switch data[1] {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		return 2, nil
	case 'u':
		if len(data) < 6 {
			return 0, fmt.Errorf("Unexpected unicode: '%s'", string(data))
		}
		for _, r := range data[2:6] {
			if !isHex(r) {
				return 0, fmt.Errorf("Unexpected unicode: '%s'", string(data[:6]))
			}
		}
		return 6, nil
	default:
		return 0, fmt.Errorf("Unexpected Escape String: '%s'", string(data[:2]))
	}
}

func tokenizeString(data []rune) (Token, int, error) {

	if l := len(data); l < 2 || (l == 2 && data[0] != '"') {
		return Token{}, 0, fmt.Errorf("Unexpected token: '%s'", string(data))
	}

	i := 1
	for ; i < len(data); i++ {
		if data[i] == '"' {
			break
		}

		if data[i] == '\\' {
			n, err := validStringEscapeSequence(data[i:])
			if err != nil {
				return Token{}, 0, err
			}
			i += n - 1
		}
	}

	return New(String, data[:i+1]), i + 1, nil
}
