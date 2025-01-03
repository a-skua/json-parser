package state

import (
	"testing"
)

func TestNumber_Next(t *testing.T) {
	tests := map[string]struct {
		state   Number
		input   string
		want    Number
		wantErr string
	}{
		"NumberStart: next sign": {
			state: NumberStart,
			input: "-",
			want:  NumberSign,
		},
		"NumberStart: next zero": {
			state: NumberStart,
			input: "0",
			want:  NumberZero,
		},
		"NumberStart: next digit": {
			state: NumberStart,
			input: "1",
			want:  NumberInteger,
		},
		"NumberSign: next digit": {
			state: NumberSign,
			input: "1",
			want:  NumberInteger,
		},
		"NumberSign: next zero": {
			state: NumberSign,
			input: "0",
			want:  NumberZero,
		},
		"NumberZero: next dot": {
			state: NumberZero,
			input: ".",
			want:  NumberFractionSymol,
		},
		"NumberZero: next 'e'": {
			state: NumberZero,
			input: "e",
			want:  NumberExponentSymbol,
		},
		"NumberZero: next 'E'": {
			state: NumberZero,
			input: "E",
			want:  NumberExponentSymbol,
		},
		"NumberZero: next non symbols": {
			state: NumberZero,
			input: "0",
			want:  NumberEnd,
		},
		"NumberInteger: next digit": {
			state: NumberInteger,
			input: "1",
			want:  NumberInteger,
		},
		"NumberInteger: next '.'": {
			state: NumberInteger,
			input: ".",
			want:  NumberFractionSymol,
		},
		"NumberInteger: next 'e'": {
			state: NumberInteger,
			input: "e",
			want:  NumberExponentSymbol,
		},
		"NumberInteger: next 'E'": {
			state: NumberInteger,
			input: "E",
			want:  NumberExponentSymbol,
		},
		"NumberInteger: next non digit": {
			state: NumberInteger,
			input: ",",
			want:  NumberEnd,
		},
		"NumberFractionSymol: next digit": {
			state: NumberFractionSymol,
			input: "1",
			want:  NumberFraction,
		},
		"NumberFractionSymol: next non digit": {
			state:   NumberFractionSymol,
			input:   ",",
			wantErr: "Expected digit after '.': ','",
		},
		"NumberFraction: next digit": {
			state: NumberFraction,
			input: "1",
			want:  NumberFraction,
		},
		"NumberFraction: next 'e'": {
			state: NumberFraction,
			input: "e",
			want:  NumberExponentSymbol,
		},
		"NumberFraction: next 'E'": {
			state: NumberFraction,
			input: "E",
			want:  NumberExponentSymbol,
		},
		"NumberFraction: non digit": {
			state: NumberFraction,
			input: ",",
			want:  NumberEnd,
		},
		"NumberExponentSymbol: next '-'": {
			state: NumberExponentSymbol,
			input: "-",
			want:  NumberExponentSign,
		},
		"NumberExponentSymbol: next '+'": {
			state: NumberExponentSymbol,
			input: "+",
			want:  NumberExponentSign,
		},
		"NumberExponentSymbol: next digit": {
			state: NumberExponentSymbol,
			input: "0",
			want:  NumberExponent,
		},
		"NumberExponentSymbol: next non digit": {
			state:   NumberExponentSymbol,
			input:   ",",
			wantErr: "Expected digit or sign after 'e' or 'E': ','",
		},
		"NumberExponentSign: next digit": {
			state: NumberExponentSign,
			input: "0",
			want:  NumberExponent,
		},
		"NumberExponentSign: next non digit": {
			state:   NumberExponentSign,
			input:   ",",
			wantErr: "Expected digit after sign: ','",
		},
		"NumberExponent: next digit": {
			state: NumberExponent,
			input: "0",
			want:  NumberExponent,
		},
		"NumberExponent: next non digit": {
			state: NumberExponent,
			input: ",",
			want:  NumberEnd,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.state.Next([]rune(tt.input)[0])
			if err != nil && err.Error() != tt.wantErr {
				t.Fatalf("Number.Next(%s) error: %v (want: %v)", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("Number.Next(%s) = %v (want: %v)", tt.input, got, tt.want)
			}
		})
	}
}

func TestNumber_IsEnd(t *testing.T) {
	tests := map[string]struct {
		state Number
		want  bool
	}{
		"NumberStart": {
			state: NumberStart,
			want:  false,
		},
		"NumberSign": {
			state: NumberSign,
			want:  false,
		},
		"NumberZero": {
			state: NumberZero,
			want:  false,
		},
		"NumberInteger": {
			state: NumberInteger,
			want:  false,
		},
		"NumberFractionSymol": {
			state: NumberFractionSymol,
			want:  false,
		},
		"NumberFraction": {
			state: NumberFraction,
			want:  false,
		},
		"NumberExponentSymbol": {
			state: NumberExponentSymbol,
			want:  false,
		},
		"NumberExponentSign": {
			state: NumberExponentSign,
			want:  false,
		},
		"NumberExponent": {
			state: NumberExponent,
			want:  false,
		},
		"NumberEnd": {
			state: NumberEnd,
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.state.IsEnd()
			if got != tt.want {
				t.Fatalf("Number.IsEnd() = %v (want: %v)", got, tt.want)
			}
		})
	}
}

func TestNumber_Valid(t *testing.T) {
	tests := map[string]struct {
		state Number
		want  bool
	}{
		"NumberStart": {
			state: NumberStart,
			want:  false,
		},
		"NumberSign": {
			state: NumberSign,
			want:  false,
		},
		"NumberZero": {
			state: NumberZero,
			want:  true,
		},
		"NumberInteger": {
			state: NumberInteger,
			want:  true,
		},
		"NumberFractionSymol": {
			state: NumberFractionSymol,
			want:  false,
		},
		"NumberFraction": {
			state: NumberFraction,
			want:  true,
		},
		"NumberExponentSymbol": {
			state: NumberExponentSymbol,
			want:  false,
		},
		"NumberExponentSign": {
			state: NumberExponentSign,
			want:  false,
		},
		"NumberExponent": {
			state: NumberExponent,
			want:  true,
		},
		"NumberEnd": {
			state: NumberEnd,
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.state.Valid()
			if got != tt.want {
				t.Fatalf("Number.Valid() = %v (want: %v)", got, tt.want)
			}
		})
	}
}
