package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/a-skua/json-parser/node"
)

func main() {
	text := os.Args[1]

	nodes, err := node.Lex(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, n := range nodes {
		fmt.Println(sprintNode(n, 0))
	}
}

func sprintNode(n node.Node, depth int) string {
	switch n.Type() {
	case node.TypeObject:
		str := "{\n"
		for _, f := range n.Value().([]node.ObjectField) {
			str += fmt.Sprintf("%s\"%s\": %s,\n", indent(depth+1), f.Key, sprintNode(f.Value, depth+1))
		}
		return str + indent(depth) + "}"
	case node.TypeArray:
		str := "[\n"
		for _, v := range n.Value().([]node.Node) {
			str += fmt.Sprintf("%s%s,\n", indent(depth+1), sprintNode(v, depth+1))
		}
		return str + indent(depth) + "]"
	default:
		return n.String()
	}
}

func indent(n int) string {
	return strings.Repeat("  ", n)
}
