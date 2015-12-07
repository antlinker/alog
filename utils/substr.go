package utils

import "strings"

// SubstrByStartAfter 字符串截取
// 从尾部开始查找，返回查找位置到尾部的字符串
func SubstrByStartAfter(str string, start string) string {
	pos := strings.LastIndex(str, start)

	if pos > -1 {
		pos += len(start)
		return string([]byte(str)[pos:])
	}
	return str
}

// SubStrByStartBefore 字符串截取
// 从尾部开始查找，返回从头到查找位置
func SubStrByStartBefore(str string, start string) string {
	pos := strings.LastIndex(str, start)

	if pos > -1 {
		return string([]byte(str)[0:pos])
	}
	return str
}
