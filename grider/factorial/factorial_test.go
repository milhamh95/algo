package factorial

import "testing"

func Test_genFactorial1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3! is 6",
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genFactorial1(tt.args.n); got != tt.want {
				t.Errorf("genFactorial1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genFactorial2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3! is 6",
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genFactorial2(tt.args.n); got != tt.want {
				t.Errorf("genFactorial2() = %v, want %v", got, tt.want)
			}
		})
	}
}
