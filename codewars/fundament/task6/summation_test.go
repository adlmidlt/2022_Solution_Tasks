package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Basic Tests", func() {
	examples := [...][2]int{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}
	for i := 0; i < len(examples); i++ {
		v := examples[i]
		It(fmt.Sprintf("Testing for %d", v[0]), func() {
			Expect(Summation(v[0])).To(Equal(v[1]))
		})
	}
})

func TestSummation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success1",
			args: args{
				n: 2,
			},
			want: 3,
		},
		{
			name: "success2",
			args: args{
				n: 213,
			},
			want: 22791,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summation(tt.args.n); got != tt.want {
				t.Errorf("Summation() = %v, want %v", got, tt.want)
			}
		})
	}
}
