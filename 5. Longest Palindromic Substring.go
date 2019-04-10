package main

//方法一：暴力循环，因为前两重循环是找所有子串，第三重循环判断回文
//方法二：dp，其实思路和方法一一样，只不过判断回文的时候用dp记录了，所以复杂度收益O(n)，dp[i][j]
//记录从i到j是否为回文，若dp[i+1][j-1]是回文，则str[i]=str[j]时dp[i][j]也是回文，利用这一特性求解
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start := 0
	maxLen := 1
	dp := make([][]bool, 0)
	for i := 0; i < n; i++ {
		tmp := make([]bool, n)
		dp = append(dp, tmp)
	}
	dp[n-1][n-1] = true
	res := []byte{}
	//逆序创建dp
	for i := n - 2; i >= 0; i-- {
		dp[i][i] = true
		for j := n - 1; j >= i; j-- {
			if j == i+1 {
				dp[i][j] = s[j] == s[i] //因为若s[i]!=s[j]则i到j组不成回文 举例abba，a和a相等，判断bb相等，这时是回文
			}
			if j > i+1 { //根据dp[i+1][j-1]判断当前i到j是否为回文
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if j-i+1 > maxLen && dp[i][j] {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	//fmt.Println("start:",start)
	//fmt.Println("maxlen:",maxLen)
	for i := start; i < maxLen+start; i++ {
		res = append(res, s[i])
	}
	return string(res)
}