package PushSwap

import (
	"sort"
)

func FindValues(Stack *[]int) (int, int, int, int) {
	var sortedStack []int
	sortedStack = append(sortedStack, *Stack...)
	sort.Ints(sortedStack)

	var min int = sortedStack[0]
	var max int = sortedStack[len(sortedStack)-1]
	var sLenght int = len(*Stack)
	var median int = (sortedStack)[sLenght/2]

	return min, max, sLenght, median
}
