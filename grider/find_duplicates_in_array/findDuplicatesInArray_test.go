package find_duplicates_in_array

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,1,2,2,2,3,3,3,3,3], return [1,2,3]",
			args: args{
				nums: []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "[3,6,1,7,1,5,1,3,9], return [1,3]",
			args: args{
				nums: []int{3, 6, 1, 7, 1, 5, 1, 3, 9},
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
