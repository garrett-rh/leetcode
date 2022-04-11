package baseball

import (
	"strconv"
)

func CalPoints(ops []string) int {
	if len(ops) == 1 {
		op, _ := strconv.Atoi(ops[0])
		return op
	}
	var (
		score  = make([]int, len(ops))
		top    int
		result int
	)
	for _, op := range ops {
		switch op {
		case "D":
			score[top] = 2 * score[top-1]
			result += score[top]
			top++
		case "+":
			score[top] = score[top-1] + score[top-2]
			result += score[top]
			top++
		case "C":
			result -= score[top-1]
			top--
		default:
			intOp, _ := strconv.Atoi(op)
			score[top] = intOp
			result += score[top]
			top++
		}

	}

	return result
}
