// Package word provides utilities for word games.
// 包 word 提供了文字游戏的实用程序。
package word

// IsPalindrome reports whether s reads the same forward and backward.
// IsPalindrome 报告 s 向前和向后读取是否相同。
// (Our first attempt.)
//（我们的第一次尝试。）
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
