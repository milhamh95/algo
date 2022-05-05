package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,3,4], output = [24,12,8,6]",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "nums = [-1,1,0,-3,3], output = [0,0,9,0,0]",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productExceptSelf2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,3,4], output = [24,12,8,6]",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "nums = [-1,1,0,-3,3], output = [0,0,9,0,0]",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf1() = %v, want %v", got, tt.want)
			}
		})
	}
}
