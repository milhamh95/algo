package matrix

import (
	"reflect"
	"testing"
)

func TestMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "value is 3",
			args: args{
				n: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "value is 4",
			args: args{
				n: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Matrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
