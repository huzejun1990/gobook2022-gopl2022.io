// Package intset provides a set of integers based on a bit vector.
// 包 intset 提供了一组基于位向量的整数。
package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers. // IntSet 是一组小的非负整数。
// Its zero value represents the empty set. // 它的零值代表空集。
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
// has 报告集合是否包含非负值 x。
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
// Add 将非负值 x 添加到集合中。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
// UnionWith 将 s 设置为 s 和 t 的并集。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string
// String returns the set as a string of the form "{1 2 3}".
// String 以“{1 2 3}”形式的字符串形式返回集合。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i*j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
