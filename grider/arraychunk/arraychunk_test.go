package arrayChunk

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		data         []int
		subSliceSize int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[1,2,3,4],2 return [[1,2], [3,4]]",
			args: args{
				data:         []int{1, 2, 3, 4},
				subSliceSize: 2,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "[1,2,3,4,5],2 return [[1,2], [3,4], [5]]",
			args: args{
				data:         []int{1, 2, 3, 4, 5},
				subSliceSize: 2,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
				{5},
			},
		},
		{
			name: "[1,2,3,4,5],10 return [[1,2,3,4,5]]",
			args: args{
				data:         []int{1, 2, 3, 4, 5},
				subSliceSize: 10,
			},
			want: [][]int{
				{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.data, tt.args.subSliceSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
