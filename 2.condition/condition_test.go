package main

import "testing"

// test function auto create
func Test_isOdd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// with key and value
		{name: "test case1", args: args{1}, want: true},
		// only value name ,args, want
		{"test case2", args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.n); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
