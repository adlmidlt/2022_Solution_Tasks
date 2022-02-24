package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Sample Tests", func() {
	It("Testing [1,2]", func() { Expect(SquareSum([]int{1, 2})).To(Equal(5)) })
	It("Testing [0,3,4,5]", func() { Expect(SquareSum([]int{0, 3, 4, 5})).To(Equal(50)) })
	It("Testing []", func() { Expect(SquareSum([]int{})).To(Equal(0)) })
})

func TestSquareSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success1",
			args: args{
				numbers: []int{1, 2},
			},
			want: 5,
		},
		{
			name: "success1",
			args: args{
				numbers: []int{0, 3, 4, 5},
			},
			want: 50,
		},
		{
			name: "success1",
			args: args{
				numbers: []int{},
			},
			want: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareSum(tt.args.numbers); got != tt.want {
				t.Errorf("SquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
