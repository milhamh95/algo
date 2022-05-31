package spiral_matrix_II

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "n = 3, otuput = [[1,2,3],[8,9,4],[7,6,5]]",
			args: args{
				n: 3,
			},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			name: "n = 1, output = [[1]]",
			args: args{
				n: 1,
			},
			want: [][]int{
				[]int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
