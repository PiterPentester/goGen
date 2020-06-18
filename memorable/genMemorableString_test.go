package memorable

import (
	"reflect"
	"testing"

	"github.com/PiterPentester/goGen/abracadabra"
)

func TestCurlToAddr(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fail - curl not exist addr",
			args: args{
				addr: "ngsgszdfgsa",
			},
			wantErr: true,
		},

		{
			name: "curl to randomtext.me",
			args: args{
				addr: "http://www.randomtext.me/api/gibberish",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CurlToAddr(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurlToAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetTextOut(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "fail - empty txt",
			args: args{
				txt: "",
			},
			wantErr: true,
		},

		{
			name: "good",
			args: args{
				txt: "text_out:{'Works!'}",
			},
			want:    "Works!",
			wantErr: false,
		},
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

func TestSplitByTag(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "fail - empty txt",
			args: args{
				txt: "",
			},
			wantErr: true,
		},

		{
			name: "good",
			args: args{
				txt: "works<p>both<p>method<p>war<p>",
			},
			want:    []string{"works", "both", "method", "war "},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitByTag(tt.args.txt)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitByTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("splitByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitBySpace(t *testing.T) {
	type args struct {
		res []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "fail - empty res",
			args: args{
				res: []string{},
			},
			wantErr: true,
		},
		{
			name: "good",
			args: args{
				res: []string{"test text", "split by", "space works"},
			},
			want:    []string{"test", "text", "split", "space", "works"},
			wantErr: false,
		},
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

func TestGenOfflineWordsAndGetRand(t *testing.T) {
	type args struct {
		numOfWords int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "fail - less than 3 words",
			args: args{
				numOfWords: 0,
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "good",
			args: args{
				numOfWords: 4,
			},
			want:    []string{"Works", "quality", "sound", "tested"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genOfflineWords(tt.args.numOfWords)
			if (err != nil) != tt.wantErr {
				t.Errorf("genOfflineWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("genOfflineWords() = %v, want %v", got, tt.want)
			}

			// Check two functions in one test to avoid code duplication
			get, e := GetRandWords(tt.args.numOfWords)
			if (e != nil) != tt.wantErr {
				t.Errorf("GetRandWords() error = %v, wantErr %v", e, tt.wantErr)
				return
			}
			if len(get) != len(tt.want) {
				t.Errorf("GetRandWords() = %v, want %v", get, tt.want)
			}
		})
	}
}

func TestGenMemorablePass(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name  string
		args  args
		check bool
	}{
		{
			name: "fail - bad complexity",
			args: args{
				words: []string{"simple", "insecure", "pass"},
			},
			check: false,
		},
		{
			name: "good complexity",
			args: args{
				words: []string{"Simple", "inSecure78*()", "pAss"},
			},
			check: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abracadabra.CheckStrength(GenMemorablePass(tt.args.words)); got != tt.check {
				t.Errorf("CheckStrength for GenMemorablePass() = %v, want %v", got, tt.check)
			}
		})
	}
}
