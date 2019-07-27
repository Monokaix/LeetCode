package main

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 { //说明没有最大公约数
		return ""
	}
	if len(str2) > len(str1) {
		return gcdOfStrings(str2, str1)
	}
	return str1[:gcd(len(str1), len(str2))]

}
//辗转相除求最大公约数
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
