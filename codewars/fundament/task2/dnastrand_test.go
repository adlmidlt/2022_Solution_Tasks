package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("Basic Tests", func() {
		Expect(DNAStrand("AAAA")).To(Equal("TTTT"))
		Expect(DNAStrand("ATTGC")).To(Equal("TAACG"))
		Expect(DNAStrand("GTAT")).To(Equal("CATA"))
	})
})

func TestDNAStrand(t *testing.T) {
	type args struct {
		dna string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				dna: "ATTGC",
			},
			want: "TAACG",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DNAStrand(tt.args.dna); got != tt.want {
				t.Errorf("DNAStrand() = %v, want %v", got, tt.want)
			}
		})
	}
}
