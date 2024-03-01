package main

import (
	"testing"
)

func TestUnwrapString(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "First", args: args{inputStr: "a4bc2d5e"}, want: "aaaabccddddde", wantErr: false,
		},
		{
			name: "Second", args: args{inputStr: "abcd"}, want: "abcd", wantErr: false,
		},
		{
			name: "Third", args: args{inputStr: "45"}, want: "", wantErr: true,
		},
		{
			name: "Forth", args: args{inputStr: ""}, want: "", wantErr: false,
		},
		{
			name: "Fifth", args: args{inputStr: "qwe\\4\\5"}, want: "qwe45", wantErr: false,
		},
		{
			name: "Sixth", args: args{inputStr: "qwe\\4\\5"}, want: "qwe45", wantErr: false,
		},
		{
			name: "Seventh", args: args{inputStr: "qwe\\45"}, want: "qwe44444", wantErr: false,
		},
		{
			name: "Eighth", args: args{inputStr: "qwe\\\\5"}, want: "qwe\\\\\\\\\\", wantErr: false,
		},
		{
			name: "Ninth", args: args{inputStr: "\\\\\\qer"}, want: "", wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnwrapString(tt.args.inputStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnwrapString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnwrapString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
