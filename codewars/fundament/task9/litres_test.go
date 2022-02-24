package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Test Examples", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(Litres(2)).To(Equal(1))
		Expect(Litres(1.4)).To(Equal(0))
		Expect(Litres(12.3)).To(Equal(6))
		Expect(Litres(0.82)).To(Equal(0))
		Expect(Litres(11.8)).To(Equal(5))
		Expect(Litres(1787)).To(Equal(893))
		Expect(Litres(0)).To(Equal(0))
	})
})

func TestLitres(t *testing.T) {
	type args struct {
		time float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success1",
			args: args{
				time: 2,
			},
			want: 1,
		},
		{
			name: "success2",
			args: args{
				time: 1.4,
			},
			want: 0,
		},
		{
			name: "success3",
			args: args{
				time: 12.3,
			},
			want: 6,
		},
		{
			name: "success4",
			args: args{
				time: 0.82,
			},
			want: 0,
		},
		{
			name: "success5",
			args: args{
				time: 11.8,
			},
			want: 5,
		},
		{
			name: "success6",
			args: args{
				time: 1787,
			},
			want: 893,
		},
		{
			name: "success7",
			args: args{
				time: 0,
			},
			want: 0,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Litres(tt.args.time); got != tt.want {
				t.Errorf("Litres() = %v, want %v", got, tt.want)
			}
		})
	}
}
