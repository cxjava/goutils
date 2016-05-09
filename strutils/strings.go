package strutils

import (
	"strings"
	"unicode/utf8"
)

func TrimSpace(s string) string {
	return strings.TrimRight(strings.TrimLeft(s, "\r\n\t "), "\r\n\t ")
}

func SubStringBetween(str, begin, end string) string {
	if len(str) < 1 || len(begin) < 1 || len(end) < 1 {
		return ""
	}

	if strings.Index(str, begin) < 0 {
		return ""
	}
	exp := strings.Split(str, begin)
	if len(exp) <= 1 {
		return ""
	}

	sMember := exp[1]
	if strings.Index(sMember, end) < 0 {
		return ""
	}
	exp = strings.Split(sMember, end)
	if len(exp) == 1 {
		return ""
	}
	return exp[0]
}

func SubStr(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	return utf8.RuneCountInString(str[:result])
}
