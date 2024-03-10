package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	lines, _ := ReadLines("input.txt")
	type args struct {
		linesInput        []string
		delimiter         string
		isNeededDelimiter bool
		neededFields      []int
		typeOfFlag        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "first", args: args{linesInput: lines, delimiter: ",", isNeededDelimiter: true, neededFields: []int{0, 2}, typeOfFlag: 2}, want: []string{"hello,my", "zdarova moi,krtuyoi", "nihao zhi,est"}},
		{name: "second", args: args{linesInput: lines, delimiter: ",", isNeededDelimiter: true, neededFields: []int{0}, typeOfFlag: 1}, want: []string{"hello", "zdarova moi", "nihao zhi"}},
		{name: "third", args: args{linesInput: lines, delimiter: ",", isNeededDelimiter: true, neededFields: []int{0, 1, 2}, typeOfFlag: 3}, want: []string{"hello,my,dear", "zdarova moi,krtuyoi,mir", "nihao zhi,est"}},
		{name: "forth", args: args{linesInput: lines, delimiter: " ", isNeededDelimiter: false, neededFields: []int{0, 1, 2}, typeOfFlag: 3}, want: []string{"hello,my,dear,world", "wassup my sweet", "zdarova moi,krtuyoi,mir", "nihao zhi,est"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cut(tt.args.linesInput, tt.args.delimiter, tt.args.isNeededDelimiter, tt.args.neededFields, tt.args.typeOfFlag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFFlagValue(t *testing.T) {
	type args struct {
		fFlagValue string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseFFlagValue(tt.args.fFlagValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFFlagValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFFlagValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseFFlagValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

//func TestReadLines(t *testing.T) {
//	type args struct {
//		fileName string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := ReadLines(tt.args.fileName)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReadLines() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ReadLines() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
