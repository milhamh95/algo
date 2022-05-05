package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "apple reverse to elppa",
			args: args{
				s: "apple",
			},
			want: "elppa",
		},
		{
			name: "hello reverse to olleh",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
