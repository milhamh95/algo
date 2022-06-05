package fibonacci

import "testing"

func Test_generateFibonacci1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 2, output = 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "n = 3, output = 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFibonacci1(tt.args.n); got != tt.want {
				t.Errorf("generateFibonacci1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateFibonacci2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 2, output = 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "n = 3, output = 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFibonacci2(tt.args.n); got != tt.want {
				t.Errorf("generateFibonacci2() = %v, want %v", got, tt.want)
			}
		})
	}
}
