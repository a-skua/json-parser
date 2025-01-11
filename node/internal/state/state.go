package state

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=Array
type Array uint8

const (
	_ Array = iota
	ArrayFirstValue
	ArrayValue
	ArraySeparator
)

func NewArray() Array {
	return ArrayFirstValue
}

func (a Array) Next() Array {
	switch a {
	case ArrayValue, ArrayFirstValue:
		return ArraySeparator
	default:
		return ArrayValue
	}
}

func (a Array) IsValue() bool {
	return a == ArrayValue || a == ArrayFirstValue
}

func (a Array) IsSeparator() bool {
	return a == ArraySeparator || a == ArrayFirstValue
}

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=Object
type Object uint8

const (
	_ Object = iota
	ObjectKey
	ObjectColon
	ObjectValue
	ObjectSeparator
)

func NewObject() Object {
	return ObjectKey
}

func (o Object) Next() Object {
	switch o {
	case ObjectKey:
		return ObjectColon
	case ObjectColon:
		return ObjectValue
	case ObjectValue:
		return ObjectSeparator
	default:
		return ObjectKey
	}
}

func (o Object) IsKey() bool {
	return o == ObjectKey
}

func (o Object) IsColon() bool {
	return o == ObjectColon
}

func (o Object) IsValue() bool {
	return o == ObjectValue
}

func (o Object) IsSeparator() bool {
	return o == ObjectSeparator
}
