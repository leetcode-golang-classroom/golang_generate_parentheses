package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	n := 3
	for idx := 0; idx < b.N; idx++ {
		generateParenthesis(n)
	}
}
func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "n = 3",
			args: args{n: 3},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n = 1",
			args: args{n: 1},
			want: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
