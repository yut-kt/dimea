package dimea

import (
	"fmt"
	"math"
)

// Hamming distance between two strings of equal length is the number of positions at which the corresponding symbols are different.
// https://en.wikipedia.org/wiki/Hamming_distance
func Hamming(x, y string) (int, error) {
	xRune, yRune := []rune(x), []rune(y)
	if len(xRune) != len(yRune) {
		return 0, fmt.Errorf("lengtn is different between x[%d] and y[%d]", len(xRune), len(yRune))
	}

	var sum int
	for i := range xRune {
		if xRune[i] != yRune[i] {
			sum++
		}
	}
	return sum, nil
}

/******************************
Levenshtein
******************************/

// Levenshtein distance is a string metric for measuring the difference between two sequences.
// Informally, the Levenshtein distance between two words is the minimum number of single-character edits
// (insertions, deletions or substitutions) required to change one word into the other.
// https://en.wikipedia.org/wiki/Levenshtein_distance
type Levenshtein struct {
	cost struct {
		insertions  int
		deletion  int
		substitution int
	}
}

// SetCosts is set edit (insertions, deletions or substitutions) costs.
func (l *Levenshtein) SetCosts(insertions, deletion, substitution int) *Levenshtein {
	l.cost.insertions = insertions
	l.cost.deletion = deletion
	l.cost.substitution = substitution
	return l
}

// SetInsertionsCost is set edit (insertions) cost.
func (l *Levenshtein) SetInsertionsCost(insertions int) *Levenshtein {
	l.cost.insertions = insertions
	return l
}

// SetDeletionCost is set edit (deletions) cost.
func (l *Levenshtein) SetDeletionCost(deletion int) *Levenshtein {
	l.cost.deletion = deletion
	return l
}

// SetSubstitutionCost is set edit (substitutions) cost.
func (l *Levenshtein) SetSubstitutionCost(substitution int) *Levenshtein {
	l.cost.substitution = substitution
	return l
}

// Distance is calc levenshtein distance.
func (l *Levenshtein) Distance(x, y string) int {
	xRune, yRune := []rune(x), []rune(y)

	// initialize table
	table := make([][]int, len(xRune)+1)
	for i := 0; i < len(xRune)+1; i++ {
		table[i] = make([]int, len(yRune)+1)
	}
	// initialize row
	for i := 0; i < len(xRune)+1; i++ {
		table[i][0] = i * l.cost.deletion
	}
	// initialize column
	for i := 0; i < len(yRune)+1; i++ {
		table[0][i] = i * l.cost.insertions
	}

	min := func(nums ...int) int {
		n := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < n {
				n = nums[i]
			}
		}
		return n
	}

	for i := 1; i < len(xRune)+1; i++ {
		for j := 1; j < len(yRune)+1; j++ {
			var costs []int
			if xRune[i-1] == yRune[j-1] {
				costs = append(costs, table[i-1][j-1])
			} else {
				costs = append(costs, table[i-1][j-1]+l.cost.substitution)
			}
			costs = append(costs, table[i-1][j]+l.cost.insertions)
			costs = append(costs, table[i][j-1]+l.cost.deletion)

			table[i][j] = min(costs...)
		}
	}

	return table[len(xRune)][len(yRune)]
}

// StdDistance is calc standardized levenshtein distance.
func (l *Levenshtein) StdDistance(x, y string) float64 {
	d := l.Distance(x, y)
	return float64(d) / math.Max(float64(len([]rune(x))), float64(len([]rune(y))))
}
