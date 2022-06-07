package print_star_based_on_num

import (
	"fmt"
	"testing"
)

func Test_print1(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "n = 5, output = 1 * 3 4 *",
			args: args{
				num: 5,
			},
			want: "1 * 3 4 *",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(tt.args.num); got != tt.want {
				fmt.Println(got)
				fmt.Println(tt.want == got)
				t.Errorf("print() = %v, want %v", got, tt.want)
			}
		})
	}
}
