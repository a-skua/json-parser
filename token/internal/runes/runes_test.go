package runes

import "testing"

func TestIsDigit(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"0": {
			input: "0",
			want:  true,
		},
		"9": {
			input: "9",
			want:  true,
		},
		"a": {
			input: "a",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsDigit(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsDigit(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}
