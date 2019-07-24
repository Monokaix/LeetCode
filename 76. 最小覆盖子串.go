package main

//两个指针滑动窗口,每个窗口维护一个当前的解,每次比较得到最优值
//何时i移动:只要给定的串T不在当前窗口,i就一直移动,同时记录每个字符出现的次数
//何时begin移动:两种情况,前提是begin<i
//1、begin当前指向的字符不在给定的串T中
//2、begin当前指向的字串出现在串T中,但当前窗口中该字符数量大于串T中对应的字符
func minWindow(s string, t string) string {
	begin := 0
	res := ""
	//计算t中每个字符出现次数
	map_t := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		map_t[t[i]]++
	}
	map_cur := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		map_cur[s[i]]++
		for begin < i {
			ch := s[begin]
			if map_t[ch] == 0 { //当前字符不在t中
				begin++
			} else if map_t[ch] < map_cur[ch] { //当前字符在t中,但当前字符个数大于t中对应字符个数
				map_cur[ch]--
				begin++
			} else { //得到一个新的解
				break
			}
		}
		if helper2(map_t, map_cur) {
			if res == "" || len(res) > i-begin+1 {
				res = s[begin : i+1]
			}
		}
	}
	return res
}

//该函数功能：判断map_s中的字符是否存在于map_t中
func helper2(map_s map[byte]int, map_t map[byte]int) bool {
	for k, v := range map_s {
		if v > map_t[k] {
			return false
		}
	}
	return true
}