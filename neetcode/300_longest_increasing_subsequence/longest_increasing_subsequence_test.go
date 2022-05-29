package longest_increasing_subsequence

import "testing"

func Test_lengthOfLIS1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,1,7] return 3",
			args: args{
				nums: []int{2, 3, 1, 7},
			},
			want: 3,
		},
		{
			name: "[2,3,5,7] return 4",
			args: args{
				nums: []int{2, 3, 5, 7},
			},
			want: 4,
		},
		{
			name: "[7,7,7,7,7,7,7] return 1",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "[0,1,3,2,3] return 4",
			args: args{
				nums: []int{0, 1, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "[0,1,0] return 2",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS1(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,1,7] return 3",
			args: args{
				nums: []int{2, 3, 1, 7},
			},
			want: 3,
		},
		{
			name: "[2,3,5,7] return 4",
			args: args{
				nums: []int{2, 3, 5, 7},
			},
			want: 4,
		},
		{
			name: "[7,7,7,7,7,7,7] return 1",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "[0,1,0,3] return 3",
			args: args{
				nums: []int{0, 1, 0, 3},
			},
			want: 3,
		},
		{
			name: "[0,1,0,3,2,3] return 4",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "[0,1,0] return 2",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,1,7] return 3",
			args: args{
				nums: []int{2, 3, 1, 7},
			},
			want: 3,
		},
		{
			name: "[2,3,5,7] return 4",
			args: args{
				nums: []int{2, 3, 5, 7},
			},
			want: 4,
		},
		{
			name: "[7,7,7,7,7,7,7] return 1",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "[0,1,0,3] return 3",
			args: args{
				nums: []int{0, 1, 0, 3},
			},
			want: 3,
		},
		{
			name: "[0,1,0,3,2,3] return 4",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "[0,1,0] return 2",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS3(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS1() = %v, want %v", got, tt.want)
			}
		})
	}
}
