package reverseint

import "testing"

func TestReverseInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "15 reverse to 51",
			args: args{
				n: 15,
			},
			want: 51,
		},
		{
			name: "123 reverse to 321",
			args: args{
				n: 123,
			},
			want: 321,
		},
		{
			name: "981 reverse to 189",
			args: args{
				n: 981,
			},
			want: 189,
		},
		{
			name: "500 reverse to 5",
			args: args{
				n: 500,
			},
			want: 5,
		},
		{
			name: "-15 reverse to -51",
			args: args{
				n: -15,
			},
			want: -51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt(tt.args.n); got != tt.want {
				t.Errorf("ReverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
