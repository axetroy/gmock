package function

import (
	"testing"
)

func TestBase64Encoding(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{str: "hello world"},
			want: "aGVsbG8gd29ybGQ=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encoding(tt.args.str); got != tt.want {
				t.Errorf("Base64Encoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64Decoding(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{str: "aGVsbG8gd29ybGQ="},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Decoding(tt.args.str); got != tt.want {
				t.Errorf("Base64Decoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
