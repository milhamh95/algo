package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tenet is a palindrome",
			args: args{
				s: "tenet",
			},
			want: true,
		},
		{
			name: "noon is a palindrome",
			args: args{
				s: "noon",
			},
			want: true,
		},
		{
			name: "sator isn't a palindrome",
			args: args{
				s: "sator",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.args.s); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
