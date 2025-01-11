package node

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/a-skua/json-parser/node/internal/state"
	"github.com/a-skua/json-parser/token"
)

var (
	ErrEON     = errors.New("End of Node")
	ErrEOA     = errors.New("End of Array")
	ErrEOO     = errors.New("End of Object")
	ErrIsComma = errors.New("Token is comma")
	ErrIsColon = errors.New("Token is colon")
)

type Type uint8

const (
	_ Type = iota
	TypeObject
	TypeArray
	TypeString
	TypeNumber
	TypeBoolean
	TypeNull
)

type Node interface {
	Type() Type
	Value() interface{}
	String() string
}

type String struct {
	value string
}

func newString(t token.Token) String {
	return String{t.Value[1 : len(t.Value)-1]}
}

func (s String) Type() Type {
	return TypeString
}

func (s String) Value() interface{} {
	return s.value
}

func (s String) String() string {
	return "\"" + s.value + "\""
}

type Number struct {
	value float64
}

func newNumber(t token.Token) (Number, error) {
	value, err := strconv.ParseFloat(t.Value, 64)
	return Number{value}, err
}

func (n Number) Type() Type {
	return TypeNumber
}

func (n Number) Value() interface{} {
	return n.value
}

func (n Number) String() string {
	return strconv.FormatFloat(n.value, 'g', -1, 64)
}

func Lex(input string) ([]Node, error) {
	var err error
	nodes := make([]Node, 0)

	lexer := NewLexer(token.NewTokenizer([]rune(input)))
	for {
		var node Node
		node, err = lexer.Next()
		if err != nil {
			break
		}
		nodes = append(nodes, node)
	}

	if err == ErrEON {
		return nodes, nil
	}

	return nodes, err
}

type Boolean struct {
	value bool
}

func newBoolean(t token.Token) Boolean {
	return Boolean{t.Type == token.True}
}

func (b Boolean) Type() Type {
	return TypeBoolean
}

func (b Boolean) Value() interface{} {
	return b.value
}

type Null struct{}

func newNull() Null {
	return Null{}
}

func (n Null) Type() Type {
	return TypeNull
}

func (n Null) Value() interface{} {
	return nil
}

func (n Null) String() string {
	return "null"
}

func (b Boolean) String() string {
	if b.value {
		return "true"
	} else {
		return "false"
	}
}

type Array struct {
	nodes []Node
}

func (a Array) Type() Type {
	return TypeArray
}

func (a Array) Value() interface{} {
	return a.nodes
}

func (a Array) String() string {
	str := "["
	for i, node := range a.nodes {
		if i > 0 {
			str += ","
		}
		str += node.String()
	}
	return str + "]"
}

type Lexer struct {
	tokenizer token.Tokenizer
}

func NewLexer(tokenizer token.Tokenizer) Lexer {
	return Lexer{tokenizer: tokenizer}
}

type ObjectField struct {
	Key   string
	Value Node
}

type Object struct {
	fields []ObjectField
}

func (o Object) Type() Type {
	return TypeObject
}

func (o Object) Value() interface{} {
	return o.fields
}

func (o Object) String() string {
	str := "{"
	for i, field := range o.fields {
		if i > 0 {
			str += ","
		}
		str += "\"" + field.Key + "\":" + field.Value.String()
	}
	return str + "}"
}

func (l *Lexer) Next() (Node, error) {
	var t token.Token
	var err error
	for t, err = l.tokenizer.Next(); err == nil; t, err = l.tokenizer.Next() {
		switch t.Type {
		case token.String:
			return newString(t), nil
		case token.Number:
			return newNumber(t)
		case token.True, token.False:
			return newBoolean(t), nil
		case token.Null:
			return newNull(), nil
		case token.Whitespace:
			continue
		case token.LeftBracket:
			return l.parseArray()
		case token.RightBracket:
			return nil, ErrEOA
		case token.LeftBrace:
			return l.parseObject()
		case token.RightBrace:
			return nil, ErrEOO
		case token.Comma:
			return nil, ErrIsComma
		case token.Colon:
			return nil, ErrIsColon
		default:
			return nil, fmt.Errorf("Unexpected Token: %v", t)
		}
	}

	if err == token.ErrEOT {
		return nil, ErrEON
	}

	return nil, err
}

func (l *Lexer) parseArray() (Node, error) {
	nodes := make([]Node, 0)
	for state := state.NewArray(); ; state = state.Next() {
		node, err := l.Next()
		if state.IsSeparator() && err == ErrIsComma {
			continue
		}

		if state.IsSeparator() && err == ErrEOA {
			break
		}

		if err == ErrEOA {
			return nil, errors.New("Unexpected End of Array")
		}

		if err == ErrIsComma {
			return nil, errors.New("Unexpected Comma")
		}

		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return Array{nodes}, nil
}

func (l *Lexer) parseObject() (Node, error) {
	fields := make([]ObjectField, 0)
	for state := state.NewObject(); ; state = state.Next() {
		key, err := l.Next()
		if state.IsSeparator() && err == ErrIsComma {
			continue
		}
		if err == ErrEOO {
			break
		}

		if !state.IsKey() || key.Type() != TypeString {
			return nil, fmt.Errorf("Unexpected Token: %v", key)
		}
		state = state.Next()

		colon, err := l.Next()
		if !state.IsColon() || err != ErrIsColon {
			return nil, fmt.Errorf("Unexpected Token: %v", colon)
		}
		state = state.Next()

		value, err := l.Next()
		if !state.IsValue() || err != nil {
			return nil, err
		}

		fields = append(fields, ObjectField{key.Value().(string), value})
	}

	return Object{fields}, nil
}
