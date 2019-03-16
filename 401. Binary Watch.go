package main

import "fmt"

var arr = []int{1, 2, 4, 8, 16, 32}

func readBinaryWatch(num int) []string {
	res := make([]string,0)
	leds := make([]bool,10)
	dfs2(0,num,&res,leds)
	return res
}
func dfs2(index int, n int, res *[]string, leds []bool) {
	h := 0 //时针
	m := 0 //分针
	if n == 0 {
		m,h= cal(leds[:6]),cal(leds[6:])
		if h < 12 && m<60{ //时间要合法
			*res = append(*res,fmt.Sprintf("%d:%02d",h,m))
		}
		return
	}

	for i := index; i < n; i++ { //剪枝操作，很有学问，假如n为3，则实际可选的位置只有剩下的7个加3个中的一个，因此可以推算得到i<11-n
		leds[i] = true
		dfs2(i+1, n-1, res, leds)
		leds[i] = false
	}
}

func cal(leds []bool) int{
	sum := 0
	for i := 0; i < len(leds); i++ {
		if leds[i] {
			sum += arr[i]
		}
	}
	return sum
}
