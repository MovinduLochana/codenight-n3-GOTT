package main

import (
	"reflect"
	"testing"
)

func TestCheckout(t *testing.T) {
	got, err := Checkout([]string{"Coffee", "Tea"}, []int{2, 3})
	want := []int{4, 6}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("Checkout = %v, %v; want %v, nil", got, err, want)
	}
}
