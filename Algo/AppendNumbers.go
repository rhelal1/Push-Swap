package PushSwap

import (
	"fmt"
	"strconv"
)


// Append all number to A stack
func AppendNumbers(userInput []string, A *[]int) bool {
	for i := 0; i < len(userInput); i++ {
		numToAppend, err := strconv.Atoi(userInput[i])
		if err != nil {
			fmt.Println("Error")
			return true
		}

		// Check for duplicates
		if i != 0 {
			for k := 0; k < len(*A); k++ {
				if (*A)[k] == numToAppend {
					fmt.Println("Error")
					return true
				}
			}
		}

		(*A) = append((*A), numToAppend)
	}
	return false
}
