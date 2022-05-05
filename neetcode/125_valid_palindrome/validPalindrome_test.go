package validpalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "race a car isn't a palindrome",
			args: args{
				s: "race a car",
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
			name: "Never odd or even is a palindrome",
			args: args{
				s: "Never odd or even",
			},
			want: true,
		},
		{
			name: "A man, a plan, a canal: Panama is a palindrome",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
