package validation

import "testing"

func TestNumericDot(t *testing.T) {
	type args struct {
		validValue interface{}
		params     []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"0123456789.02", []string{}}, true},
		{"", args{"0123456789.afadf", []string{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumericDot(tt.args.validValue, tt.args.params...); got != tt.want {
				t.Errorf("NumericDot() = %v, want %v", got, tt.want)
			}
		})
	}
}
