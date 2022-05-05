package validanagram

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "anagram and nagaram is an anagram",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "anagram and nagarama isn't an anagram",
			args: args{
				s: "anagram",
				t: "nagarama",
			},
			want: false,
		},
		{
			name: "rat and car isn't an anagram",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "anagram and nagaram is an anagram",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "anagram and nagarama isn't an anagram",
			args: args{
				s: "anagram",
				t: "nagarama",
			},
			want: false,
		},
		{
			name: "rat and car isn't an anagram",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "aacc and ccac isn't an anagram",
			args: args{
				s: "aacc",
				t: "ccac",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram2() = %v, want %v", got, tt.want)
			}
		})
	}
}
