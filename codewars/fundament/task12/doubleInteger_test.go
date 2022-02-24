package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should work for some examples", func() {
		Expect(DoubleInteger(1)).To(Equal(2))
	})
})

func TestDoubleInteger(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success1",
			args: args{
				i: 18,
			},
			want: 36,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleInteger(tt.args.i); got != tt.want {
				t.Errorf("DoubleInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
