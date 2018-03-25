// Copyright (c) 2018 hIMEI <himei@tuta.io>
package goluhn

import (
	"reflect"
	"testing"
)

var err error

func Test_errFatal(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{"err", args{err}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errFatal(tt.args.err)
		})
	}
}

func Test_chod(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"evennum", args{11}, false},
		{"odd", args{12}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chod(tt.args.num); got != tt.want {
				t.Errorf("chod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_powt(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"p", args{2}, 100},
		{"pp", args{4}, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powt(tt.args.p); got != tt.want {
				t.Errorf("powt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randt(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randt(tt.args.num); got != tt.want {
				t.Errorf("randt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aToi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"str", args{"1023"}, 1023},
		{"erstr", args{"10"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aToi(tt.args.str); got != tt.want {
				t.Errorf("aToi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base10(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ten", args{10}, true},
		{"tenten", args{100}, true},
		{"notten", args{10043}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base10(tt.args.num); got != tt.want {
				t.Errorf("base10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base9(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"num", args{6}, 3},
		{"num2", args{2}, 4},
		{"num3", args{4}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base9(tt.args.num); got != tt.want {
				t.Errorf("base9() = %v, want %v", got, tt.want)
			}
		})
	}
}

// test cases form http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers#Go
func TestCheckLuhn(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"w1", args{49927398716}, true},
		{"w2", args{49927398717}, false},
		{"w3", args{1234567812345678}, false},
		{"w4", args{1234567812345670}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLuhn(tt.args.num); got != tt.want {
				t.Errorf("CheckLuhn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuhnInRange(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"trange", args{100000, 100100}, []int{
			100008,
			100016,
			100024,
			100032,
			100040,
			100057,
			100065,
			100073,
			100081,
			100099,
		}},
		{"prange", args{49927398716, 49927398717}, []int{
			49927398716,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuhnInRange(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LuhnInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuhnByLen(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuhnByLen(tt.args.len); got != tt.want {
				t.Errorf("LuhnByLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
