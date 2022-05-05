package maxchar

import "testing"

func TestMaxChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abcccccd frequent char is c",
			args: args{
				s: "abcccccd",
			},
			want: "c",
		},
		{
			name: "ab1134d frequent char is 1",
			args: args{
				s: "ab1134d",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxChar(tt.args.s); got != tt.want {
				t.Errorf("MaxChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
