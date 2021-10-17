package dimea_test

import (
	"math/rand"
	"testing"

	"github.com/yut-kt/dimea"
)

func makeSlice() []float64 {
	s := make([]float64, 100000)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}

func BenchmarkEuclidean(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.Euclidean(x, y)
	}
}

func BenchmarkSquaredEuclidean(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.SquaredEuclidean(x, y)
	}
}

func BenchmarkCosineSimilarity(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.CosineSimilarity(x, y)
	}
}

func BenchmarkHamming(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.Hamming(x, y)
	}
}

func BenchmarkManhattan(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.Manhattan(x, y)
	}
}

func BenchmarkChebyshev(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.Chebyshev(x, y)
	}
}

func BenchmarkMinkowski(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.Minkowski(x, y, 1)
	}
}

func BenchmarkJaccardIndex(b *testing.B) {
	x, y := makeSlice(), makeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dimea.JaccardIndex(x, y)
	}
}
