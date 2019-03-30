package main

//方法一：自顶向下，记忆化搜索
func numDecodings(s string) int {
	memo := make([]int, 0)
	for i := 0; i <= len(s); i++ {
		memo = append(memo, -1)
	}
	return dp(s, len(s), memo)
}
func dp(s string, n int, memo []int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		if s[0] == '0' {
			return 0
		}
		return 1
	}
	if memo[n] != -1 { //记忆化搜索
		return memo[n]
	}
	res := 0
	//相当于f(n-1) 斐波那契
	if s[n-1] != '0' {
		res += dp(s, n-1, memo)
	}
	//相当于f(n-2),只不过要分情况讨论,大于等于20小于26或者小于20
	if s[n-2] == '1' || s[n-2] == '2' && s[n-1] <= '6' {
		res += dp(s, n-2, memo)
	}
	memo[n] = res
	return res
}

//方法二：动态规划
func numDecodings2(s string) int {
	n := len(s)
	if s == "0" || n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			dp[i+1] = 0
		} else {
			dp[i+1] = dp[i]
		}
		if i > 0 && (s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6') {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[n]
}
