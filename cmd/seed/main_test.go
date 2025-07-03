package main

import "testing"

func BenchmarkRunWithoutGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runWithoutGoroutines()
	}
}

func BenchmarkRunWithGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runWithGoroutines()
	}
}
