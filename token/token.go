package token

import (
	"errors"
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

type Tokenizer struct {
	data []rune
}

func NewTokenizer(data []rune) Tokenizer {
	return Tokenizer{data: data}
}

var ErrEOT = errors.New("End of Token")

func (t *Tokenizer) Next() (Token, error) {
	if len(t.data) == 0 {
		return Token{}, ErrEOT
	}

	switch {
	case runes.IsWhitespace(t.data[0]):
		token, n := tokenizeWhitespace(t.data)
		t.data = t.data[n:]
		return token, nil

	case runes.MaybeTrue(t.data):
		token, n, err := tokenizeTrue(t.data)
		if err != nil {
			return Token{}, err
		}
		t.data = t.data[n:]
		return token, nil

	case runes.MaybeFalse(t.data):
		token, n, err := tokenizeFalse(t.data)
		if err != nil {
			return Token{}, err
		}
		t.data = t.data[n:]
		return token, nil

	case runes.MaybeNull(t.data):
		token, n, err := tokenizeNull(t.data)
		if err != nil {
			return Token{}, err
		}
		t.data = t.data[n:]
		return token, nil

	case runes.MaybeString(t.data):
		token, n, err := tokenizeString(t.data)
		if err != nil {
			return Token{}, err
		}
		t.data = t.data[n:]
		return token, nil

	case runes.MaybeNumber(t.data[0]):
		token, n, err := tokenizeNumber(t.data)
		if err != nil {
			return Token{}, err
		}
		t.data = t.data[n:]
		return token, nil

	case runes.IsLeftBrace(t.data[0]):
		token, n := tokenizeLeftBrace(t.data)
		t.data = t.data[n:]
		return token, nil

	case runes.IsRightBrace(t.data[0]):
		token, n := tokenizeRightBrace(t.data)
		t.data = t.data[n:]
		return token, nil

	case runes.IsColon(t.data[0]):
		token, n := tokenizeColon(t.data)
		t.data = t.data[n:]
		return token, nil

	case runes.IsComma(t.data[0]):
		token, n := tokenizeComma(t.data)
		t.data = t.data[n:]
		return token, nil

	case runes.IsLeftBracket(t.data[0]):
		token, n := tokenizeLeftBracket(t.data)
		t.data = t.data[n:]
		return token, nil

	case runes.IsRightBracket(t.data[0]):
		token, n := tokenizeRightBracket(t.data)
		t.data = t.data[n:]
		return token, nil

	default:
		return Token{}, fmt.Errorf("Unexpected token: '%s'", string(t.data[0]))
	}
}

func Tokenize(data []rune) ([]Token, error) {
	tokenizer := NewTokenizer(data)
	tokens := []Token{}

	for {
		t, err := tokenizer.Next()
		if err == ErrEOT {
			break
		}
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, t)
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

func tokenizeTrue(data []rune) (Token, int, error) {
	if l := len(data); (l < 4) ||
		(l == 4 && string(data) != "true") {
		return Token{}, 0, fmt.Errorf("Unexpected Token: '%s'", string(data))
	}
	if l := len(data); 4 < l &&
		!runes.IsWhitespace(data[4]) &&
		!runes.IsComma(data[4]) &&
		!runes.IsRightBracket(data[4]) &&
		!runes.IsRightBrace(data[4]) {
		return Token{}, 0, fmt.Errorf("Unexpected Token: true'%s'", string(data[4]))
	}

	return New(True, data[:4]), 4, nil
}

func tokenizeFalse(data []rune) (Token, int, error) {
	if l := len(data); (l < 5) ||
		(l == 5 && string(data) != "false") {
		return Token{}, 0, fmt.Errorf("Unexpected Token:'%s'", string(data))
	}

	if l := len(data); 5 < l &&
		!runes.IsWhitespace(data[5]) &&
		!runes.IsComma(data[5]) &&
		!runes.IsRightBracket(data[5]) &&
		!runes.IsRightBrace(data[5]) {
		return Token{}, 0, fmt.Errorf("Unexpected Token: false'%s'", string(data[5]))
	}

	return New(False, data[:5]), 5, nil
}

func tokenizeNull(data []rune) (Token, int, error) {
	if l := len(data); (l < 4) ||
		(l == 4 && string(data) != "null") {
		return Token{}, 0, fmt.Errorf("Unexpected Token: '%s'", string(data))
	}

	if l := len(data); 4 < l &&
		!runes.IsWhitespace(data[4]) &&
		!runes.IsComma(data[4]) &&
		!runes.IsRightBracket(data[4]) &&
		!runes.IsRightBrace(data[4]) {
		return Token{}, 0, fmt.Errorf("Unexpected Token: null'%s'", string(data[4]))
	}

	return New(Null, data[:4]), 4, nil
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

func tokenizeString(data []rune) (Token, int, error) {
	state := state.NewString()
	strings := []rune{}

	for _, r := range data {
		var err error
		state, err = state.Next(r)
		if err != nil {
			return Token{}, 0, fmt.Errorf("Invalid string: %s'%s' (%w)", string(strings), string(r), err)
		}
		if state.IsEnd() {
			break
		}
		strings = append(strings, r)
	}

	if !state.Valid() {
		return Token{}, 0, fmt.Errorf("Invalid string: '%s'", string(strings))
	}

	return New(String, strings), len(strings), nil
}
