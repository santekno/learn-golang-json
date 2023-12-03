package main

import (
	"reflect"
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

func TestGenerateObjectJSON(t *testing.T) {
	type args struct {
		data Customer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success generate object JSON",
			args: args{
				data: Customer{
					FirstName:  "Santekno",
					MiddleName: "Ihsan",
					LastName:   "Arif",
					Hobbies:    []string{"badminton", "renang", "coding"},
				},
			},
			want: string(`{"first_name":"Santekno","middle_name":"Ihsan","last_name":"Arif","hobbies":["badminton","renang","coding"]}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateObjectJSON(tt.args.data); got != tt.want {
				t.Errorf("GenerateObjectJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertObjectJSON(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want Customer
	}{
		{
			name: "success conversion object JSON",
			args: args{
				data: string(`{"first_name":"Santekno","middle_name":"Ihsan","last_name":"Arif","hobbies":["badminton","renang","coding"]}`),
			},
			want: Customer{
				FirstName:  "Santekno",
				MiddleName: "Ihsan",
				LastName:   "Arif",
				Hobbies:    []string{"badminton", "renang", "coding"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertObjectJSON(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertObjectJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
