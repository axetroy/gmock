package function

import "testing"

func TestEscape(t *testing.T) {
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
			args: args{str: "\"hello world\""},
			want: "\\\"hello world\\\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Escape(tt.args.str); got != tt.want {
				t.Errorf("Escape() = %v, want %v", got, tt.want)
			}
		})
	}
}
