package after_ops

func FinalValueAfterOperations(operations []string) int {
	var final int
	for _, op := range operations {
		switch op[1] {
		case '+':
			final++
		case '-':
			final--
		}
	}
	return final
}
