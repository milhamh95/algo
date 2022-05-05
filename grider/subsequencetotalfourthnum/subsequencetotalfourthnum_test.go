package subsequencetotalfourthnum

import (
	"reflect"
	"testing"
)

func Test_subsequencenum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n = 5, total = 10, return [1,2,3,4]",
			args: args{
				n: 5,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "n = 6, total = 19, return [2,3,4,10]",
			args: args{
				n: 6,
			},
			want: []int{2, 3, 4, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsequencenum(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsequencenum() = %v, want %v", got, tt.want)
			}
		})
	}
}
