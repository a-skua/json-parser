package token

import (
	"github.com/google/go-cmp/cmp"
	"testing"
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
