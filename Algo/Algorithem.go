package PushSwap

import (
	"fmt"
)

// Pushes the top first element of one stack to another (pa, pb)
func PushTop(exStack *[]int, impStack *[]int, name string) {
	if len(*exStack) == 0 {
		fmt.Println("Not enought elements")
		return
	}
	fmt.Println("p" + name)
	tempVar := (*exStack)[0]
	*exStack = append((*exStack)[:0], (*exStack)[0+1:]...)

	if len(*impStack) == 0 {
		*impStack = append(*impStack, tempVar)
		return
	}

	*impStack = append((*impStack)[:0+1], (*impStack)[0:]...)
	(*impStack)[0] = tempVar
}

// Swap first 2 elements of stack (sb,sa)
func Swap(stack *[]int, name string) {
	if len(*stack) < 2 {
		fmt.Println("Not enought elements")
		return
	}

	fmt.Println("s" + name)
	tempVar := (*stack)[0]
	(*stack)[0] = (*stack)[1]
	(*stack)[1] = tempVar
}

// Shifts up all elements of stack by 1 (ra, rb)
func Rotate(stack *[]int, name string) {
	if name != "bb" {
		fmt.Println("r" + name)
	}

	newStack := make([]int, len(*stack))

	newStack[len(*stack)-1] = (*stack)[0]

	for i := 1; i < len(*stack); i++ {
		newStack[i-1] = (*stack)[i]
	}

	*stack = newStack
}

// Shifts down all elements of stack by 1 (rra, rrb)
func ReverseRotate(stack *[]int, name string) {
	if name != "bb" {
		fmt.Println("rr" + name)
	}

	newStack := make([]int, len(*stack))

	newStack[0] = (*stack)[len(*stack)-1]

	for i := 1; i < len(*stack); i++ {
		newStack[i] = (*stack)[i-1]
	}

	*stack = newStack
}

// Executes swap function for both stacks (ss)
func SwapBoth(ThisA *[]int, ThisB *[]int) {
	Swap(ThisA, "a")
	Swap(ThisB, "b")
}

// Executes rotate function for both stacks (rr)
func RotateBoth(ThisA *[]int, ThisB *[]int) {
	fmt.Println("rr")
	Rotate(ThisA, "bb")
	Rotate(ThisB, "bb")
}

// Executes reverse rotate function for both stacks (rrr)
func ReverseRotateBoth(ThisA *[]int, ThisB *[]int) {
	fmt.Println("rrr")
	ReverseRotate(ThisA, "bb")
	ReverseRotate(ThisB, "bb")
}

func CheckRotate(stack []int, numberToAppend int, rotateType string) int {
	var counter int

	if len(stack) == 1 {
		return 0
	} else if stack[0] > numberToAppend && stack[len(stack)-1] < numberToAppend {
		return 0
	}

	var copyStack []int = make([]int, len(stack))
	copy(copyStack, stack)

	if rotateType == "r" {
		for {
			tempVar := copyStack[len(copyStack)-1]
			copyStack[len(copyStack)-1] = copyStack[0]

			for i := 1; i < len(copyStack)-1; i++ {
				copyStack[i-1] = copyStack[i]
			}

			copyStack[len(copyStack)-2] = tempVar
			counter++

			if copyStack[0] > numberToAppend && copyStack[len(copyStack)-1] < numberToAppend {
				break
			}
		}
	} else {
		for {
			tempVar := copyStack[0]
			copyStack[0] = copyStack[len(copyStack)-1]

			for i := len(copyStack) - 1; i > 1; i-- {
				copyStack[i] = copyStack[i-1]
			}

			copyStack[1] = tempVar
			counter++

			if copyStack[0] > numberToAppend && copyStack[len(copyStack)-1] < numberToAppend {
				break
			}
		}
	}

	return counter
}