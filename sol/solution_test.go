package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	for idx := 0; idx < b.N; idx++ {
		solve(board)
	}
}
func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Example1",
			args: args{board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "Example2",
			args: args{board: [][]byte{
				{'X'},
			}},
			want: [][]byte{
				{'X'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if solve(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("pacificAtlantic() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
