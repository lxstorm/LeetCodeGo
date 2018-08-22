package PalindromeNumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := 0
	tmpX := x
	for tmpX != 0 {
		tmp = tmp*10 + tmpX%10
		tmpX = tmpX / 10
	}
	if tmp == x {
		return true
	}
	return false

}
