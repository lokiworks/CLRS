package main

import (
	"reflect"
	"testing"
)

func Test_bucketSort(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{
			name: "WithCase1",
			args: args{[]float64{0.897, 0.565, 0.656, 0.1234}},
			want: []float64{0.1234, 0.565, 0.656, 0.897},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bucketSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
