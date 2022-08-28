package sol

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ans := []string{}
	var backTrack func(cur string, openCount int, closeCount int)
	backTrack = func(cur string, openCount int, closeCount int) {
		if len(cur) == 2*n {
			ans = append(ans, cur)
			return
		}
		if openCount < n {
			cur += "("
			backTrack(cur, openCount+1, closeCount)
			cur = cur[:len(cur)-1]
		}
		if closeCount < openCount {
			cur += ")"
			backTrack(cur, openCount, closeCount+1)
			cur = cur[:len(cur)-1]
		}
	}
	backTrack("", 0, 0)
	return ans
}
