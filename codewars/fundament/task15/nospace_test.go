package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Tests", func() {
	It("should work for sample tests", func() {
		Expect(NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B")).To(Equal("8j8mBliB8gimjB8B8jlB"))
		Expect(NoSpace("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")).To(Equal("88Bifk8hB8BB8BBBB888chl8BhBfd"))
		Expect(NoSpace("8aaaaa dddd r     ")).To(Equal("8aaaaaddddr"))
		Expect(NoSpace("jfBm  gk lf8hg  88lbe8 ")).To(Equal("jfBmgklf8hg88lbe8"))
		Expect(NoSpace("8j aam")).To(Equal("8jaam"))
	})
})

func TestNoSpace(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success1",
			args: args{
				word: "8 j 8   mBliB8g  imjB8B8  jl  B",
			},
			want: "8j8mBliB8gimjB8B8jlB",
		},
		{
			name: "success2",
			args: args{
				word: "8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd",
			},
			want: "88Bifk8hB8BB8BBBB888chl8BhBfd",
		},
		{
			name: "success3",
			args: args{
				word: "8aaaaa dddd r     ",
			},
			want: "8aaaaaddddr",
		},
		{
			name: "success4",
			args: args{
				word: "jfBm  gk lf8hg  88lbe8 ",
			},
			want: "jfBmgklf8hg88lbe8",
		},
		{
			name: "success5",
			args: args{
				word: "8j aam",
			},
			want: "8jaam",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoSpace(tt.args.word); got != tt.want {
				t.Errorf("NoSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
