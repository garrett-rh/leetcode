package fizz

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) []string {
	var (
		b      strings.Builder
		answer []string
	)
	for i := 1; i <= n; i++ {

		if i%3 == 0 {
			b.WriteString("Fizz")
		}
		if i%5 == 0 {
			b.WriteString("Buzz")
		}
		if len(b.String()) == 0 {
			b.WriteString(strconv.Itoa(i))
		}

		answer = append(answer, b.String())

		b.Reset()
	}

	return answer
}
