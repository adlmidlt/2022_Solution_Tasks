package main

import "testing"

func TestPositiveSum(t *testing.T) {
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
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "succes2",
			args: args{
				numbers: []int{-1, 2, 3, 4, -5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PositiveSum(tt.args.numbers); got != tt.want {
				t.Errorf("PositiveSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
