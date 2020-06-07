package abracadabra

import (
	"regexp"
	"testing"
	"unicode"
)

// checkStrength(password) - test our password complexity.
// TRUE if contains UPPERCASE & lowercase & special symbols.
// Else => FALSE
func checkStrength(pass string) bool {
	isDigit := false
	isUpper := false
	isLower := false
	isSymbol := false

	if len(pass) < 8 {
		return false
	}

	for _, c := range pass {
		if unicode.IsDigit(c) {
			isDigit = true
		}
		if unicode.IsUpper(c) {
			isUpper = true
		}
		if unicode.IsLower(c) {
			isLower = true
		}
	}

	reg, err := regexp.Compile("[!^a-zA-Z0-9]+")
	if err != nil {
		return false
	}

	processedString := reg.ReplaceAllString(pass, "")
	if processedString != "" {
		isSymbol = true
	}

	if isDigit && isUpper && isLower && isSymbol {
		return true
	}
	return false
}

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
			name: "fail length < 8",
			args: args{
				length:  7,
				Charset: Charset,
			},
			check: false,
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
			if cmp := checkStrength(StringWithCharset(tt.args.length, Charset)); cmp != tt.check {
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
			name: "fail length < 8",
			args: args{
				length: 7,
			},
			check: false,
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
			if cmp := checkStrength(String(tt.args.length)); cmp != tt.check {
				t.Errorf("Check strength StringWithCharset() = %v, want %v", cmp, tt.check)
			}
		})
	}
}
