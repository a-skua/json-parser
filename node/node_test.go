package node

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    string
		wantErr string
	}{
		"empty": {
			input: "",
			want:  "[]",
		},
		"string": {
			input: `"hello"`,
			want:  `["hello"]`,
		},
		"number": {
			input: "123.45",
			want:  "[123.45]",
		},
		"string & number": {
			input: `"hello"
123`,
			want: `["hello" 123]`,
		},
		"boolean (true)": {
			input: "true",
			want:  "[true]",
		},
		"boolean (false)": {
			input: "false",
			want:  "[false]",
		},
		"null": {
			input: "null",
			want:  "[null]",
		},
		"array": {
			input: `["hello", 123, true, false, null]
["hello"]
[123]
[true]
[false]
[null]
[
  ["hello", 123],
  [true, false, null]
]
[]`,
			want: `[["hello",123,true,false,null] ["hello"] [123] [true] [false] [null] [["hello",123],[true,false,null]] []]`,
		},
		"object": {
			input: `{"key1": "value1", "key2": 123, "key3": true, "key4": false, "key5": null}
{"array": ["hello", 123, true, false, null]}
{}`,
			want: `[{"key1":"value1","key2":123,"key3":true,"key4":false,"key5":null} {"array":["hello",123,true,false,null]} {}]`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			nodes, err := Lex(tt.input)
			if err != nil && err.Error() != tt.wantErr {
				t.Fatalf("Lex(%s) error: %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			got := fmt.Sprint(nodes)
			if tt.want != got {
				t.Fatalf("Lex(%s) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
