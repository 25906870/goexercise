package strrelated

// LongestValidParentheses .
func LongestValidParentheses(s string) int {

	stack := make([]byte, 0, 1)
	slen := len(s)
	sum := 0
	maxsum := 0
	tmpsum := 0
	befend := 0
	tmptag := 0
	for index := 0; index < slen; index++ {
		if s[index] == '(' {

			if len(stack) == 0 {
				tmptag = index
				sum = 0
			}
			stack = append(stack, ')')

		} else if s[index] == ')' {

			if len(stack) < 1 {
				sum = 0
				tmptag = index + 1
				continue
			}
			stack = stack[:len(stack)-1]
			sum += 2

			if len(stack) == 0 {

				if befend+1 == tmptag {
					sum = sum + tmpsum

				}

				tmpsum = sum
				befend = index
			} else {

			}

			if sum > maxsum {
				maxsum = sum
			}

		} else {
			sum = 0
			stack = stack[0:0]
		}
	}
	return maxsum
}
