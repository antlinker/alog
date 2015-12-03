package utils

import (
	"testing"
)

func TestParse(t *testing.T) {
	data := map[string]interface{}{
		"Foo": "Bar",
	}
	buf, err := NewParseTmpl("Foo={{.Foo}}").Parse(data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(buf.String())
}
