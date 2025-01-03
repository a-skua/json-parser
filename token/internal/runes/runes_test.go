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

func TestIsWhitespace(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"space": {
			input: " ",
			want:  true,
		},
		"newline": {
			input: "\n",
			want:  true,
		},
		"carriage return": {
			input: "\r",
			want:  true,
		},
		"tab": {
			input: "\t",
			want:  true,
		},
		"a": {
			input: "a",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsWhitespace(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsWhitespace(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsLeftBrace(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"{": {
			input: "{",
			want:  true,
		},
		"}": {
			input: "}",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsLeftBrace(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsLeftBrace(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsRightBrace(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"{": {
			input: "{",
			want:  false,
		},
		"}": {
			input: "}",
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsRightBrace(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsRightBrace(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsColon(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		":": {
			input: ":",
			want:  true,
		},
		";": {
			input: ";",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsColon(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsColon(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsComma(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		",": {
			input: ",",
			want:  true,
		},
		".": {
			input: ".",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsComma(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsComma(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsLeftBracket(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"[": {
			input: "[",
			want:  true,
		},
		"]": {
			input: "]",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsLeftBracket(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsLeftBracket(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsRightBracket(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"[": {
			input: "[",
			want:  false,
		},
		"]": {
			input: "]",
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsRightBracket(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsRightBracket(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}
