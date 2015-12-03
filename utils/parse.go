package utils

import (
	"bytes"
	"text/template"
)

// ParseTmpl 提供模板解析接口
type ParseTmpl interface {
	// Parse 提供模板解析函数
	Parse(interface{}) (*bytes.Buffer, error)
}

// NewParseTmpl 创建新的ParseTmpl实例
func NewParseTmpl(tmpl interface{}) ParseTmpl {
	var t *template.Template
	if v, ok := tmpl.(string); ok {
		t = template.Must(template.New("").Parse(v))
	} else if v, ok := tmpl.(*template.Template); ok {
		t = v
	}
	return &parse{tmpl: t}
}

type parse struct {
	tmpl *template.Template
}

func (p *parse) Parse(data interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := p.tmpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
