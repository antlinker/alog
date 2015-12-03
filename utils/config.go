package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config 读取配置文件
type Config interface {
	// Read 将配置文件信息解析到指定的对象
	Read(interface{}) error
}

// NewConfig 创建Config实例
func NewConfig(file string) Config {
	if !filepath.IsAbs(file) {
		fPath, err := filepath.Abs(file)
		if err == nil {
			file = fPath
		}
	}
	return &config{file: file}
}

type config struct {
	file string
}

func (c *config) Read(v interface{}) error {
	buf, err := c.readFile()
	if err != nil {
		return err
	}
	switch strings.ToLower(filepath.Ext(c.file)) {
	case ".yaml":
		return yaml.Unmarshal(buf.Bytes(), v)
	case ".json":
		return json.Unmarshal(buf.Bytes(), v)
	}
	return nil
}

func (c *config) readFile() (*bytes.Buffer, error) {
	file, err := os.Open(c.file)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	return buf, nil
}
