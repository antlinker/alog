package alog

import (
	"errors"

	"gopkg.in/alog.v1/log"
	"gopkg.in/alog.v1/utils"
)

// 解析配置文件
func parseConfig(config interface{}) (*log.LogConfig, error) {
	cfg := new(log.LogConfig)
	if config != nil {
		if v, ok := config.(string); ok {
			err := utils.NewConfig(v).Read(cfg)
			if err != nil {
				return nil, err
			}
		} else if v, ok := config.(log.LogConfig); ok {
			cfg = &v
		} else if v, ok := config.(*log.LogConfig); ok {
			cfg = v
		} else {
			return nil, errors.New("Wrong configuration.")
		}
	}
	return cfg, nil
}

// 加载配置文件默认参数
func loadDefaultConfig(config *log.LogConfig) {
	if (*config).Global.IsEnabled == 0 {
		(*config).Global.IsEnabled = log.DefaultEnabled
	}
	if (*config).Console.Item.Tmpl == "" {
		(*config).Console.Item.Tmpl = log.DefaultConsoleTmpl
	}
	if (*config).Console.Item.TimeTmpl == "" {
		(*config).Console.Item.TimeTmpl = log.DefaultConsoleTimeTmpl
	}
	if (*config).Global.Interval == 0 {
		(*config).Global.Interval = log.DefaultInterval
	}
	if (*config).Global.Buffer.Engine == 0 {
		(*config).Global.Buffer.Engine = log.MEMORY_BUFFER
	}
	if (*config).Global.TargetStore == "" {
		(*config).Global.TargetStore = log.DefaultGlobalKey
	}
	if (*config).Store.File == nil {
		(*config).Store.File = map[string]log.FileConfig{
			log.DefaultGlobalKey: log.FileConfig{},
		}
	}
	if (*config).Global.FileCaller == 0 {
		(*config).Global.FileCaller = log.DefaultFileCaller
	}
}

func defaultConfig() *log.LogConfig {
	config := new(log.LogConfig)
	loadDefaultConfig(config)
	return config
}
