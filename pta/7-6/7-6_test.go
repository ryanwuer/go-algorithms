package main

import (
	"reflect"
	"testing"
)

func Test_getGapValue(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "case2",
			args: args{n: 3},
			want: 4,
		},
		{
			name: "case3",
			args: args{n: 4},
			want: 10,
		},
		{
			name: "case4",
			args: args{n: 5},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGapValue(tt.args.n); got != tt.want {
				t.Errorf("getGapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "case1",
			args: args{k: 1},
			want: [][]int{},
		},
		{
			name: "case2",
			args: args{k: 2},
			want: [][]int{},
		},
		{
			name: "case3",
			args: args{k: 26},
			want: [][]int{
				{4, 5, 7, 10},
			},
		},
		{
			name: "case4",
			args: args{k: 55},
			want: [][]int{
				{27, 28},
				{17, 18, 20},
				{7, 8, 10, 13, 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
