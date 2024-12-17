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
			input: "true",
			want: []Token{
				{Type: True, Value: "true"},
			},
		},
		"true: ng": {
			input:   "true,",
			wantErr: "Unexpected token: ','",
		},
		"false: ok": {
			input: "false",
			want: []Token{
				{Type: False, Value: "false"},
			},
		},
		"false: ng": {
			input:   "false,",
			wantErr: "Unexpected token: ','",
		},
		"null: ok": {
			input: "null",
			want: []Token{
				{Type: Null, Value: "null"},
			},
		},
		"null: ng": {
			input:   "null,",
			wantErr: "Unexpected token: ','",
		},
		"string: ok": {
			input: `"Hello, 世界!"`,
			want: []Token{
				{Type: String, Value: `"Hello, 世界!"`},
			},
		},
		"string: emplyt": {
			input: `""`,
			want: []Token{
				{Type: String, Value: `""`},
			},
		},
		"string: escape quotation mark (\\\")": {
			input: `"\""`,
			want: []Token{
				{Type: String, Value: `"\""`},
			},
		},
		"string: escape reverse solidus (\\\\)": {
			input: `"\\"`,
			want: []Token{
				{Type: String, Value: `"\\"`},
			},
		},
		"string: escape solidus (\\/)": {
			input: `"\/"`,
			want: []Token{
				{Type: String, Value: `"\/"`},
			},
		},
		"string: escape backspace (\\b)": {
			input: `"\b"`,
			want: []Token{
				{Type: String, Value: `"\b"`},
			},
		},
		"string: escape formfeed (\\f)": {
			input: `"\f"`,
			want: []Token{
				{Type: String, Value: `"\f"`},
			},
		},
		"string: escape linefeed (\\n)": {
			input: `"\n"`,
			want: []Token{
				{Type: String, Value: `"\n"`},
			},
		},
		"string: escape carriage return (\\r)": {
			input: `"\r"`,
			want: []Token{
				{Type: String, Value: `"\r"`},
			},
		},
		"string: escape horizontal tab (\\t)": {
			input: `"\t"`,
			want: []Token{
				{Type: String, Value: `"\t"`},
			},
		},
		"string: escape unicode (\\uXXXX)": {
			input: `"\u3042"`,
			want: []Token{
				{Type: String, Value: `"\u3042"`},
			},
		},
		"string: ng invalid escape": {
			input:   `"\a"`,
			wantErr: "Unexpected Escape String: '\\a'",
		},
		"string: ng invalid unicode (not hex digit)": {
			input:   `"\u30G2"`,
			wantErr: "Unexpected unicode: '\\u30G2'",
		},
		"number: 0": {
			input: "0",
			want: []Token{
				{Type: Number, Value: "0"},
			},
		},
		"number: -0": {
			input: "-0",
			want: []Token{
				{Type: Number, Value: "-0"},
			},
		},
		"number: 123": {
			input: "123",
			want: []Token{
				{Type: Number, Value: "123"},
			},
		},
		"number: -123": {
			input: "-123",
			want: []Token{
				{Type: Number, Value: "-123"},
			},
		},
		"number: 0.0": {
			input: "0.0",
			want: []Token{
				{Type: Number, Value: "0.0"},
			},
		},
		"number: -0.0": {
			input: "-0.0",
			want: []Token{
				{Type: Number, Value: "-0.0"},
			},
		},
		"number: 0e0": {
			input: "0e0",
			want: []Token{
				{Type: Number, Value: "0e0"},
			},
		},
		"number: 0E0": {
			input: "0E0",
			want: []Token{
				{Type: Number, Value: "0E0"},
			},
		},
		"number: 0e+0": {
			input: "0e+0",
			want: []Token{
				{Type: Number, Value: "0e+0"},
			},
		},
		"number: 0e-0": {
			input: "0e-0",
			want: []Token{
				{Type: Number, Value: "0e-0"},
			},
		},
		"number: 0.1": {
			input: "0.1",
			want: []Token{
				{Type: Number, Value: "0.1"},
			},
		},
		"number: 1e-1": {
			input: "1e-1",
			want: []Token{
				{Type: Number, Value: "1e-1"},
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
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Tokenize([]rune(tt.input))
			if err != nil && err.Error() != tt.wantErr {
				t.Fatalf("Tokenize() error: %v (want: %v)", err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("Tokenize() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
