package app

import (
	"reflect"
	"regexp"
	"testing"
)

func TestNameToRegExp(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *regexp.Regexp
	}{
		{
			name: "basic",
			args: args{
				name: "profile",
			},
			want: regexp.MustCompile("^profile$"),
		},
		{
			name: "basic",
			args: args{
				name: "[id]",
			},
			want: regexp.MustCompile("^[\\w\\d-]+$"),
		},
		{
			name: "basic",
			args: args{
				name: "[id].name",
			},
			want: regexp.MustCompile("^[\\w\\d-]+\\.name$"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameToRegExp(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NameToRegExp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchFileName(t *testing.T) {
	type args struct {
		rule string
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{
				rule: "user_[id]",
				name: "user_123",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				rule: "[id].profile",
				name: "123.profile",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchFileName(tt.args.rule, tt.args.name); got != tt.want {
				t.Errorf("MatchFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
