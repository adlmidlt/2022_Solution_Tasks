package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("RemoveChar", func() {
	Describe("Fixed Tests", func() {
		It("eloquent", func() {
			Expect(RemoveChar("eloquent")).To(Equal("loquen"))
		})
		It("country", func() {
			Expect(RemoveChar("country")).To(Equal("ountr"))
		})
		It("person", func() {
			Expect(RemoveChar("person")).To(Equal("erso"))
		})
		It("place", func() {
			Expect(RemoveChar("place")).To(Equal("lac"))
		})
	})
})

func TestRemoveChar(t *testing.T) {
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
				word: "eloquent",
			},
			want: "loquen",
		},
		{
			name: "success2",
			args: args{
				word: "country",
			},
			want: "ountr",
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveChar(tt.args.word); got != tt.want {
				t.Errorf("RemoveChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
