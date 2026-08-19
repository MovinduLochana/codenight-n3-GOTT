package main

import (
	"reflect"
	"testing"
)

func TestAppendThree(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{"empty", []int{}, []int{1, 2, 3}},
		{"single", []int{9}, []int{9, 1, 2, 3}},
		{"existing", []int{1, 2}, []int{1, 2, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendThree(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendThree(%v) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}