package state

import (
	"testing"
)

func TestString_IsEnd(t *testing.T) {
	tests := map[string]struct {
		state String
		want  bool
	}{
		"StringStart": {
			state: StringStart,
			want:  false,
		},
		"StringEnd": {
			state: StringEnd,
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.state.IsEnd()
			if got != tt.want {
				t.Fatalf("String.IsEnd() = %v (want: %v)", got, tt.want)
			}
		})
	}
}

func TestString_Valid(t *testing.T) {
	tests := map[string]struct {
		state String
		want  bool
	}{
		"StringLastQuote": {
			state: StringLastQuote,
			want:  true,
		},
		"StringEnd": {
			state: StringEnd,
			want:  true,
		},
		"StringStart": {
			state: StringStart,
			want:  false,
		},
		"StringFirstQuote": {
			state: StringFirstQuote,
			want:  false,
		},
		"StringCodepoint": {
			state: StringCodepoint,
			want:  false,
		},
		"StringEscapeSymbol": {
			state: StringEscapeSymbol,
			want:  false,
		},
		"StringEscape": {
			state: StringEscape,
			want:  false,
		},
		"StringHexDigitSymbol": {
			state: StringHexDigitSymbol,
			want:  false,
		},
		"StringHexDigit1": {
			state: StringHexDigit1,
			want:  false,
		},
		"StringHexDigit2": {
			state: StringHexDigit2,
			want:  false,
		},
		"StringHexDigit3": {
			state: StringHexDigit3,
			want:  false,
		},
		"StringHexDigit4": {
			state: StringHexDigit4,
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.state.Valid()
			if got != tt.want {
				t.Fatalf("String.Valid() = %v (want: %v)", got, tt.want)
			}
		})
	}
}

func TestString_Next(t *testing.T) {
	tests := map[string]struct {
		state   String
		input   string
		want    String
		wantErr string
	}{
		"StringStart: next quote": {
			state: StringStart,
			input: "\"",
			want:  StringFirstQuote,
		},
		"StringStart: next non quote": {
			state:   StringStart,
			input:   "0",
			wantErr: "Expected '\"' at the start of a string: 0",
		},
		"StringFirstQuote: next codepoint": {
			state: StringFirstQuote,
			input: "a",
			want:  StringCodepoint,
		},
		"StringFirstQuote: next escape symbol": {
			state: StringFirstQuote,
			input: "\\",
			want:  StringEscapeSymbol,
		},
		"StringFirstQuote: next quote": {
			state: StringFirstQuote,
			input: "\"",
			want:  StringLastQuote,
		},
		"StringCodepoint: next codepoint": {
			state: StringCodepoint,
			input: "a",
			want:  StringCodepoint,
		},
		"StringCodepoint: next escape symbol": {
			state: StringCodepoint,
			input: "\\",
			want:  StringEscapeSymbol,
		},
		"StringCodepoint: next quote": {
			state: StringCodepoint,
			input: "\"",
			want:  StringLastQuote,
		},
		"StringEscapeSymbol: next quotation mark": {
			state: StringEscapeSymbol,
			input: "\"",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next reverse solidus": {
			state: StringEscapeSymbol,
			input: "\\",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next solidus": {
			state: StringEscapeSymbol,
			input: "/",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next backspace": {
			state: StringEscapeSymbol,
			input: "b",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next formfeed": {
			state: StringEscapeSymbol,
			input: "f",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next linefeed": {
			state: StringEscapeSymbol,
			input: "n",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next carriage return": {
			state: StringEscapeSymbol,
			input: "r",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next tab": {
			state: StringEscapeSymbol,
			input: "t",
			want:  StringEscape,
		},
		"StringEscapeSymbol: next 'u'": {
			state: StringEscapeSymbol,
			input: "u",
			want:  StringHexDigitSymbol,
		},
		"StringEscapeSymbol: next non escape codepoint": {
			state:   StringEscapeSymbol,
			input:   "a",
			wantErr: "Unexpected Escape String: 'a'",
		},
		"StringEscape: next codepoint": {
			state: StringEscape,
			input: "a",
			want:  StringCodepoint,
		},
		"StringEscape: next escape symbol": {
			state: StringEscape,
			input: "\\",
			want:  StringEscapeSymbol,
		},
		"StringEscape: next quote": {
			state: StringEscape,
			input: "\"",
			want:  StringLastQuote,
		},
		"StringHexDigitSymbol: next hex digit": {
			state: StringHexDigitSymbol,
			input: "0",
			want:  StringHexDigit1,
		},
		"StringHexDigitSymbol: next non hex digit": {
			state:   StringHexDigitSymbol,
			input:   "g",
			wantErr: "Unexpected unicode: 'g'",
		},
		"StringHexDigit1: next hex digit": {
			state: StringHexDigit1,
			input: "0",
			want:  StringHexDigit2,
		},
		"StringHexDigit1: next non hex digit": {
			state:   StringHexDigit1,
			input:   "g",
			wantErr: "Unexpected unicode: 'g'",
		},
		"StringHexDigit2: next hex digit": {
			state: StringHexDigit2,
			input: "0",
			want:  StringHexDigit3,
		},
		"StringHexDigit2: next non hex digit": {
			state:   StringHexDigit2,
			input:   "g",
			wantErr: "Unexpected unicode: 'g'",
		},
		"StringHexDigit3: next hex digit": {
			state: StringHexDigit3,
			input: "0",
			want:  StringHexDigit4,
		},
		"StringHexDigit3: next non hex digit": {
			state:   StringHexDigit3,
			input:   "g",
			wantErr: "Unexpected unicode: 'g'",
		},
		"StringHexDigit4: next hex digit": {
			state: StringHexDigit4,
			input: "0",
			want:  StringCodepoint,
		},
		"StringHexDigit4: next non hex digit": {
			state: StringHexDigit4,
			input: "g",
			want:  StringCodepoint,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.state.Next(rune(tt.input[0]))
			if err != nil && err.Error() != tt.wantErr {
				t.Fatalf("String.Next(%s) error: %v (want: %v)", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("String.Next(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}
