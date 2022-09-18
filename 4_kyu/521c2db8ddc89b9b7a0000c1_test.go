package kata

import (
	"reflect"
	"testing"
)

func TestSnail(t *testing.T) {
	type args struct {
		snaipMap [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"sample",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"empty",
			args{
				[][]int{},
			},
			[]int{},
		},
		{
			"empty subarray",
			args{
				[][]int{{}},
			},
			[]int{},
		},
		{
			"one element",
			args{
				[][]int{{1}},
			},
			[]int{1},
		},
		{
			"two array",
			args{
				[][]int{{1, 2}, {4, 3}},
			},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snail(tt.args.snaipMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Snail() = %v, want %v", got, tt.want)
			}
		})
	}
}
