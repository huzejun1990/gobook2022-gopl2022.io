// Package word provides utilities for word games.
// 包 word 提供了文字游戏的实用程序。
package word

import "unicode"

// IsPalindrome reports whether s reads the same forward and backward.
// IsPalindrome 报告 s 向前和向后读取是否相同。
// Letter case is ignored, as are non-letters.
// 字母大小写被忽略，非字母也是如此。
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//!-
