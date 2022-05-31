package search_a_2d_matrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3, output = true",
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13, output = false",
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix1(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3, output = true",
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13, output = false",
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix1(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix2(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3, output = true",
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13, output = false",
			args: args{
				matrix: [][]int{
					[]int{1, 3, 5, 7},
					[]int{10, 11, 16, 20},
					[]int{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix2(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
