package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func SubstrByStartAfter(str string, start string) string {

	pos := strings.LastIndex(str, start)

	if pos > -1 {
		pos += len(start)
		return string([]byte(str)[pos:])
	}
	return str

}
func SubStrByStartBefore(str string, start string) string {
	pos := strings.LastIndex(str, start)

	if pos > -1 {
		return string([]byte(str)[0:pos])
	}
	return str
}

type pkgpath struct {
}

func GetGoPath() string {
	t := reflect.TypeOf(pkgpath{})
	_, wfile, _, _ := runtime.Caller(0)
	return SubStrByStartBefore(wfile, t.PkgPath())

}
