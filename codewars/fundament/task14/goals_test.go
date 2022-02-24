package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Sample tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Goals(0, 0, 0)).To(Equal(0))
		Expect(Goals(43, 10, 5)).To(Equal(58))
	})
})

func TestGoals(t *testing.T) {
	type args struct {
		laLigaGoals          int
		copaDelReyGoals      int
		championsLeagueGoals int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success1",
			args: args{
				laLigaGoals:          0,
				copaDelReyGoals:      0,
				championsLeagueGoals: 0,
			},
			want: 0,
		},
		{
			name: "success1",
			args: args{
				laLigaGoals:          43,
				copaDelReyGoals:      10,
				championsLeagueGoals: 5,
			},
			want: 58,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Goals(tt.args.laLigaGoals, tt.args.copaDelReyGoals, tt.args.championsLeagueGoals); got != tt.want {
				t.Errorf("Goals() = %v, want %v", got, tt.want)
			}
		})
	}
}
