package memorable

import (
	"reflect"
	"testing"
)

func TestCurlToAddr(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CurlToAddr(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurlToAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CurlToAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTextOut(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTextOut(tt.args.txt)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTextOut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTextOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByTag(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByTag(tt.args.txt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitBySpace(t *testing.T) {
	type args struct {
		res []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitBySpace(tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitBySpace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitBySpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOutput(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOutput()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandWords(t *testing.T) {
	type args struct {
		numOfWords int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandWords(tt.args.numOfWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenMemorablePass(t *testing.T) {
	type args struct {
		words []string
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
			if got := GenMemorablePass(tt.args.words); got != tt.want {
				t.Errorf("GenMemorablePass() = %v, want %v", got, tt.want)
			}
		})
	}
}
