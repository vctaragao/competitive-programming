package generateparentheses

// Using the backtracking approach I have to think about
// 1. possibilites to be used in each step
// 2. The rules for witch I will follow in order to use a given set of possibilites
//
// In this problem I have two possibilites ["(", ")"]
// And the rules to follow are:
// 1. Base Case: open == close == n
// 2. Rule for ( : open < n
// 3. Rule for ) : close < open

func generateParenthesis(n int) []string {
	return solve(n, 0, 0, "", []string{})
}

func solve(n, open, closed int, str string, result []string) []string {
	if open == n && closed == n {
		result = append(result, str)
	}

	if open < n {
		str += "("
		result = solve(n, open+1, closed, str, result)
		str = str[:len(str)-1]
	}

	if closed < open {
		str += ")"
		result = solve(n, open, closed+1, str, result)
		str = str[:len(str)-1]
	}

	return result
}
