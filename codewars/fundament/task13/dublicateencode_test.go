package main

import (
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should return the correct value", func() {
		Expect(DuplicateEncode("din")).To(Equal("((("))
		Expect(DuplicateEncode("recede")).To(Equal("()()()"))
		Expect(DuplicateEncode("(( @")).To(Equal("))(("))
	})

	It("should ignore case", func() {
		Expect(DuplicateEncode("Success")).To(Equal(")())())"))
	})
})

func TestDuplicateEncode(t *testing.T) {
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
				word: "din",
			},
			want: "(((",
		},
		{
			name: "success2",
			args: args{
				word: "recede",
			},
			want: "()()()",
		},
		{
			name: "success3",
			args: args{
				word: "(( @",
			},
			want: "))((",
		},
		{
			name: "success4",
			args: args{
				word: "Success",
			},
			want: ")())())",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DuplicateEncode(tt.args.word); got != tt.want {
				t.Errorf("DuplicateEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
