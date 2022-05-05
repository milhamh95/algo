package group_anagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams1(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "strs = [\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"] return [[\"bat\"],[\"nat\",\"tan\"],[\"ate\",\"eat\",\"tea\"]]",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams1(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagrams2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "strs = [\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"] return [[\"bat\"],[\"nat\",\"tan\"],[\"ate\",\"eat\",\"tea\"]]",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams2(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams1() = %v, want %v", got, tt.want)
			}
		})
	}
}
