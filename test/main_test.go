package codestrings_test

import (
	"io/ioutil"
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
			want: []string{"hello", "world", "it's a beautiful day", "it\\\\'s a beautiful day"},
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
			want: []string{"hello", "world", "it\\\\'s a beautiful day"},
		},
		{
			name: "multiple strings in either single or double quotes",
			args: args{
				source:           `var a = 'hello'; var b = "world"; var c = 'it\\'s a beautiful day'; var d = "it's a beautiful day"`,
				stringDelimiters: []string{"'", "\""},
			},
			want: []string{"hello", "world", "it\\\\'s a beautiful day", "it's a beautiful day"},
		},
		{
			name: "strings in an expression",
			args: args{
				source:           `var a = "hello" + "world"`,
				stringDelimiters: []string{"\""},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "strings in an expression searching for double and single quotes",
			args: args{
				source:           `var a = "hello" + "world" + "it's a beautiful day"`,
				stringDelimiters: []string{"\"", "'"},
			},
			want: []string{"hello", "world", "it's a beautiful day"},
		},
		{
			name: "strings in an expression with mixed quotes",
			args: args{
				source:           `var a = "hello" + 'world' + "it's a beautiful day"`,
				stringDelimiters: []string{"\"", "'"},
			},
			want: []string{"hello", "world", "it's a beautiful day"},
		},
		{
			name: "strings in an expression with escaped quotes",
			args: args{
				source:           `var a = "hello" + "world" + 'it\'s a beautiful day'`,
				stringDelimiters: []string{"\"", "'"},
			},
			want: []string{"hello", "world", "it\\'s a beautiful day"},
		},
		{
			name: "string in the middle of an expression with white space",
			args: args{
				source:           `const d = a + " " + b + "hello" + name;`,
				stringDelimiters: []string{"\""},
			},
			want: []string{"hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cs.ExtractStrings(tt.args.source, tt.args.stringDelimiters); !equal(got, tt.want) {
				t.Errorf("ExtractStrings() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkExtractStrings(b *testing.B) {
	// load example.js
	source, err := ioutil.ReadFile("example.js")
	if err != nil {
		b.Fatal(err)
	}
	delimiters := []string{"\"", "'", "`"}
	for i := 0; i < b.N; i++ {
		cs.ExtractStrings(string(source), delimiters)
	}
}
