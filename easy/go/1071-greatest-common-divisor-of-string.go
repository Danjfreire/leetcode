
import "strings"

// O(min(n,m) * (m+n))
func gcdOfStrings(str1 string, str2 string) string {
	var minStr string

	if len(str1) > len(str2) {
		minStr = str2
	} else {
		minStr = str1
	}

	for len(minStr) > 0 {
		if isGcd(str1, str2, minStr) {
			break
		}

		minStr = minStr[:len(minStr) -1]
	}

	return minStr 
}

func isGcd(str1 string, str2 string, gcd string) bool{
	len1 := len(str1)
	len2 := len(str2)
	gcdLen := len(gcd)

	if len1 % gcdLen > 0 || len2 % gcdLen < 0 {
		return false
	}

	divideStr1 := strings.Repeat(gcd, len1 / gcdLen) == str1
	divideStr2 := strings.Repeat(gcd, len2 / gcdLen) == str2

	return divideStr1 && divideStr2
}