package main

import (
	"bytes"
	"io"
	"sync"
	"testing"
)

func Test_amp(t *testing.T) {
	type args struct {
		setting string
		in      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := amp(tt.args.setting, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("amp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("amp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_feedbackAmp(t *testing.T) {
	type args struct {
		read io.Reader
		wg   *sync.WaitGroup
		tag  string
	}
	tests := []struct {
		name      string
		args      args
		wantWrite string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			write := &bytes.Buffer{}
			feedbackAmp(tt.args.read, write, tt.args.wg, tt.args.tag)
			if gotWrite := write.String(); gotWrite != tt.wantWrite {
				t.Errorf("feedbackAmp() = %v, want %v", gotWrite, tt.wantWrite)
			}
		})
	}
}

func Test_logger_Write(t *testing.T) {
	type fields struct {
		path string
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				path: tt.fields.path,
			}
			got, err := l.Write(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}
