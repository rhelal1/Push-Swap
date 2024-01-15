package PushSwap

import (
	"reflect"
	"sort"
)

func SmallSortA(Stack *[]int) {
	min, max, sLenght, median := FindValues(Stack)
	if sort.IntsAreSorted(*Stack) {
		return
	}

	if sLenght == 2 {
		if (*Stack)[0] > (*Stack)[1] {
			Swap(Stack, "a")
			return
		}
	}

	if (*Stack)[0] == max {
		Rotate(Stack, "a")
	}

	if (*Stack)[0] == median {
		if (*Stack)[1] == min {
			Swap(Stack, "a")
		} else {
			ReverseRotate(Stack, "a")
		}
		return
	}

	if (*Stack)[1] == max {
		ReverseRotate(Stack, "a")
		Swap(Stack, "a")
	}
}

func SmallSortB(Stack *[]int) {
	min, max, sLenght, _ := FindValues(Stack)

	var sortedValues []int
	sortedValues = append(sortedValues, *Stack...)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedValues)))

	if reflect.DeepEqual(*Stack, sortedValues) {
		return
	}

	if sLenght == 2 {
		if (*Stack)[0] < (*Stack)[1] {
			Swap(Stack, "b")
			return
		}
	}

	if (*Stack)[2] == max {
		if (*Stack)[1] == min {
			ReverseRotate(Stack, "b")
		} else {
			Rotate(Stack, "b")
			Swap(Stack, "b")
		}

		return
	}

	if (*Stack)[1] == max {
		if (*Stack)[0] == min {
			Swap(Stack, "b")
		} else {
			Rotate(Stack, "b")
		}

		return
	}

	ReverseRotate(Stack, "b")
	Swap(Stack, "b")
}

func MergeSmallStacks(ThisA *[]int, ThisB *[]int, max int, min int) {
	for {
		if len(*ThisB) == 0 {
			break
		}

		var counter int
		if (*ThisB)[0] == max {
			PushTop(ThisB, ThisA, "a")
			Rotate(ThisA, "a")
			continue
		} else if (*ThisB)[0] == min {
			PushTop(ThisB, ThisA, "a")
			continue
		}

		if (*ThisB)[0] < (*ThisA)[0] {
			PushTop(ThisB, ThisA, "a")
		} else {
			for i := 0; i < len(*ThisA); i++ {
				if i == 0 {
					if (*ThisA)[0] > (*ThisB)[0] && (*ThisA)[len(*ThisA)-1] < (*ThisB)[0] {
						counter = i + 1
						break
					}
				} else if i == len(*ThisA)-1 {
					if (*ThisA)[0] < (*ThisB)[0] && (*ThisA)[len(*ThisA)-1] > (*ThisB)[0] {
						counter = i + 1
						break
					}
				} else if (*ThisA)[i] < (*ThisB)[0] && (*ThisA)[i+1] > (*ThisB)[0] {
					counter = i + 1
					break
				} else if (*ThisA)[i-1] < (*ThisB)[0] && (*ThisA)[i] > (*ThisB)[0] {
					counter = i - 1
					break
				}
			}

			for a := 1; a < len(*ThisA); a++ {
				if (counter)-a == 0 || counter == 0 {
					for b := 0; b < a; b++ {
						Rotate(ThisA, "a")
					}

					PushTop(ThisB, ThisA, "a")

					if len(*ThisB) != 0 {
						if (*ThisB)[0] < (*ThisA)[0] && (*ThisB)[0] > (*ThisA)[len(*ThisA)-1] {
							PushTop(ThisB, ThisA, "a")

						}
					}

					for b := 0; b < a; b++ {
						ReverseRotate(ThisA, "a")
					}
					break
				} else if counter+a == len(*ThisA) || counter == len(*ThisA) {
					for b := 0; b < a; b++ {
						ReverseRotate(ThisA, "a")
					}

					PushTop(ThisB, ThisA, "a")

					if len(*ThisB) != 0 {
						if (*ThisB)[0] < (*ThisA)[0] && (*ThisB)[0] > (*ThisA)[len(*ThisA)-1] {
							PushTop(ThisB, ThisA, "a")
							a++
						}
					}

					for b := 0; b < a+1; b++ {
						Rotate(ThisA, "a")
					}
					break
				}
			}
		}
	}
}
