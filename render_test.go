package gmock

import (
	"reflect"
	"testing"
)

func TestParseQuery(t *testing.T) {
	type args struct {
		queryStr string
	}
	tests := []struct {
		name         string
		args         args
		wantQueryMap map[string]interface{}
	}{
		{
			name: "basic",
			args: args{queryStr: "foo=bar"},
			wantQueryMap: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "empty",
			args: args{queryStr: "foo="},
			wantQueryMap: map[string]interface{}{
				"foo": "",
			},
		},
		{
			name: "array",
			args: args{queryStr: "foo=123&foo=321&foo=abc"},
			wantQueryMap: map[string]interface{}{
				"foo": []string{"123", "321", "abc"},
			},
		},
		{
			name: "invalid",
			args: args{queryStr: "invalid"},
			wantQueryMap: map[string]interface{}{
				"invalid": "",
			},
		},
		{
			name: "multiple parameters",
			args: args{queryStr: "foo=123&bar=abc"},
			wantQueryMap: map[string]interface{}{
				"foo": "123",
				"bar": "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQueryMap := ParseQuery(tt.args.queryStr); !reflect.DeepEqual(gotQueryMap, tt.wantQueryMap) {
				t.Errorf("ParseQuery() = %v, want %v", gotQueryMap, tt.wantQueryMap)
			}
		})
	}
}
