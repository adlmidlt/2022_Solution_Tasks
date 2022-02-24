package main

import "testing"

func TestRepeatStr(t *testing.T) {
	type args struct {
		repetitions int
		value       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success1",
			args: args{
				repetitions: 4,
				value:       "a",
			},
			want: "aaaa",
		},
		{
			name: "success2",
			args: args{
				repetitions: 3,
				value:       "hello ",
			},
			want: "hello hello hello ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatStr(tt.args.repetitions, tt.args.value); got != tt.want {
				t.Errorf("RepeatStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
