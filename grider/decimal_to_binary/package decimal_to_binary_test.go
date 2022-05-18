package decimal_to_binary

import "testing"

func Test_decimalToBinary(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input 7, return 111",
			args: args{
				num: 7,
			},
			want: "111",
		},
		{
			name: "input 13, return 1101",
			args: args{
				num: 13,
			},
			want: "1101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decimalToBinary(tt.args.num); got != tt.want {
				t.Errorf("decimalToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
