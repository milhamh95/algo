package trapping_rain_water

import "testing"

func Test_trap1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "height = [4,2,0,3,2,5], output = 9",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
		{
			name: "height = [1,1], output = 0",
			args: args{
				height: []int{1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap1(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap2(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "height = [4,2,0,3,2,5], output = 9",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
		{
			name: "height = [1,1], output = 0",
			args: args{
				height: []int{1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap2(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap3(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "height = [4,2,0,3,2,5], output = 9",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
		{
			name: "height = [1,1], output = 0",
			args: args{
				height: []int{1, 1},
			},
			want: 0,
		},
		{
			name: "height = [4,5,1,3], output = 1",
			args: args{
				height: []int{4, 5, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap3(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}
