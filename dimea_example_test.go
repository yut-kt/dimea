package dimea_test

import (
	"fmt"
	"math"

	"github.com/yut-kt/dimea"
)

func makeSliceXY() ([]float64, []float64) {
	return []float64{1, 2, 3, 4, 5}, []float64{5, 4, 3, 2, 1}
}

func ExampleEuclidean() {
	x, y := makeSliceXY()
	v, _ := dimea.Euclidean(x, y)
	fmt.Println(v)
	// Output:
	// 6.324555320336759
}

func ExampleSquaredEuclidean() {
	x, y := makeSliceXY()
	v, _ := dimea.SquaredEuclidean(x, y)
	fmt.Println(v)
	// Output:
	// 40
}

func ExampleCosineSimilarity() {
	x, y := makeSliceXY()
	v, _ := dimea.CosineSimilarity(x, y)
	fmt.Println(v)
	// Output:
	// 0.6363636363636364
}

func ExampleHamming() {
	x, y := "karolin", "kathrin"
	v, _ := dimea.Hamming(x, y)
	fmt.Println(v)
	// Output:
	// 3
}

func ExampleManhattan() {
	x, y := makeSliceXY()
	v, _ := dimea.Manhattan(x, y)
	fmt.Println(v)
	// Output:
	// 12
}

func ExampleChebyshev() {
	x, y := makeSliceXY()
	v, _ := dimea.Chebyshev(x, y)
	fmt.Println(v)
	// Output:
	// 4
}

func ExampleMinkowski() {
	x, y := makeSliceXY()
	pSlice := []float64{1, 2, math.Inf(1)}
	for _, p := range pSlice {
		v, _ := dimea.Minkowski(x, y, p)
		fmt.Println(v)
	}
	fmt.Println()
	// Output:
	// 12
	// 6.324555320336759
	// 4
}

func ExampleJaccardIndex() {
	x, y := makeSliceXY()
	v, _ := dimea.JaccardIndex(x, y)
	fmt.Println(v)
	// Output:
	// 1
}

func ExampleLevenshtein_Distance() {
	fmt.Println(new(dimea.Levenshtein).SetCosts(7, 7, 10).Distance("agent", "agency"))
	// Output:
	// 17
}

func ExampleLevenshtein_StdDistance() {
	fmt.Printf(
		"%f",
		new(dimea.Levenshtein).
			SetInsertionsCost(7).
			SetDeletionCost(7).
			SetSubstitutionCost(10).
			StdDistance("agent", "agency"),
	)
	// Output:
	// 2.833333
}
