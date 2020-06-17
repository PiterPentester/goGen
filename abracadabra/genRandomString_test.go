package abracadabra

import (
	"testing"
)

func TestStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		Charset string
	}
	tests := []struct {
		name  string
		args  args
		check bool
	}{
		{
			name: "length < 8 => len = 16",
			args: args{
				length:  7,
				Charset: Charset,
			},
			check: true,
		},
		{
			name: "test length 10",
			args: args{
				length:  10,
				Charset: Charset,
			},
			check: true,
		},

		{
			name: "test length 20",
			args: args{
				length:  20,
				Charset: Charset,
			},
			check: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if cmp := CheckStrength(StringWithCharset(tt.args.length, Charset)); cmp != tt.check {
				t.Errorf("Check strength StringWithCharset() = %v, want %v", cmp, tt.check)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name  string
		args  args
		check bool
	}{
		{
			name: "length < 8  => 16",
			args: args{
				length: 7,
			},
			check: true,
		},
		{
			name: "test length 10",
			args: args{
				length: 10,
			},
			check: true,
		},

		{
			name: "test length 20",
			args: args{
				length: 20,
			},
			check: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if cmp := CheckStrength(String(tt.args.length)); cmp != tt.check {
				t.Errorf("Check strength StringWithCharset() = %v, want %v", cmp, tt.check)
			}
		})
	}
}
