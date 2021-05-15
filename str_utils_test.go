package main

import "testing"

func TestRemoveFuzzy(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestRemoveFuzzy",
			args: args{
				str: "Hello",
			},
			want: "He",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if str := RemoveFuzzy(tt.args.str); str != tt.want {
				t.Errorf("[ERROR] got = %v; want = %v", str, tt.want)
			}
		})
	}
}
