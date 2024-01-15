package PushSwap

func FullSort(ThisA *[]int, ThisB *[]int) {
	for {
		if len(*ThisB) == 0 {
			break
		}

		var scoreList []ScoreData = make([]ScoreData, len(*ThisB))

		var steps int
		var sLenght int = len(scoreList)
		var changeDirection bool
		for i := 0; i < sLenght; i++ {
			scoreList[i].SC4B = steps

			if sLenght%2 == 0 {
				if sLenght/2 > i {
					steps++
				} else {
					steps--
				}
			} else {
				if i == (sLenght/2)+1 {
					scoreList[i].SC4B = steps - 1
					steps--
				}

				if sLenght/2 >= i {
					steps++
				} else {
					steps--
				}
			}

			if (i == (sLenght/2)+1 && sLenght%2 == 1) || (i == (sLenght/2) && sLenght%2 == 0) {
				changeDirection = true
			}

			if !changeDirection {
				scoreList[i].TMIB = "r"
			} else {
				scoreList[i].TMIB = "rr"
			}
		}

		var reverseRotateResult, rotateResult, bestResult, bestResultIndex int
		for i := 0; i < len(*ThisB); i++ {
			rotateResult = CheckRotate(*ThisA, (*ThisB)[i], "r")
			reverseRotateResult = CheckRotate(*ThisA, (*ThisB)[i], "rr")

			if rotateResult <= reverseRotateResult {
				scoreList[i].SC4A = rotateResult
				scoreList[i].TMIA = "r"
			} else {
				scoreList[i].SC4A = reverseRotateResult
				scoreList[i].TMIA = "rr"
			}

			if rotateResult == reverseRotateResult {
				if (scoreList[i].TMIA != scoreList[i].TMIB) && scoreList[i].TMIB == "rr" {
					scoreList[i].SC4A = reverseRotateResult
					scoreList[i].TMIA = "rr"
				}
			}

			possibleBestResult := scoreList[i].SC4A + scoreList[i].SC4B
			if bestResult == 0 || (bestResult > possibleBestResult) {
				bestResult = scoreList[i].SC4A + scoreList[i].SC4B
				bestResultIndex = i
			}
		}

		var stepsLeftInA, stepsLeftInB int = scoreList[bestResultIndex].SC4A, scoreList[bestResultIndex].SC4B
		for {
			if stepsLeftInA == 0 && stepsLeftInB == 0 {
				PushTop(ThisB, ThisA, "a")
				break
			}

			if scoreList[bestResultIndex].TMIA == scoreList[bestResultIndex].TMIB {
				if stepsLeftInA != 0 && stepsLeftInB != 0 {
					if scoreList[bestResultIndex].TMIA == "r" {
						RotateBoth(ThisA, ThisB)
					} else {
						ReverseRotateBoth(ThisA, ThisB)
					}
					stepsLeftInA--
					stepsLeftInB--
					continue
				}
			}

			if stepsLeftInA != 0 {
				FullSortMove(scoreList[bestResultIndex].TMIA, ThisA, &stepsLeftInA, "a")
			}

			if stepsLeftInB != 0 {
				FullSortMove(scoreList[bestResultIndex].TMIB, ThisB, &stepsLeftInB, "b")
			}
		}
	}

	min, _, stackSize, _ := FindValues(ThisA)
	var RotateType string
	var stepsCount int
	for i := 0; i < stackSize; i++ {
		if (*ThisA)[i] == min {
			if stackSize%2 == 0 {
				if stackSize/2 <= i {
					RotateType = "rr"
					stepsCount = stackSize - i
				} else {
					RotateType = "r"
					stepsCount = i
				}
			} else {
				if stackSize/2 < i {
					RotateType = "rr"
					stepsCount = stackSize - i
				} else {
					RotateType = "r"
					stepsCount = i
				}
			}
			break
		}
	}

	if RotateType == "r" {
		for i := 0; i < stepsCount; i++ {
			Rotate(ThisA, "a")
		}
	} else {
		for i := 0; i < stepsCount; i++ {
			ReverseRotate(ThisA, "a")
		}
	}
}

func FullSortMove(typeMove string, stack *[]int, stepsLeft *int, name string) {
	if typeMove == "r" {
		Rotate(stack, name)
	} else {
		ReverseRotate(stack, name)
	}
	*stepsLeft--
}
