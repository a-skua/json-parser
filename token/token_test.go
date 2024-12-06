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
