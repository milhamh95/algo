package palindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: " x = 123, output false",
			args: args{
				x: 123,
			},
			want: false,
		},
		{
			name: " x = 121, output true",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: " x = -121, output false",
			args: args{
				x: -121,
			},
			want: false,
		},

		{
			name: " x = 10, output false",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome1(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: " x = 123, output false",
			args: args{
				x: 123,
			},
			want: false,
		},
		{
			name: " x = 121, output true",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: " x = -121, output false",
			args: args{
				x: -121,
			},
			want: false,
		},

		{
			name: " x = 10, output false",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome1(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
