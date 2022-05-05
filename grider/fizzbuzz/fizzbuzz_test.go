package fizzbuzz

import "testing"

func TestFizzBuzz2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "3 return fizz",
			args: args{
				n: 3,
			},
			want: "fizz",
		},
		{
			name: "5 return buzz",
			args: args{
				n: 5,
			},
			want: "buzz",
		},
		{
			name: "15 return fizzbuzz",
			args: args{
				n: 15,
			},
			want: "fizzbuzz",
		},
		{
			name: "1 return 1",
			args: args{
				n: 1,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz2(tt.args.n); got != tt.want {
				t.Errorf("FizzBuzz2() = %v, want %v", got, tt.want)
			}
		})
	}
}
