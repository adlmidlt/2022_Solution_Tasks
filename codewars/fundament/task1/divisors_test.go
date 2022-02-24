package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Should pass some basic tests", func() {
	It("Divisors(1)", func() { Expect(Divisors(1)).To(Equal(1)) })
	It("Divisors(10)", func() { Expect(Divisors(10)).To(Equal(4)) })
	It("Divisors(11)", func() { Expect(Divisors(11)).To(Equal(2)) })
	It("Divisors(54)", func() { Expect(Divisors(54)).To(Equal(8)) })
	It("Divisors(64)", func() { Expect(Divisors(64)).To(Equal(7)) })
})

func TestDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				n: 500000,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divisors(tt.args.n); got != tt.want {
				t.Errorf("Divisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
