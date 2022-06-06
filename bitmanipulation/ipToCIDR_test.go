package bitmanipulation

import (
	"fmt"
	"testing"
)

func Test_ipToCIDR(t *testing.T) {
	type args struct {
		ip string
		n  int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"happy",
			args{
				ip: "255.0.0.7",
				n:  10,
			},
			[]string{"255.0.0.7/32", "255.0.0.8/29", "255.0.0.16/32"},
		},
		{
			"0000",
			args{
				ip: "0.0.0.0",
				n:  1,
			},
			[]string{"255.0.0.7/32", "255.0.0.8/29", "255.0.0.16/32"},
		},
	}
	for _, tt := range tests {
		got := ipToCIDR(tt.args.ip, tt.args.n)
		for _, v := range got {
			fmt.Println(v)
		}

	}
}

func TestZero(t *testing.T) {
	var x int64
	x = 0
	v := x & -x
	fmt.Println(v)
}
