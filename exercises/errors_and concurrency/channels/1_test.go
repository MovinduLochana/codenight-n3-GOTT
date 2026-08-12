package main

import "testing"

func TestProducer(t *testing.T) {
	ch := make(chan int, 3)
	Producer(ch, 3)
	count := 0
	for range ch {
		count++
	}
	if count != 3 {
		t.Errorf("Producer count = %d; want 3", count)
	}
}
