package test

import (
	"github.com/phpdi/ant/util"
	"testing"
)

func TestMd5(t *testing.T) {
	str := util.Md5([]byte("123456"))
	if str != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("md5 fail")
	}
}

func TestKeep0Add(t *testing.T) {
	s := "001"
	ns := util.Keep0Add(s, 2)
	if ns != "003" {
		t.Error("fail")
	}
}

func Test_snakeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ls",
			args: args{s: "XxYy"},
			want: "xx_yy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.snakeString(tt.args.s); got != tt.want {
				t.Errorf("snakeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
