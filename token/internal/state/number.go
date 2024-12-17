package state

import (
	"fmt"

	"github.com/a-skua/json-parser/token/internal/runes"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=Number
type Number uint8

const (
	_ Number = iota
	NumberStart
	NumberSign
	NumberZero
	NumberInteger
	NumberFractionSymol
	NumberFraction
	NumberExponentSymbol
	NumberExponentSign
	NumberExponent
	NumberEnd
)

func NewNumber() Number {
	return NumberStart
}

func (s Number) IsEnd() bool {
	return s == NumberEnd
}

func (s Number) Valid() bool {
	switch s {
	case NumberZero, NumberInteger, NumberFraction, NumberExponent, NumberEnd:
		return true
	default:
		return false
	}
}

func (s Number) Next(r rune) (Number, error) {
	switch s {
	case NumberStart:
		return s.startNext(r)
	case NumberSign:
		return s.signNext(r)
	case NumberZero:
		return s.zeroNext(r)
	case NumberInteger:
		return s.integerNext(r)
	case NumberFractionSymol:
		return s.fractionSymbolNext(r)
	case NumberFraction:
		return s.fractionNext(r)
	case NumberExponentSymbol:
		return s.exponentSymbolNext(r)
	case NumberExponentSign:
		return s.exponentSignNext(r)
	case NumberExponent:
		return s.exponentNext(r)
	}

	return NumberEnd, nil
}

// start => '-' | '0' | [1-9]
func (s Number) startNext(r rune) (Number, error) {
	if r == '-' {
		return NumberSign, nil
	}

	if r == '0' {
		return NumberZero, nil
	}

	if runes.IsDigit(r) {
		return NumberInteger, nil
	}

	return 0, fmt.Errorf("Expected sign or digit: '%s'", string(r))
}

// '-' => [1-9] | 0
func (s Number) signNext(r rune) (Number, error) {
	if r == '0' {
		return NumberZero, nil
	}

	if runes.IsDigit(r) {
		return NumberInteger, nil
	}

	return 0, fmt.Errorf("Expected digit after sign: '%s'", string(r))
}

// '0' => '.' | 'e' | 'E' | end
func (s Number) zeroNext(r rune) (Number, error) {
	if r == '.' {
		return NumberFractionSymol, nil
	}

	if r == 'e' || r == 'E' {
		return NumberExponentSymbol, nil
	}

	return NumberEnd, nil
}

// [0-9] => [0-9] | '.' | 'e' | 'E' | end
func (s Number) integerNext(r rune) (Number, error) {
	if runes.IsDigit(r) {
		return NumberInteger, nil
	}

	if r == '.' {
		return NumberFractionSymol, nil
	}

	if r == 'e' || r == 'E' {
		return NumberExponentSymbol, nil
	}

	return NumberEnd, nil
}

// '.' => [0-9]
func (s Number) fractionSymbolNext(r rune) (Number, error) {
	if runes.IsDigit(r) {
		return NumberFraction, nil
	}

	return 0, fmt.Errorf("Expected digit after '.': '%s'", string(r))
}

// [0-9] => [0-9] | 'e' | 'E' | end
func (s Number) fractionNext(r rune) (Number, error) {
	if runes.IsDigit(r) {
		return NumberFraction, nil
	}

	if r == 'e' || r == 'E' {
		return NumberExponentSymbol, nil
	}

	return NumberEnd, nil
}

// 'e' | 'E' => '+' | '-' | [0-9]
func (s Number) exponentSymbolNext(r rune) (Number, error) {
	if r == '+' || r == '-' {
		return NumberExponentSign, nil
	}

	if runes.IsDigit(r) {
		return NumberExponent, nil
	}

	return 0, fmt.Errorf("Expected digit or sign after 'e' or 'E': '%s'", string(r))
}

// '+' | '-' => [0-9]
func (s Number) exponentSignNext(r rune) (Number, error) {
	if runes.IsDigit(r) {
		return NumberExponent, nil
	}

	return 0, fmt.Errorf("Expected digit after sign: '%s'", string(r))
}

// [0-9] => [0-9] | end
func (s Number) exponentNext(r rune) (Number, error) {
	if runes.IsDigit(r) {
		return NumberExponent, nil
	}

	return NumberEnd, nil
}
