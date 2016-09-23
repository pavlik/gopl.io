package main

import "testing"

func BenchmarkAlgo(b *testing.B) {
	b.Run("Slow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slowAlgorithm()
		}
	})

	b.Run("Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fastAlgorithm()
		}
	})
}
