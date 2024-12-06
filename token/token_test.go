package token

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    []Token
		wantErr bool
	}{
		"whitespace: empty": {
			input: "",
			want:  []Token{},
		},
		"whitespace: space": {
			input: " ",
			want: []Token{
				{Type: Whitespace, Value: " "},
			},
		},
		"whitespace: linefeed": {
			input: "\n",
			want: []Token{
				{Type: Whitespace, Value: "\n"},
			},
		},
		"whitespace: carriage return": {
			input: "\r",
			want: []Token{
				{Type: Whitespace, Value: "\r"},
			},
		},
		"whitespace: horizontal tab": {
			input: "\t",
			want: []Token{
				{Type: Whitespace, Value: "\t"},
			},
		},
		"whitespace: multiple": {
			input: "    \r\n\t",
			want: []Token{
				{Type: Whitespace, Value: "    \r\n\t"},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Tokenize([]rune(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("Tokenize() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("Tokenize() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
