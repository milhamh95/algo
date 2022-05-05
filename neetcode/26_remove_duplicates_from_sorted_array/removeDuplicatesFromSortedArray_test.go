package remove_duplicates_from_sorted_array

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,1,2], return 2, nums = [1,2,_]",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "nums = [1,2,3,4], return 4, nums = [1,2,3,4]",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
