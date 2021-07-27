package main

//dp 用cost[i][j]记录word1[0...i-1]变换到word2[0...j-1]的消耗
//有四种情况到达word1[i]和word2[j]
//1、已知word1[i-1]到word2[j-1]的消耗
//2、已知word1[i-1]到word2[j]的消耗
//3、word1[i]到word2[j-1]的消耗
//4、word1[i]直接到word2[j]，则消耗不变，即word1[i]==word2[j]，cost[i][j] = cost[i-1][j-1]
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	cost := make([][]int, 0)
	for i := 0; i <= m; i++ {
		tmp := make([]int, n+1)
		cost = append(cost, tmp)
	}
	//初始化从word1[0...i]到word2[0]每次的消耗
	for i := 0; i <= m; i++ {
		cost[i][0] = i
	}
	//初始化从word1[0]到word2[0...j]每次的消耗
	for j := 0; j <= n; j++ {
		cost[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] { //上文总结的第4种情况
				cost[i][j] = cost[i-1][j-1]
			} else { //其余三种情况
				cost[i][j] = 1 + min(cost[i-1][j-1], min(cost[i-1][j], cost[i][j-1]))
			}
		}
	}
	return cost[m][n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
