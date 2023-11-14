package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

type data struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestLogJSON(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "encode string",
			args: args{
				data: string("santekno"),
			},
			want: `"santekno"`,
		},
		{
			name: "encode number",
			args: args{
				data: 2,
			},
			want: "2",
		},
		{
			name: "encode boolean",
			args: args{
				data: true,
			},
			want: "true",
		},
		{
			name: "encode array string",
			args: args{
				data: []string{"santekno", "ihsan"},
			},
			want: `["santekno","ihsan"]`,
		},
		{
			name: "encode array object",
			args: args{
				data: data{
					FirstName:  "Ihsan",
					MiddleName: "Arif",
					LastName:   "Rahman",
				},
			},
			want: `{"FirstName":"Ihsan","MiddleName":"Arif","LastName":"Rahman"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, LogJSON(tt.args.data), tt.want)
		})
	}
}
