package app

import (
	"net/http"
	"net/url"
	"os"
	"path"
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
			name: "params",
			args: args{
				name: "[id]",
			},
			want: regexp.MustCompile("^[\\w\\d-]+$"),
		},
		{
			name: "params & special characters 1",
			args: args{
				name: "[id].name",
			},
			want: regexp.MustCompile("^[\\w\\d-]+\\.name$"),
		},
		{
			name: "params & another special characters 2",
			args: args{
				name: "[id]-profile",
			},
			want: regexp.MustCompile("^[\\w\\d-]+\\-profile$"),
		},
		{
			name: "params & another special characters 3",
			args: args{
				name: "^[id]-profile",
			},
			want: regexp.MustCompile("^\\^[\\w\\d-]+\\-profile$"),
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
				rule: "[id]",
				name: "123",
			},
			want: true,
		},
		{
			name: "basic & common characters",
			args: args{
				rule: "user_[id]",
				name: "user_123",
			},
			want: true,
		},
		{
			name: "params & special characters 1",
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

func TestLookup(t *testing.T) {
	cwd, _ := os.Getwd()

	helloWorld := path.Join(cwd, "__test__", "hello_world", "hello.get.json")
	paramsUserID := path.Join(cwd, "__test__", "params", "user", "[id].get.json")
	paramsCardID := path.Join(cwd, "__test__", "params", "[card_id]", "detail.get.json")

	type args struct {
		rootDir string
		method  string
		u       *url.URL
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "basic",
			args: args{
				rootDir: path.Join(cwd, "__test__", "hello_world"),
				method:  http.MethodGet,
				u:       &url.URL{Path: "/hello"},
			},
			want: &helloWorld,
		},
		{
			name: "params 1",
			args: args{
				rootDir: path.Join(cwd, "__test__", "params"),
				method:  http.MethodGet,
				u:       &url.URL{Path: "/user/user_id"},
			},
			want: &paramsUserID,
		},
		{
			name: "params 2",
			args: args{
				rootDir: path.Join(cwd, "__test__", "params"),
				method:  http.MethodGet,
				u:       &url.URL{Path: "/card_id/detail"},
			},
			want: &paramsCardID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lookup(tt.args.rootDir, tt.args.method, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
