package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToJSONString(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "nil", args: args{i: nil}, want: `null`},
		{name: "empty list", args: args{i: []string{}}, want: `[]`},
		{name: "empty map", args: args{i: map[string]string{}}, want: `{}`},
		{name: "simple key value", args: args{i: map[string]string{"key": "value"}}, want: `{"key":"value"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.JSONEq(t, tt.want, ToJSONString(tt.args.i))
		})
	}
}
