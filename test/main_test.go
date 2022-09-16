package codestrings_test

import (
	"testing"

	cs "github.com/alexferrari88/codestrings"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestExtractStrings(t *testing.T) {
	type args struct {
		source           string
		stringDelimiters []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				source:           "",
				stringDelimiters: []string{},
			},
			want: []string{},
		},
		{
			name: "empty string delimiters",
			args: args{
				source:           `var a = "hello"`,
				stringDelimiters: []string{},
			},
			want: []string{"hello"},
		},
		{
			name: "single string",
			args: args{
				source:           `var a = "hello"`,
				stringDelimiters: []string{"\""},
			},
			want: []string{"hello"},
		},
		{
			name: "multiple strings",
			args: args{
				source:           `var a = "hello"; var b = 'world'`,
				stringDelimiters: []string{"\"", "'"},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "multiple strings with escaped quotes",
			args: args{
				source:           `var a = "hello"; var b = 'world'; var c = "it's a beautiful day"`,
				stringDelimiters: []string{"\"", "'"},
			},
			want: []string{"hello", "world", "it's a beautiful day"},
		},
		{
			name: "multiple strings with escaped quotes and escaped backslashes",
			args: args{
				source:           `var a = "hello"; var b = 'world'; var c = "it's a beautiful day"; var d = "it\\'s a beautiful day"`,
				stringDelimiters: []string{"\"", "'"},
			},
			want: []string{"hello", "world", "it's a beautiful day", "it\\'s a beautiful day"},
		},
		{
			name: "single string in single quotes",
			args: args{
				source:           `var a = 'hello'`,
				stringDelimiters: []string{"'"},
			},
			want: []string{"hello"},
		},
		{
			name: "multiple strings in single quotes",
			args: args{
				source:           `var a = 'hello'; var b = 'world'`,
				stringDelimiters: []string{"'"},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "multiple strings in single quotes with escaped quotes",
			args: args{
				source:           `var a = 'hello'; var b = 'world'; var c = 'it\\'s a beautiful day'`,
				stringDelimiters: []string{"'"},
			},
			want: []string{"hello", "world", "it\\'s a beautiful day"},
		},
		{
			name: "multiple strings in either single or double quotes",
			args: args{
				source:           `var a = 'hello'; var b = "world"; var c = 'it\\'s a beautiful day'; var d = "it's a beautiful day"`,
				stringDelimiters: []string{"'", "\""},
			},
			want: []string{"hello", "world", "it\\'s a beautiful day", "it's a beautiful day"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cs.ExtractStrings(tt.args.source, tt.args.stringDelimiters); !equal(got, tt.want) {
				t.Errorf("ExtractStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
