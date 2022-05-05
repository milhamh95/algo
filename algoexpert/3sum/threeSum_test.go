package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[] target 0, return [][]",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: [][]int{},
		},
		{
			name: "[12, 3, 1, 2, -6, 5, -8, 6] target 0, return [[-8, 2, 6],[-8, 3, 5],[-6, 1, 5]]",
			args: args{
				nums:   []int{12, 3, 1, 2, -6, 5, -8, 6},
				target: 0,
			},
			want: [][]int{
				{-8, 2, 6},
				{-8, 3, 5},
				{-6, 1, 5},
			},
		},
		{
			name: "nums = [-1,0,1,2,-1,-4] target 0, return [[-1,-1,2],[-1,0,1],[-1,0,1]]",
			args: args{
				nums:   []int{-1, 0, 1, 2, -1, -4},
				target: 0,
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
				{-1, 0, 1},
			},
		},
		{
			name: "nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 15] target 5, return []",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
				target: 5,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
