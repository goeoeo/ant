package test

import (
	"fmt"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

type ValidItem struct {
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64
	Int   int

	UInt8  uint8
	UInt16 uint16
	UInt32 uint32
	UInt64 uint64
	UInt   uint

	Float32 float32
	Float64 float64

	Bool bool

	String string
	Byte   []byte

	//兼容proto
	WpDouble *wrapperspb.DoubleValue
	WpFloat  *wrapperspb.FloatValue
	WpInt64  *wrapperspb.Int64Value
	WpUInt64 *wrapperspb.UInt64Value
	WpInt32  *wrapperspb.Int32Value
	WpUInt32 *wrapperspb.UInt32Value
	WpBool   *wrapperspb.BoolValue
	WpString *wrapperspb.StringValue
	WpBytes  *wrapperspb.BytesValue
}

func eqErr(a, b error) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil {
		return a.Error() == b.Error()
	}

	return false
}

//为空测试
func Test_Required(t *testing.T) {
	funcName := "Required"
	errMsg := "%s"
	errMsg += defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"Int8", ValidItem{Int8: 0}, nil},
		{"Int8", ValidItem{Int8: 1}, nil},
		{"Int16", ValidItem{Int16: 0}, nil},
		{"Int16", ValidItem{Int16: 1}, nil},
		{"Int32", ValidItem{Int32: 0}, nil},
		{"Int32", ValidItem{Int32: 1}, nil},
		{"Int64", ValidItem{Int64: 0}, nil},
		{"Int64", ValidItem{Int64: 1}, nil},
		{"Int", ValidItem{Int: 0}, nil},
		{"Int", ValidItem{Int: 1}, nil},
		{"UInt8", ValidItem{UInt8: 0}, nil},
		{"UInt8", ValidItem{UInt8: 1}, nil},
		{"UInt16", ValidItem{UInt16: 0}, nil},
		{"UInt16", ValidItem{UInt16: 1}, nil},
		{"UInt32", ValidItem{UInt32: 0}, nil},
		{"UInt32", ValidItem{UInt32: 1}, nil},
		{"UInt64", ValidItem{UInt64: 0}, nil},
		{"UInt64", ValidItem{UInt64: 1}, nil},
		{"UInt", ValidItem{UInt: 0}, nil},
		{"UInt", ValidItem{UInt: 1}, nil},
		{"Float32", ValidItem{Float32: 0.0}, nil},
		{"Float32", ValidItem{Float32: 1.1}, nil},
		{"Float64", ValidItem{Float64: 0.0}, nil},
		{"Float64", ValidItem{Float64: 1.1}, nil},

		{"Bool", ValidItem{Bool: false}, nil},
		{"Bool", ValidItem{Bool: true}, nil},

		{"String", ValidItem{String: ""}, fmt.Errorf(errMsg, "String")},
		{"String", ValidItem{String: "1"}, nil},
		{"Byte", ValidItem{Byte: []byte("")}, fmt.Errorf(errMsg, "Byte")},
		{"Byte", ValidItem{Byte: []byte("1")}, nil},

		{"WpDouble", ValidItem{WpDouble: nil}, fmt.Errorf(errMsg, "WpDouble")},
		{"WpDouble", ValidItem{WpDouble: new(wrapperspb.DoubleValue)}, nil},
		{"WpFloat", ValidItem{WpFloat: nil}, fmt.Errorf(errMsg, "WpFloat")},
		{"WpFloat", ValidItem{WpFloat: new(wrapperspb.FloatValue)}, nil},
		{"WpInt32", ValidItem{WpInt32: nil}, fmt.Errorf(errMsg, "WpInt32")},
		{"WpInt32", ValidItem{WpInt32: new(wrapperspb.Int32Value)}, nil},

		{"WpInt64", ValidItem{WpInt64: nil}, fmt.Errorf(errMsg, "WpInt64")},
		{"WpInt64", ValidItem{WpInt64: new(wrapperspb.Int64Value)}, nil},

		{"WpUInt32", ValidItem{WpUInt32: nil}, fmt.Errorf(errMsg, "WpUInt32")},
		{"WpUInt32", ValidItem{WpUInt32: new(wrapperspb.UInt32Value)}, nil},

		{"WpUInt64", ValidItem{WpUInt64: nil}, fmt.Errorf(errMsg, "WpUInt64")},
		{"WpUInt64", ValidItem{WpUInt64: new(wrapperspb.UInt64Value)}, nil},

		{"WpBool", ValidItem{WpBool: nil}, fmt.Errorf(errMsg, "WpBool")},
		{"WpBool", ValidItem{WpBool: new(wrapperspb.BoolValue)}, nil},

		{"WpString", ValidItem{WpString: nil}, fmt.Errorf(errMsg, "WpString")},
		{"WpString", ValidItem{WpString: new(wrapperspb.StringValue)}, fmt.Errorf(errMsg, "WpString")},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "sss"}}, nil},

		{"WpBytes", ValidItem{WpBytes: nil}, fmt.Errorf(errMsg, "WpBytes")},
		{"WpBytes", ValidItem{WpBytes: new(wrapperspb.BytesValue)}, fmt.Errorf(errMsg, "WpBytes")},
		{"WpBytes", ValidItem{WpBytes: &wrapperspb.BytesValue{Value: []byte("ss")}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: `valid:"` + funcName + `"`}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

//严格模式为空测试
func Test_RequiredStrict(t *testing.T) {
	funcName := "Required"
	errMsg := "%s"
	errMsg += defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"Int8", ValidItem{Int8: 0}, fmt.Errorf(errMsg, "Int8")},
		{"Int8", ValidItem{Int8: 1}, nil},

		{"Int16", ValidItem{Int16: 0}, fmt.Errorf(errMsg, "Int16")},
		{"Int16", ValidItem{Int16: 1}, nil},

		{"Int32", ValidItem{Int32: 0}, fmt.Errorf(errMsg, "Int32")},
		{"Int32", ValidItem{Int32: 1}, nil},
		{"Int64", ValidItem{Int64: 0}, fmt.Errorf(errMsg, "Int64")},
		{"Int64", ValidItem{Int64: 1}, nil},
		{"Int", ValidItem{Int: 0}, fmt.Errorf(errMsg, "Int")},
		{"Int", ValidItem{Int: 1}, nil},
		{"UInt8", ValidItem{UInt8: 0}, fmt.Errorf(errMsg, "UInt8")},
		{"UInt8", ValidItem{UInt8: 1}, nil},
		{"UInt16", ValidItem{UInt16: 0}, fmt.Errorf(errMsg, "UInt16")},
		{"UInt16", ValidItem{UInt16: 1}, nil},
		{"UInt32", ValidItem{UInt32: 0}, fmt.Errorf(errMsg, "UInt32")},
		{"UInt32", ValidItem{UInt32: 1}, nil},
		{"UInt64", ValidItem{UInt64: 0}, fmt.Errorf(errMsg, "UInt64")},
		{"UInt64", ValidItem{UInt64: 1}, nil},
		{"UInt", ValidItem{UInt: 0}, fmt.Errorf(errMsg, "UInt")},
		{"UInt", ValidItem{UInt: 1}, nil},
		{"Float32", ValidItem{Float32: 0.0}, fmt.Errorf(errMsg, "Float32")},
		{"Float32", ValidItem{Float32: 1.1}, nil},
		{"Float64", ValidItem{Float64: 0.0}, fmt.Errorf(errMsg, "Float64")},
		{"Float64", ValidItem{Float64: 1.1}, nil},

		{"Bool", ValidItem{Bool: false}, fmt.Errorf(errMsg, "Bool")},
		{"Bool", ValidItem{Bool: true}, nil},

		{"String", ValidItem{String: ""}, fmt.Errorf(errMsg, "String")},
		{"String", ValidItem{String: "1"}, nil},
		{"Byte", ValidItem{Byte: []byte("")}, fmt.Errorf(errMsg, "Byte")},
		{"Byte", ValidItem{Byte: []byte("1")}, nil},

		{"WpDouble", ValidItem{WpDouble: nil}, fmt.Errorf(errMsg, "WpDouble")},
		{"WpDouble", ValidItem{WpDouble: new(wrapperspb.DoubleValue)}, fmt.Errorf(errMsg, "WpDouble")},
		{"WpDouble", ValidItem{WpDouble: &wrapperspb.DoubleValue{Value: 1.1}}, nil},

		{"WpFloat", ValidItem{WpFloat: nil}, fmt.Errorf(errMsg, "WpFloat")},
		{"WpFloat", ValidItem{WpFloat: new(wrapperspb.FloatValue)}, fmt.Errorf(errMsg, "WpFloat")},
		{"WpFloat", ValidItem{WpFloat: &wrapperspb.FloatValue{Value: 1.1}}, nil},

		{"WpInt32", ValidItem{WpInt32: nil}, fmt.Errorf(errMsg, "WpInt32")},
		{"WpInt32", ValidItem{WpInt32: new(wrapperspb.Int32Value)}, fmt.Errorf(errMsg, "WpInt32")},
		{"WpInt32", ValidItem{WpInt32: &wrapperspb.Int32Value{Value: 1}}, nil},

		{"WpInt64", ValidItem{WpInt64: nil}, fmt.Errorf(errMsg, "WpInt64")},
		{"WpInt64", ValidItem{WpInt64: new(wrapperspb.Int64Value)}, fmt.Errorf(errMsg, "WpInt64")},
		{"WpInt64", ValidItem{WpInt64: &wrapperspb.Int64Value{Value: 1}}, nil},

		{"WpUInt32", ValidItem{WpUInt32: nil}, fmt.Errorf(errMsg, "WpUInt32")},
		{"WpUInt32", ValidItem{WpUInt32: new(wrapperspb.UInt32Value)}, fmt.Errorf(errMsg, "WpUInt32")},
		{"WpUInt32", ValidItem{WpUInt32: &wrapperspb.UInt32Value{Value: 1}}, nil},

		{"WpUInt64", ValidItem{WpUInt64: nil}, fmt.Errorf(errMsg, "WpUInt64")},
		{"WpUInt64", ValidItem{WpUInt64: new(wrapperspb.UInt64Value)}, fmt.Errorf(errMsg, "WpUInt64")},
		{"WpUInt64", ValidItem{WpUInt64: &wrapperspb.UInt64Value{Value: 1}}, nil},

		{"WpBool", ValidItem{WpBool: nil}, fmt.Errorf(errMsg, "WpBool")},
		{"WpBool", ValidItem{WpBool: new(wrapperspb.BoolValue)}, fmt.Errorf(errMsg, "WpBool")},
		{"WpBool", ValidItem{WpBool: &wrapperspb.BoolValue{Value: true}}, nil},

		{"WpString", ValidItem{WpString: nil}, fmt.Errorf(errMsg, "WpString")},
		{"WpString", ValidItem{WpString: new(wrapperspb.StringValue)}, fmt.Errorf(errMsg, "WpString")},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "sss"}}, nil},

		{"WpBytes", ValidItem{WpBytes: nil}, fmt.Errorf(errMsg, "WpBytes")},
		{"WpBytes", ValidItem{WpBytes: new(wrapperspb.BytesValue)}, fmt.Errorf(errMsg, "WpBytes")},
		{"WpBytes", ValidItem{WpBytes: &wrapperspb.BytesValue{Value: []byte("ss")}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: `valid:"` + funcName + `(true)"`}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_Max(t *testing.T) {
	funcName := "Max"
	maxNum := "10"
	validFunName := `valid:"` + funcName + `(` + maxNum + `)"`
	errMsg := "%s"
	errMsg += defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"Int", ValidItem{Int: 11}, fmt.Errorf(errMsg, "Int", maxNum)},
		{"Int", ValidItem{Int: 0}, nil},

		{"WpInt64", ValidItem{WpInt64: nil}, nil},
		{"WpInt64", ValidItem{WpInt64: new(wrapperspb.Int64Value)}, nil},
		{"WpInt64", ValidItem{WpInt64: &wrapperspb.Int64Value{Value: 11}}, fmt.Errorf(errMsg, "WpInt64", maxNum)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_Min(t *testing.T) {
	funcName := "Min"
	maxNum := "10"
	validFunName := `valid:"` + funcName + `(` + maxNum + `)"`
	errMsg := "%s"
	errMsg += defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"Int", ValidItem{Int: 11}, nil},
		{"Int", ValidItem{Int: 0}, fmt.Errorf(errMsg, "Int", maxNum)},

		{"WpInt64", ValidItem{WpInt64: nil}, nil},
		{"WpInt64", ValidItem{WpInt64: new(wrapperspb.Int64Value)}, fmt.Errorf(errMsg, "WpInt64", maxNum)},
		{"WpInt64", ValidItem{WpInt64: &wrapperspb.Int64Value{Value: 11}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_Range(t *testing.T) {
	funcName := "Range"
	maxNum := "10"
	minNum := "5"
	validFunName := `valid:"` + fmt.Sprintf("%s(%s,%s)", funcName, minNum, maxNum) + `"`
	errMsg := "%s"
	errMsg += defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"Int", ValidItem{Int: 6}, nil},
		{"Int", ValidItem{Int: 11}, fmt.Errorf(errMsg, "Int", minNum, maxNum)},
		{"Int", ValidItem{Int: 0}, fmt.Errorf(errMsg, "Int", minNum, maxNum)},

		{"WpInt64", ValidItem{WpInt64: nil}, nil},
		{"WpInt64", ValidItem{WpInt64: &wrapperspb.Int64Value{Value: 6}}, nil},
		{"WpInt64", ValidItem{WpInt64: new(wrapperspb.Int64Value)}, fmt.Errorf(errMsg, "WpInt64", minNum, maxNum)},
		{"WpInt64", ValidItem{WpInt64: &wrapperspb.Int64Value{Value: 11}}, fmt.Errorf(errMsg, "WpInt64", minNum, maxNum)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_MinSize(t *testing.T) {
	funcName := "MinSize"
	minNum := "5"
	validFunName := `valid:"` + fmt.Sprintf("%s(%s)", funcName, minNum) + `"`
	errMsg := "%s"
	errMsg += defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"String", ValidItem{String: "w我爱中国"}, nil},
		{"String", ValidItem{String: ""}, fmt.Errorf(errMsg, "String", minNum)},

		{"WpString", ValidItem{WpInt64: nil}, nil},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱中国"}}, nil},
		{"WpString", ValidItem{WpString: new(wrapperspb.StringValue)}, fmt.Errorf(errMsg, "WpString", minNum)},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱中"}}, fmt.Errorf(errMsg, "WpString", minNum)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_MaxSize(t *testing.T) {
	funcName := "MaxSize"
	minNum := "5"
	validFunName := `valid:"` + fmt.Sprintf("%s(%s)", funcName, minNum) + `"`
	errMsg := "%s" + defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"String", ValidItem{String: ""}, nil},
		{"String", ValidItem{String: "w我爱中国aa"}, fmt.Errorf(errMsg, "String", minNum)},

		{"WpString", ValidItem{WpInt64: nil}, nil},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱中国11"}}, fmt.Errorf(errMsg, "WpString", minNum)},
		{"WpString", ValidItem{WpString: new(wrapperspb.StringValue)}, nil},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱中"}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_Length(t *testing.T) {
	funcName := "Length"
	minNum := "5"
	validFunName := `valid:"` + fmt.Sprintf("%s(%s)", funcName, minNum) + `"`
	errMsg := "%s" + defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"String", ValidItem{String: "w我爱中国"}, nil},
		{"String", ValidItem{String: "w我爱中国aa"}, fmt.Errorf(errMsg, "String", minNum)},

		{"WpString", ValidItem{WpInt64: nil}, nil},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱中国11"}}, fmt.Errorf(errMsg, "WpString", minNum)},
		{"WpString", ValidItem{WpString: new(wrapperspb.StringValue)}, fmt.Errorf(errMsg, "WpString", minNum)},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱中1"}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}

func Test_Alpha(t *testing.T) {
	funcName := "Alpha"
	validFunName := `valid:"` + fmt.Sprintf("%s", funcName) + `"`
	errMsg := "%s" + defaultValidationConfig.GetMessageTmpls(funcName)

	tests := []struct {
		name string
		args ValidItem
		want error
	}{
		{"String", ValidItem{String: "afaZ"}, nil},
		{"String", ValidItem{String: "afaZ1"}, fmt.Errorf(errMsg, "String")},

		{"WpString", ValidItem{WpInt64: nil}, nil},
		{"WpString", ValidItem{WpString: new(wrapperspb.StringValue)}, nil},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "aadfAA"}}, nil},
		{"WpString", ValidItem{WpString: &wrapperspb.StringValue{Value: "w我爱aa中国11"}}, fmt.Errorf(errMsg, "WpString")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidate().SetFieldTag(map[string]string{tt.name: validFunName}).Valid(tt.args); !eqErr(got, tt.want) {
				t.Errorf("%s() = %v, want %v", funcName, got, tt.want)
			}
		})
	}
}
