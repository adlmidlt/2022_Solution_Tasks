package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("basic tests", func() {
		Expect(ToAlternatingCase("hello world")).To(Equal("HELLO WORLD"))
		Expect(ToAlternatingCase("HELLO WORLD")).To(Equal("hello world"))
		Expect(ToAlternatingCase("hello WORLD")).To(Equal("HELLO world"))
		Expect(ToAlternatingCase("HeLLo WoRLD")).To(Equal("hEllO wOrld"))
		Expect(ToAlternatingCase("12345")).To(Equal("12345"))
		Expect(ToAlternatingCase("1a2b3c4d5e")).To(Equal("1A2B3C4D5E"))
		Expect(ToAlternatingCase("String.prototype.toAlternatingCase")).To(Equal("sTRING.PROTOTYPE.TOaLTERNATINGcASE"))
		Expect(ToAlternatingCase(ToAlternatingCase("Hello World"))).To(Equal("Hello World"))
		Expect(ToAlternatingCase("altERnaTIng cAsE")).To(Equal("ALTerNAtiNG CaSe"))
		Expect(ToAlternatingCase("ALTerNAtiNG CaSe")).To(Equal("altERnaTIng cAsE"))
		Expect(ToAlternatingCase("altERnaTIng cAsE <=> ALTerNAtiNG CaSe")).To(Equal("ALTerNAtiNG CaSe <=> altERnaTIng cAsE"))
		Expect(ToAlternatingCase("ALTerNAtiNG CaSe <=> altERnaTIng cAsE")).To(Equal("altERnaTIng cAsE <=> ALTerNAtiNG CaSe"))
	})
})

func TestToAlternatingCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success1",
			args: args{
				str: "hello world",
			},
			want: "HELLO WORLD",
		},
		{
			name: "success2",
			args: args{
				str: "12345",
			},
			want: "12345",
		},
		{
			name: "success3",
			args: args{
				str: "String.prototype.toAlternatingCase",
			},
			want: "sTRING.PROTOTYPE.TOaLTERNATINGcASE",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAlternatingCase(tt.args.str); got != tt.want {
				t.Errorf("ToAlternatingCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
