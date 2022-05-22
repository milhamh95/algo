package sum_multiplier

import "testing"

func Test_sumMultliplier(t *testing.T) {
	type args struct {
		nums      []int
		minTarget int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: " [2, 5, 6, -6, 16, 2, 3, 6, 5, 3] return true",
			args: args{
				nums:      []int{2, 5, 6, -6, 16, 2, 3, 6, 5, 3},
				minTarget: 96,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumMultliplier(tt.args.nums, tt.args.minTarget); got != tt.want {
				t.Errorf("sumMultliplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
