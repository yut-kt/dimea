package main

import (
	"fmt"
	"github.com/yut-kt/dimea"
)

func main() {
	fmt.Println(dimea.NewLevenshtein().SetCosts(7, 7, 10).Distance("agent", "agency"))
}
