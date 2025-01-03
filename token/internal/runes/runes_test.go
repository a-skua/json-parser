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

func TestIsHex(t *testing.T) {
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
			want:  true,
		},
		"f": {
			input: "f",
			want:  true,
		},
		"A": {
			input: "A",
			want:  true,
		},
		"F": {
			input: "F",
			want:  true,
		},
		"g": {
			input: "g",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsHex(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("IsHex(%s) = %v (want: %v)", tt.input, got, tt.want)
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

func TestMaybeTrue(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"true": {
			input: "true",
			want:  true,
		},
		"false": {
			input: "false",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := MaybeTrue([]rune(tt.input))
			if got != tt.want {
				t.Fatalf("MaybeTrue(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaybeFalse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"true": {
			input: "true",
			want:  false,
		},
		"false": {
			input: "false",
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := MaybeFalse([]rune(tt.input))
			if got != tt.want {
				t.Fatalf("MaybeFalse(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaybeNull(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"null": {
			input: "null",
			want:  true,
		},
		"false": {
			input: "false",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := MaybeNull([]rune(tt.input))
			if got != tt.want {
				t.Fatalf("MaybeNull(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaybeString(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"string": {
			input: "\"\"",
			want:  true,
		},
		"false": {
			input: "false",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := MaybeString([]rune(tt.input))
			if got != tt.want {
				t.Fatalf("MaybeString(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaybeNumber(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"digit": {
			input: "0",
			want:  true,
		},
		"negative": {
			input: "-",
			want:  true,
		},
		"letter": {
			input: "a",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := MaybeNumber(rune(tt.input[0]))
			if got != tt.want {
				t.Fatalf("MaybeNumber(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}
