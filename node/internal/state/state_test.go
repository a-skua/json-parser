package state

import (
	"testing"
)

func TestArray_Next(t *testing.T) {
	tests := map[string]struct {
		input Array
		want  Array
	}{
		"ArrayValue": {
			input: ArrayValue,
			want:  ArraySeparator,
		},
		"ArraySeparator": {
			input: ArraySeparator,
			want:  ArrayValue,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.Next()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestArray_IsValue(t *testing.T) {
	tests := map[string]struct {
		input Array
		want  bool
	}{
		"ArrayValue": {
			input: ArrayValue,
			want:  true,
		},
		"ArraySeparator": {
			input: ArraySeparator,
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.IsValue()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestArray_IsSeparator(t *testing.T) {
	tests := map[string]struct {
		input Array
		want  bool
	}{
		"ArrayValue": {
			input: ArrayValue,
			want:  false,
		},
		"ArraySeparator": {
			input: ArraySeparator,
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.IsSeparator()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestObject_Next(t *testing.T) {
	tests := map[string]struct {
		input Object
		want  Object
	}{
		"ObjectKey": {
			input: ObjectKey,
			want:  ObjectColon,
		},
		"ObjectColon": {
			input: ObjectColon,
			want:  ObjectValue,
		},
		"ObjectValue": {
			input: ObjectValue,
			want:  ObjectSeparator,
		},
		"ObjectSeparator": {
			input: ObjectSeparator,
			want:  ObjectKey,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.Next()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestObject_IsKey(t *testing.T) {
	tests := map[string]struct {
		input Object
		want  bool
	}{
		"ObjectKey": {
			input: ObjectKey,
			want:  true,
		},
		"ObjectColon": {
			input: ObjectColon,
			want:  false,
		},
		"ObjectValue": {
			input: ObjectValue,
			want:  false,
		},
		"ObjectSeparator": {
			input: ObjectSeparator,
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.IsKey()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestObject_IsColon(t *testing.T) {
	tests := map[string]struct {
		input Object
		want  bool
	}{
		"ObjectKey": {
			input: ObjectKey,
			want:  false,
		},
		"ObjectColon": {
			input: ObjectColon,
			want:  true,
		},
		"ObjectValue": {
			input: ObjectValue,
			want:  false,
		},
		"ObjectSeparator": {
			input: ObjectSeparator,
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.IsColon()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestObject_IsValue(t *testing.T) {
	tests := map[string]struct {
		input Object
		want  bool
	}{
		"ObjectKey": {
			input: ObjectKey,
			want:  false,
		},
		"ObjectColon": {
			input: ObjectColon,
			want:  false,
		},
		"ObjectValue": {
			input: ObjectValue,
			want:  true,
		},
		"ObjectSeparator": {
			input: ObjectSeparator,
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.IsValue()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestObject_IsSeparator(t *testing.T) {
	tests := map[string]struct {
		input Object
		want  bool
	}{
		"ObjectKey": {
			input: ObjectKey,
			want:  false,
		},
		"ObjectColon": {
			input: ObjectColon,
			want:  false,
		},
		"ObjectValue": {
			input: ObjectValue,
			want:  false,
		},
		"ObjectSeparator": {
			input: ObjectSeparator,
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.input.IsSeparator()
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
