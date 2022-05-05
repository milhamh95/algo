package palindrome2

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
			name: "opera isn't a palindrome",
			args: args{
				s: "opera",
			},
			want: false,
		},
		{
			name: "noon is a palindrome",
			args: args{
				s: "noon",
			},
			want: true,
		},
		{
			name: "taco cat is a palindrome",
			args: args{
				s: "taco cat",
			},
			want: true,
		},
		{
			name: "Tenet is a palindrome",
			args: args{
				s: "Tenet",
			},
			want: true,
		},
		{
			name: "Never odd or even is a palindrome",
			args: args{
				s: "Never odd or even",
			},
			want: true,
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
