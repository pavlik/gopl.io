package main

import "testing"

func BenchmarkSlowAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowAlgorithm()
	}
}

func BenchmarkFastAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastAlgorithm()
	}
}
