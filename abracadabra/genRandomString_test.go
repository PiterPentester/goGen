package abracadabra

import (
	"testing"
	"unicode"
)

func checkStrength(pass string) bool {
	isDigit := false
	for _, c := range pass {
		if unicode.IsDigit(c) {
			isDigit = true
		}
	}
	return isDigit
}

func TestStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		Charset string
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
				Charset: Charset,
			},
			want: "1234567890",
		},

		{
			name: "test length",
			args: args{
				length:  20,
				Charset: Charset,
			},
			want: "12345678901234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(StringWithCharset(tt.args.length, tt.args.Charset)); got != len(tt.want) {
				t.Errorf("StringWithCharset() = %v, want %v", got, tt.want)
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
		{
			name: "test length",
			args: args{
				length: 10,
			},
			want: "1234567890",
		},

		{
			name: "test length",
			args: args{
				length: 20,
			},
			want: "12345678901234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(String(tt.args.length)); got != len(tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
