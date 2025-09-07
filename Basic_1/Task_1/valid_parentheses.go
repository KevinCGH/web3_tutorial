package Task_1

func ValidParentheses(s string) bool {
	n := len(s)
	// 奇数个括号不可能配对成功
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}

	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			// 栈顶括号是否配对
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
