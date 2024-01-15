package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"PushSwap/Algo"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	if os.Args[1] == "" {
		return
	}

	userInput := strings.Split(os.Args[1], " ")
	var A []int
	var B []int
	var sorted []int

	err := PushSwap.AppendNumbers(userInput, &A)
	if err {
		return
	}

	if sort.IntsAreSorted(A) {
		return
	}

	_ = B
	sorted = append(sorted, A...)
	sort.Ints(sorted)

	if len(A) == 1 {
		return
	} else if len(A) == 2 {
		if (A)[0] > (A)[1] {
			PushSwap.Swap(&A, "a")
			return
		}
	} else if len(A) == 3 {
		PushSwap.SmallSortA(&A)
		return
	}

	min, max, stackSize, median := PushSwap.FindValues(&A)

	if len(A) <= 6 {
		if A[0] > A[1] && A[1] == min {
			PushSwap.Swap(&A, "a")
		}

		for i := 0; i < stackSize/2; i++ {
			PushSwap.PushTop(&A, &B, "b")
		}

		PushSwap.SmallSortA(&A)
		PushSwap.SmallSortB(&B)
		PushSwap.MergeSmallStacks(&A, &B, max, min)
		return
	}

	if len(A) > 6 {
		var index int = 0
		for {
			if A[index] != min && A[index] != max {
				PushSwap.PushTop(&A, &B, "b")

				if B[0] > median && len(B) != 1 {
					PushSwap.Rotate(&B, "b")
				}
			} else {
				PushSwap.Rotate(&A, "a")
			}

			if len(A) == 2 {
				break
			}
		}
		PushSwap.FullSort(&A, &B)
	}
}
