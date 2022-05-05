package intersectionofthreesortedarrays

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8] return [1,5]",
			args: args{
				nums1: []int{1, 2, 3, 4, 5},
				nums2: []int{1, 2, 5, 7, 9},
				nums3: []int{1, 3, 4, 5, 8},
			},
			want: []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2, tt.args.nums3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
