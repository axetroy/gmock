package gmock_test

import (
	"net/http"
	"net/url"
	"os"
	path "path/filepath"
	"reflect"
	"regexp"
	"testing"

	"github.com/axetroy/gmock"
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
			want: regexp.MustCompile(`^[\w\d-]+$`),
		},
		{
			name: "params & special characters 1",
			args: args{
				name: "[id].name",
			},
			want: regexp.MustCompile(`^[\w\d-]+\.name$`),
		},
		{
			name: "params & another special characters 2",
			args: args{
				name: "[id]-profile",
			},
			want: regexp.MustCompile(`^[\w\d-]+\-profile$`),
		},
		{
			name: "params & another special characters 3",
			args: args{
				name: "^[id]-profile",
			},
			want: regexp.MustCompile(`^\^[\w\d-]+\-profile$`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gmock.NameToRegExp(tt.args.name); !reflect.DeepEqual(got, tt.want) {
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
			if got := gmock.MatchFileName(tt.args.rule, tt.args.name); got != tt.want {
				t.Errorf("MatchFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookup(t *testing.T) {
	cwd, _ := os.Getwd()

	helloWorld := path.Join(cwd, "__test__", "hello_world", "hello.get.json")
	paramsUserID := path.Join(cwd, "__test__", "params", "user", "[id].get.json")
	paramsCardID := path.Join(cwd, "__test__", "params", "[zoo_id]", "detail.get.json")

	type args struct {
		rootDir string
		method  string
		u       *url.URL
	}
	tests := []struct {
		name  string
		args  args
		want  *string
		want1 map[string]string
	}{
		{
			name: "basic",
			args: args{
				rootDir: path.Join(cwd, "__test__", "hello_world"),
				method:  http.MethodGet,
				u:       &url.URL{Path: "/hello"},
			},
			want:  &helloWorld,
			want1: map[string]string{},
		},
		{
			name: "params 1",
			args: args{
				rootDir: path.Join(cwd, "__test__", "params"),
				method:  http.MethodGet,
				u:       &url.URL{Path: "/user/root"},
			},
			want: &paramsUserID,
			want1: map[string]string{
				"id": "root",
			},
		},
		{
			name: "params 2",
			args: args{
				rootDir: path.Join(cwd, "__test__", "params"),
				method:  http.MethodGet,
				u:       &url.URL{Path: "/my_zoo_id/detail"},
			},
			want: &paramsCardID,
			want1: map[string]string{
				"zoo_id": "my_zoo_id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := gmock.Lookup(tt.args.rootDir, tt.args.method, tt.args.u)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lookup() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Lookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetRealFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{fileName: "profile.get.json"},
			want: "profile",
		},
		{
			name: "basic params",
			args: args{fileName: "[id].get.json"},
			want: "[id]",
		},
		{
			name: "basic & params",
			args: args{fileName: "user_[id].get.json"},
			want: "user_[id]",
		},
		{
			name: "basic & multiple params",
			args: args{fileName: "user_[id]_bank_[bank_name].get.json"},
			want: "user_[id]_bank_[bank_name]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gmock.GetRealFileName(tt.args.fileName); got != tt.want {
				t.Errorf("GetRealFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractParamsFromFileName(t *testing.T) {
	type args struct {
		fileName string
		urlPath  string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "basic",
			args: args{
				fileName: "[id].get.json",
				urlPath:  "123",
			},
			want: map[string]string{
				"id": "123",
			},
		},
		{
			name: "basic",
			args: args{
				fileName: "[user_id].get.json",
				urlPath:  "123",
			},
			want: map[string]string{
				"user_id": "123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gmock.ExtractParamsFromFileName(tt.args.fileName, tt.args.urlPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractParamsFromFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
