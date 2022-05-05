package besttimetobuyandsellstock

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max profit for prices = [7,1,5,3,6,4] is 5",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "max profit for prices = [7,6,4,3,1] is 0",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "max profit for prices = [1,2,4,2,5,7,2,4,9,0,9] is 9",
			args: args{
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 9,
		},
		{
			name: "max profit for prices = [1,3,0,9] is 9",
			args: args{
				prices: []int{1, 3, 0, 9},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max profit for prices = [7,1,5,3,6,4] is 5",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "max profit for prices = [7,6,4,3,1] is 0",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "max profit for prices = [1,2,4,2,5,7,2,4,9,0,9] is 9",
			args: args{
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 9,
		},
		{
			name: "max profit for prices = [1,3,0,9] is 9",
			args: args{
				prices: []int{1, 3, 0, 9},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max profit for prices = [7,1,5,3,6,4] is 5",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "max profit for prices = [7,6,4,3,1] is 0",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "max profit for prices = [1,2,4,2,5,7,2,4,9,0,9] is 9",
			args: args{
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 9,
		},
		{
			name: "max profit for prices = [1,3,0,9] is 9",
			args: args{
				prices: []int{1, 3, 0, 9},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit3(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
