# ALog

> 基于Golang的异步日志包

## Installation and usage

``` bash
$ go get gopkg.in/alog.v1
```

## API documentation

* [https://godoc.org/gopkg.in/alog.v1](https://godoc.org/gopkg.in/alog.v1)

## Configuration file

``` yaml
{
	# 控制台输出配置
	console: {
		# 输出日志级别（1表示Debug,2表示Info,3表示Warn,4表示Error,5表示Fatal）
  		# 0表示输出所有级别
		level: 0,
		# 日志项模板
		item: {
			# 项模板
		    # 模板字段说明：
		    # ID 唯一标识
		    # Time 日志发生时间
		    # Level 级别
		    # Tag 标签
		    # Message 日志明细
		    # FileName 文件名
		    # ShortName 短文件名
		    # FileFuncName 函数名
		    # FileLine 文件行
			tmpl: "[{{.ID}} {{.Time}} {{.Level}} {{.Tag}}] {{.Message}}",
			# 时间模板
		    # 模板字段说明：
		    # Year 年份
		    # Month 月份
		    # Day 天数
		    # Hour 小时
		    # Minute 分钟
		    # Second 秒
		    # MilliSecond 毫秒
			time: "{{.Hour}}:{{.Minute}}"
		}
	},
	# 全局配置
	global: {
		# 是否打印到控制台
	  	# 0表示不打印
	  	# 1表示打印
	  	print: 0,
	  	# 日志输出规则
	  	# 参数说明：
	  	# 0表示所有配置输出
	  	# 1表示指定Global配置输出
	  	# 2表示指定Tag配置输出
	  	# 3表示指定Level配置输出
	  	rule: 1,
	  	# 输出日志文件信息
	  	# 0表示不输出
	  	# 1表示输出
	  	showfile: 0,
	  	# 文件信息调用层级
	  	# 默认为5当前调用)
	  	caller: 5,
	  	# 读取缓冲区时间间隔（以秒为单位）
	  	interval: 1,
	  	# 目标存储
	  	# 指向store中定义的存储配置
	  	target: "file_global",
	  	# 缓冲区存储
	  	buffer: {
	  		# 存储引擎
		    # 1表示内存存储
		    # 2表示redis存储
		    engine: 1,
		    # 指向store中定义的存储配置
		    target: "redis_buffer"
	  	}
	},
	# 标签配置
	tags: [{
		# 标签名
  		names: ["ALOG","TEST"],
  		# 配置
  		config: {
  			# 是否打印到控制台
		    # 0表示不打印
		    # 1表示打印
		    print: 0,
		    # 日志级别
		    level: 0,
		    # 目标存储
		    # 指向store中定义的存储配置
		    target: ""
  		}
	}],
	# 级别配置
	levels: [{
		# 标签名
  		values: [4,5],
  		# 配置
  		config: {
  			# 是否打印到控制台
		    # 0表示不打印
		    # 1表示打印
		    print: 0,
		    # 目标存储
		    # 指向store中定义的存储配置
		    target: ""
  		}
	}],
	# 持久化存储配置
	store: {
		# redis 存储
		redis: {
			# 唯一标识键
			redis_buffer: {
				# host:port
				addr: "192.168.33.70:6379",
				# db
      			db: 0,
      			# poolsize
      			# Default is 10 connections.
      			poolsize: 10
			}
		},
		# 文件存储
		file: {
			# 唯一标识键
			file_global: {
			  # 文件存储路径
		      filepath: "logs",
		      # 文件名格式模板
		      # 模板字段说明：
		      # Year 年份
		      # Month 月份
		      # Day 天数
		      # Level 日志级别
		      # Tag 标签
		      filename: "{{.Year}}{{.Month}}{{.Day}}.log",
		      # 单个文件大小（单位KB）
		      filesize: 2048,
		      item: {
		      	# 项模板
        		# 字段说明同上
        		tmpl: '{{.ID}} {{.Time}} {{.Level}} {{.Tag}} "{{.FileName}} {{.FileFuncName}} {{.FileLine}}" {{.Message}}',
        		# 时间模板
		        # 字段说明同上
		        time: "{{.Year}}-{{.Month}}-{{.Day}} {{.Hour}}:{{.Minute}}:{{.Second}}.{{.MilliSecond}}"
		      }
			}
		}
	}
}
```

## Sample

``` go
package main

import (
	"time"

	"gopkg.in/alog.v1"
)

func main() {
	alog.RegisterAlog("config.yaml")
	alog.SetLogTag("Sample")
	alog.Debug("Debug info...")
	alog.DebugC("Debug console info...")
	alog.Info("Info info...")
	alog.InfoC("Info console info...")
	alog.Warn("Warn info...")
	alog.WarnC("Warn console info...")
	alog.Error("Error info...")
	alog.ErrorC("Error console info...")
	alog.Fatal("Fatal info...")
	alog.FatalC("Fatal console info...")
	time.Sleep(2 * time.Second)
}
```

### config.yaml:

``` yaml
{
	console: {
		level: 1,
		item: {
			tmpl: "[{{.ID}} {{.Time}} {{.Level}} {{.Tag}}] {{.Message}}",
			time: "{{.Hour}}:{{.Minute}}"
		}
	},
	global: {
	  	print: 1,rule: 1,showfile: 1,caller: 5,interval: 1,target: "file_global",
	  	buffer: {engine: 1}
	},
	store: {
		file: {
			file_global: {
		      filepath: "logs",
		      filename: "{{.Year}}{{.Month}}{{.Day}}.log",
		      filesize: 2048,
		      item: {
        		tmpl: '{{.ID}} {{.Time}} {{.Level}} {{.Tag}} "{{.ShortName}} {{.FileLine}} {{.FileFuncName}}" {{.Message}}',
		        time: "{{.Year}}-{{.Month}}-{{.Day}} {{.Hour}}:{{.Minute}}:{{.Second}}.{{.MilliSecond}}"
		      }
			}
		}
	}
}
```

### Console Output:

```
[1 15:26 Debug Sample] Debug info...
[0 15:26 Debug Sample] Debug console info...
[2 15:26 Info Sample] Info info...
[0 15:26 Info Sample] Info console info...
[3 15:26 Warn Sample] Warn info...
[0 15:26 Warn Sample] Warn console info...
[4 15:26 Error Sample] Error info...
[0 15:26 Error Sample] Error console info...
[5 15:26 Fatal Sample] Fatal info...
[0 15:26 Fatal Sample] Fatal console info...
```

### File Output:

```
1 2015-12-08 15:26:18.253 Debug Sample "main.go 12 main.main" Debug info...
2 2015-12-08 15:26:18.254 Info Sample "main.go 14 main.main" Info info...
3 2015-12-08 15:26:18.254 Warn Sample "main.go 16 main.main" Warn info...
4 2015-12-08 15:26:18.254 Error Sample "main.go 18 main.main" Error info...
5 2015-12-08 15:26:18.254 Fatal Sample "main.go 20 main.main" Fatal info...
```

## License

	Copyright 2015.All rights reserved.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.