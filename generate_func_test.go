package main

import "testing"

func TestGenPass(t *testing.T) {
	type args struct {
		base   string
		length int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestGenPass-001",
			args: args{
				base:   "0123456789abcdef",
				length: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("result = %v", RandStr(tt.args.base, tt.args.length))
		})
	}
}

func TestRandDigital(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestRandDigital-001",
			args: args{
				length: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("result = %v", RandDigital(tt.args.length))
		})
	}
}

func TestRandHex(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestRandHex-001",
			args: args{
				length: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("result = %v", RandHex(tt.args.length))
		})
	}
}

func TestRandDigitalAndAlphabet(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestRandDigitalAndAlphabet-001",
			args: args{
				length: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("result = %v", RandDigitalAndAlphabet(tt.args.length))
		})
	}
}
