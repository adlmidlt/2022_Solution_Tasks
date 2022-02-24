package main

import "testing"

func TestDNAtoRNA(t *testing.T) {
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
				dna: "GCAT",
			},
			want: "GCAU",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DNAtoRNA(tt.args.dna); got != tt.want {
				t.Errorf("DNAtoRNA() = %v, want %v", got, tt.want)
			}
		})
	}
}
