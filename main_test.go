package main

import "testing"

func TestStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		charset string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test length",
			args: args{
				length:  10,
				charset: "abcdefghijklmnopqrstuvwxyz",
			},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(StringWithCharset(tt.args.length, tt.args.charset)); got != len(tt.want) {
				t.Errorf("length StringWithCharset() = %v, want %v", got, len(tt.want))
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.length); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
