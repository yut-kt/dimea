package dimea

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
)

func checkDimension(x, y interface{}) error {
	xValue, yValue := reflect.ValueOf(x), reflect.ValueOf(y)
	if xValue.Kind() != yValue.Kind() {
		return errors.New("type is mismatching x == y")
	}

	dimMatchError := func(xLen, yLen int) error {
		return fmt.Errorf("num of dim is not same, [x=%ddim,y=%ddim]", xLen, yLen)
	}
	switch xValue.Kind() {
	case reflect.Slice:
		xDim, yDim := xValue.Len(), yValue.Len()
		if xDim != yDim {
			return dimMatchError(xDim, yDim)
		}
	case reflect.String:
		xRune, yRune := []rune(xValue.String()), []rune(yValue.String())
		if len(xRune) != len(yRune) {
			return dimMatchError(len(xRune), len(yRune))
		}
	default:
		return errors.New("only slice and string are supported")
	}
	return nil
}

// Euclidean distance is https://en.wikipedia.org/wiki/Euclidean_distance
func Euclidean(x, y []float64) (float64, error) {
	if err := checkDimension(x, y); err != nil {
		return 0, err
	}

	var sum float64
	for i := range x {
		sum += (x[i] - y[i]) * (x[i] - y[i])
	}
	return math.Sqrt(sum), nil
}

// SquaredEuclidean distance is https://en.wikipedia.org/wiki/Euclidean_distance#Squared_Euclidean_distance
func SquaredEuclidean(x, y []float64) (float64, error) {
	if err := checkDimension(x, y); err != nil {
		return 0, err
	}

	var sum float64
	for i := range x {
		sum += (x[i] - y[i]) * (x[i] - y[i])
	}
	return sum, nil
}

// CosineSimilarity is https://en.wikipedia.org/wiki/Cosine_similarity
func CosineSimilarity(x, y []float64) (float64, error) {
	if err := checkDimension(x, y); err != nil {
		return 0, err
	}

	var xy, xSum, ySum float64
	for i := range x {
		xy += x[i] * y[i]
		xSum += x[i] * x[i]
		ySum += y[i] * y[i]
	}
	return xy / (math.Sqrt(xSum) * math.Sqrt(ySum)), nil
}

// Manhattan is https://en.wikipedia.org/wiki/Taxicab_geometry
func Manhattan(x, y []float64) (float64, error) {
	if err := checkDimension(x, y); err != nil {
		return 0, err
	}

	var sum float64
	for i := range x {
		sum += math.Abs(x[i] - y[i])
	}
	return sum, nil
}

// Chebyshev distance is https://en.wikipedia.org/wiki/Chebyshev_distance
func Chebyshev(x, y []float64) (float64, error) {
	if err := checkDimension(x, y); err != nil {
		return 0, err
	}

	var max float64
	for i := range x {
		if tmp := math.Abs(x[i] - y[i]); tmp > max {
			max = tmp
		}
	}
	return max, nil
}

// Minkowski is https://en.wikipedia.org/wiki/Minkowski_distance
// p is float, but truncates other than +inf
func Minkowski(x, y []float64, p float64) (float64, error) {
	if err := checkDimension(x, y); err != nil {
		return 0, err
	}

	if math.IsInf(p, 1) {
		return Chebyshev(x, y)
	}

	// p is int
	p = math.Floor(p)
	var sum float64
	for i := range x {
		sum += math.Pow(math.Abs(x[i]-y[i]), p)
	}
	return math.Pow(sum, 1/p), nil
}

// JaccardIndex is https://en.wikipedia.org/wiki/Jaccard_index
func JaccardIndex(x, y []float64) (float64, error) {
	xLen, yLen := len(x), len(y)
	if xLen+yLen == 0 {
		return 1, nil
	}

	z := append([]float64{}, y...)
	sort.Float64s(z)
	var intersection int
	for i := range x {
		index := sort.SearchFloat64s(z, x[i])
		if index != yLen {
			intersection++
		}
	}
	return float64(intersection) / float64(xLen+yLen-intersection), nil
}
