package main

import (
	"math"
)

//方法一：暴力循环，因为前两重循环是找所有子串，第三重循环判断回文
//方法二：dp，其实思路和方法一一样，只不过判断回文的时候用dp记录了，所以复杂度是O(n^2)，dp[i][j]
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
	//逆序创建dp
	for i := n - 2; i >= 0; i-- {
		dp[i][i] = true
		for j := n - 1; j >= i; j-- {
			if j == i+1 { //当i、j相邻时,直接判断s[i]是否等于s[j]
				dp[i][j] = s[j] == s[i] //因为若s[i]!=s[j]则i到j组不成回文 举例abba，a和a相等，判断bb相等，这时是回文
			}
			if j > i+1 { //根据dp[i+1][j-1]判断当前i到j是否为回文
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			//当i到j是回文时,这里开始更新最长记录
			if j-i+1 > maxLen && dp[i][j] {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return s[start : maxLen+start]
}
//方法二
//马拉车算法分为两大部分：
//1、当前节点i在右边界mx外，直接暴力扩散
//2、当前节点i在右边界内：
//（1）i的对称点i在最大左右边界L内，则i的回文串长度为i的回文串长度
//（2）i的对称点i在最大左右边界L外，则i的回文串长度为到R的距离。
//（3）i的对称点i在最大左右边界L上，则i的回文串要计算R之后的那部分是否是回文串
func Manacher(s string) string {
	if s == "" {
		return ""
	}
	str := "$#"
	for i := 0; i < len(s); i++ {
		str += string(s[i])
		str += "#"
	}
	str += "@" //放置越界,因为不加的话第一个和最后一个都是#,会重复判断
	p := make([]int, len(str))
	id := 0 //最大回文半径的中心点
	mx := 0 //右边界
	resCenter := 0
	resLen := 0
	for i := 1; i < len(str)-1; i++ {
		if mx > i { //当前i在右边界内
			p[i] = int(math.Min(float64(p[2*id-i]), float64(mx-i))) //p[2*id-i]表示关于id对称位置处的回文半径
		} else {
			p[i] = 1
		}
		for i-p[i] >= 0 && i+p[i] < len(str)-1 && str[i+p[i]] == str[i-p[i]] { //i位置还可以继续扩列 对应i在右边界mx外和i的对称点i在最大左右边界L上时的暴力扩
			p[i]++
		}
		if i+p[i] > mx { //有更大的右边界
			id = i
			mx = i + p[i]
		}
		if p[i] > resLen { //当前位置i的回文半径更长
			resLen = p[i]
			resCenter = i
		}
	}
	//(resCenter-resLen)/2为起始索引,resLen-1为长度
	return s[(resCenter-resLen)/2 : (resCenter-resLen)/2+resLen-1]

}


