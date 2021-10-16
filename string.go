package dimea

import "math"

/**********
Levenshtein
**********/

type levenshtein struct {
	cost struct {
		insert  int
		delete  int
		replace int
	}
}

func NewLevenshtein() *levenshtein {
	return &levenshtein{cost: struct {
		insert  int
		delete  int
		replace int
	}{insert: 1, delete: 1, replace: 1}}
}

func (l *levenshtein) SetCosts(insert, delete, replace int) *levenshtein {
	l.cost.insert = insert
	l.cost.delete = delete
	l.cost.replace = replace
	return l
}

func (l *levenshtein) SetInsertCost(insert int) *levenshtein {
	l.cost.insert = insert
	return l
}

func (l *levenshtein) SetDeleteCost(delete int) *levenshtein {
	l.cost.delete = delete
	return l
}

func (l *levenshtein) SetReplaceCost(replace int) *levenshtein {
	l.cost.replace = replace
	return l
}

func (l *levenshtein) Distance(x, y string) int {
	xRune, yRune := []rune(x), []rune(y)

	// initialize table
	table := make([][]int, len(xRune)+1)
	for i := 0; i < len(xRune)+1; i++ {
		table[i] = make([]int, len(yRune)+1)
	}
	// initialize row
	for i := 0; i < len(xRune)+1; i++ {
		table[i][0] = i * l.cost.delete
	}
	// initialize column
	for i := 0; i < len(yRune)+1; i++ {
		table[0][i] = i * l.cost.insert
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
				costs = append(costs, table[i-1][j-1]+l.cost.replace)
			}
			costs = append(costs, table[i-1][j]+l.cost.insert)
			costs = append(costs, table[i][j-1]+l.cost.delete)

			table[i][j] = min(costs...)
		}
	}

	return table[len(xRune)][len(yRune)]
}

func (l *levenshtein) StdDistance(x, y string) float64 {
	d := l.Distance(x, y)
	return float64(d) / math.Max(float64(len([]rune(x))), float64(len([]rune(y))))
}
