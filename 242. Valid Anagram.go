package main

//using map

func isAnagram(s string, t string) bool {
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if _, ok := map1[s[i]]; ok {
			map1[s[i]]++
		} else {
			map1[s[i]] = 0
		}

		if _, ok := map2[t[i]]; ok {
			map2[t[i]]++
		} else {
			map2[t[i]] = 0
		}
	}
	for key := range map1 {
		 _, ok := map2[key]
		 if !ok || map1[key] != map2[key] {
			return false
		}
	}

	return true
}
func main() {
	print(isAnagram("anagram", "nagaram"))
}
