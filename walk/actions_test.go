package main

import (
	"bytes"
	"io/fs"
	"testing"
)

func Test_filterOut(t *testing.T) {
	type args struct {
		path string
		ext  string
		size int64
		info fs.FileInfo
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterOut(tt.args.path, tt.args.ext, tt.args.size, tt.args.info); got != tt.want {
				t.Errorf("filterOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := listFile(tt.args.path, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("listFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("listFile() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
