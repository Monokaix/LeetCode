package main

func reverseLeftWords(s string, n int) string {
	len := len(s)
	res := reverseBetween([]byte(s), 0, len-1)
	res = reverseBetween(res, len-n, len-1)
	res = reverseBetween(res, 0, len-n-1)

	return string(res)
}

func reverseBetween(s []byte, i, j int) []byte {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}
