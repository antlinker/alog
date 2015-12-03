package utils

import (
	"testing"
)

type User struct {
	ID   int64  `json:"id";yaml:"id"`
	Name string `json:"name";yaml:"name"`
}

func TestReadJson(t *testing.T) {
	cfg := NewConfig("config.json")
	var user User
	err := cfg.Read(&user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestReadYaml(t *testing.T) {
	cfg := NewConfig("config.yaml")
	var data map[string]interface{}
	err := cfg.Read(&data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}
