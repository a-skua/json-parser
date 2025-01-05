package token

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTokenize(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    []Token
		wantErr string
	}{
		"whitespace: empty": {
			input: "",
			want:  []Token{},
		},
		"whitespace: space": {
			input: " ",
			want: []Token{
				{Whitespace, " "},
			},
		},
		"whitespace: linefeed": {
			input: "\n",
			want: []Token{
				{Whitespace, "\n"},
			},
		},
		"whitespace: carriage return": {
			input: "\r",
			want: []Token{
				{Whitespace, "\r"},
			},
		},
		"whitespace: horizontal tab": {
			input: "\t",
			want: []Token{
				{Whitespace, "\t"},
			},
		},
		"whitespace: multiple": {
			input: "    \r\n\t",
			want: []Token{
				{Whitespace, "    \r\n\t"},
			},
		},
		"true: ok": {
			input: "true,",
			want: []Token{
				{True, "true"},
				{Comma, ","},
			},
		},
		"false: ok": {
			input: "false,",
			want: []Token{
				{False, "false"},
				{Comma, ","},
			},
		},
		"null: ok": {
			input: "null,",
			want: []Token{
				{Null, "null"},
				{Comma, ","},
			},
		},
		"string: ok": {
			input: `"Hello, 世界!"`,
			want: []Token{
				{String, `"Hello, 世界!"`},
			},
		},
		"string: emplyt": {
			input: `""`,
			want: []Token{
				{String, `""`},
			},
		},
		"string: escape quotation mark (\\\")": {
			input: `"\""`,
			want: []Token{
				{String, `"\""`},
			},
		},
		"string: escape reverse solidus (\\\\)": {
			input: `"\\"`,
			want: []Token{
				{String, `"\\"`},
			},
		},
		"string: escape solidus (\\/)": {
			input: `"\/"`,
			want: []Token{
				{String, `"\/"`},
			},
		},
		"string: escape backspace (\\b)": {
			input: `"\b"`,
			want: []Token{
				{String, `"\b"`},
			},
		},
		"string: escape formfeed (\\f)": {
			input: `"\f"`,
			want: []Token{
				{String, `"\f"`},
			},
		},
		"string: escape linefeed (\\n)": {
			input: `"\n"`,
			want: []Token{
				{String, `"\n"`},
			},
		},
		"string: escape carriage return (\\r)": {
			input: `"\r"`,
			want: []Token{
				{String, `"\r"`},
			},
		},
		"string: escape horizontal tab (\\t)": {
			input: `"\t"`,
			want: []Token{
				{String, `"\t"`},
			},
		},
		"string: escape unicode (\\uXXXX)": {
			input: `"\u3042"`,
			want: []Token{
				{String, `"\u3042"`},
			},
		},
		"string: ng invalid escape": {
			input:   `"\a"`,
			wantErr: "Invalid string: \"\\'a' (Unexpected Escape String: 'a')",
		},
		"string: ng invalid unicode (not hex digit)": {
			input:   `"\u30G2"`,
			wantErr: "Invalid string: \"\\u30'G' (Unexpected unicode: 'G')",
		},
		"number: 0": {
			input: "0",
			want: []Token{
				{Number, "0"},
			},
		},
		"number: -0": {
			input: "-0",
			want: []Token{
				{Number, "-0"},
			},
		},
		"number: 123": {
			input: "123",
			want: []Token{
				{Number, "123"},
			},
		},
		"number: -123": {
			input: "-123",
			want: []Token{
				{Number, "-123"},
			},
		},
		"number: 0.0": {
			input: "0.0",
			want: []Token{
				{Number, "0.0"},
			},
		},
		"number: -0.0": {
			input: "-0.0",
			want: []Token{
				{Number, "-0.0"},
			},
		},
		"number: 0e0": {
			input: "0e0",
			want: []Token{
				{Number, "0e0"},
			},
		},
		"number: 0E0": {
			input: "0E0",
			want: []Token{
				{Number, "0E0"},
			},
		},
		"number: 0e+0": {
			input: "0e+0",
			want: []Token{
				{Number, "0e+0"},
			},
		},
		"number: 0e-0": {
			input: "0e-0",
			want: []Token{
				{Number, "0e-0"},
			},
		},
		"number: 0.1": {
			input: "0.1",
			want: []Token{
				{Number, "0.1"},
			},
		},
		"number: 1e-1": {
			input: "1e-1",
			want: []Token{
				{Number, "1e-1"},
			},
		},
		"number: ng 0.a": {
			input:   "0.a",
			wantErr: "Invalid number: 0.'a' (Expected digit after '.': 'a')",
		},
		"number: ng 0.": {
			input:   "0.",
			wantErr: "Invalid number: '0.'",
		},
		"object": {
			input: `{"key": "value", "key2": "value2"}`,
			want: []Token{
				{LeftBrace, "{"},
				{String, `"key"`},
				{Colon, ":"},
				{Whitespace, " "},
				{String, `"value"`},
				{Comma, ","},
				{Whitespace, " "},
				{String, `"key2"`},
				{Colon, ":"},
				{Whitespace, " "},
				{String, `"value2"`},
				{RightBrace, "}"},
			},
		},
		"array": {
			input: `["value1", "value2"]`,
			want: []Token{
				{LeftBracket, "["},
				{String, `"value1"`},
				{Comma, ","},
				{Whitespace, " "},
				{String, `"value2"`},
				{RightBracket, "]"},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Tokenize([]rune(tt.input))
			if err != nil && err.Error() != tt.wantErr {
				t.Fatalf("Tokenize(%s) error: %v (want: %v)", tt.input, err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("Tokenize(%s) mismatch (-want +got):\n%s", tt.input, diff)
			}
		})
	}
}
