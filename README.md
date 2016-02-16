# ALog

> 基于Golang的异步日志包

## Installation and usage

``` bash
$ go get github.com/antlinker/alog
```

## API documentation

* [https://godoc.org/github.com/antlinker/alog](https://godoc.org/github.com/antlinker/alog)

## Configuration file

``` yaml
{
	# 控制台输出配置
	console: {
		# 输出日志级别（1表示Debug,2表示Info,3表示Warn,4表示Error,5表示Fatal）
  		# 1表示输出Debug及以上所有级别的日志
		level: 1,
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
			tmpl: "[{{.Time}}｜{{.Level}}｜{{.Tag}}]{{.ShortName}}:{{.FileLine}}:{{.Message}}",
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
		# IsEnabled 是否启用日志
		# 参数说明：
		# 1表示启用
		# 2表示不启用
		# 默认值为1
		enabled: 1,
		# 是否打印到控制台
	  	# 1表示打印
	  	# 2表示不打印
	  	# 默认为不打印
	  	print: 1,
	  	# 日志输出规则
	  	# 参数说明：
	  	# 0表示所有配置输出
	  	# 1表示指定Global配置输出
	  	# 2表示指定Tag配置输出
	  	# 3表示指定Level配置输出
	  	# 默认为所有配置输出(0)
	  	rule: 0,
	  	# 日志级别
			level: 1,
	  	# 输出文件信息
	  	# 1表示输出
	  	# 2表示不输出
	  	# 默认为输出(1)
	  	showfile: 1,
	  	# 文件信息调用层级
	  	# 默认为6(当前调用)
	  	caller: 6,
	  	# 读取缓冲区时间间隔（以秒为单位）
	  	interval: 1,
	  	# 目标存储
	  	# 指向store中定义的存储配置
	  	# target: "file_global",
	  	# 缓冲区存储
	  	# buffer: {
	  		# 存储引擎
		    # 1表示内存存储
		    # 2表示redis存储
		    # engine: 1,
		    # 指向store中定义的存储配置
		    # target: "redis_buffer"
	  	# }
	},
	# 标签配置
	tags: [{
		# 标签名
  		names: ["ALOG","TEST"],
  		# 配置
  		config: {
  			# 是否打印到控制台
		    # 1表示打印
	  		# 2表示不打印
		    print: 2,
		    # 日志级别
		    level: 1,
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
		    # 1表示打印
	  		# 2表示不打印
		    print: 2,
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
		},
		# elasticsearch存储
		elastic: {
			elastic_global: {
				# 指定ElasticSearch的请求节点
				# 默认值为http://127.0.0.1:9200
				url: "http://127.0.0.1:9200",
				# 索引模板
				# 模板字段说明：
				# Year 年份
				# Month 月份
				# Day 天数
				# Level 日志级别
				# Tag 标签
				# 默认值为{{.Year}}.{{.Month}}.{{.Day}}
				index: "{{.Year}}.{{.Month}}.{{.Day}}",
				# 文档类型模板
				# 模板字段说明：
				# Year 年份
				# Month 月份
				# Day 天数
				# Level 日志级别
				# Tag 标签
				# 默认值为ALogs
				type: "ALogs"
			}
		},
		# MongoDB存储
		mongo: {
			mongo_global: {
				# 指定MongoDB的链接地址
				# 默认值为mongodb://127.0.0.1:27017
				url: "mongodb://127.0.0.1:27017",
				# 存储数据库名称模板
				# 模板字段说明：
				# Year 年份
				# Month 月份
				# Day 天数
				# Level 日志级别
				# Tag 标签
				# 默认值为alog
				db: "alog",
				# 存储集合名称模板
				# 模板字段说明：
				# Year 年份
				# Month 月份
				# Day 天数
				# Level 日志级别
				# Tag 标签
				# 默认值为{{.Year}}{{.Month}}{{.Day}}
				collection: "{{.Year}}{{.Month}}{{.Day}}"
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

	"github.com/antlinker/alog"
)

func main() {
	alog.RegisterAlog()
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

### Console Output:

```
[13:36:58｜Debug｜Sample]main.go:13:Debug console info...
[13:36:58｜Info｜Sample]main.go:15:Info console info...
[13:36:58｜Warn｜Sample]main.go:17:Warn console info...
[13:36:58｜Error｜Sample]main.go:19:Error console info...
[13:36:58｜Fatal｜Sample]main.go:21:Fatal console info...
```

### File Output:

```
2016-01-18 13:38:09.721 Debug Sample "main.go main.main 12" Debug info...
2016-01-18 13:38:09.721 Info Sample "main.go main.main 14" Info info...
2016-01-18 13:38:09.721 Warn Sample "main.go main.main 16" Warn info...
2016-01-18 13:38:09.721 Error Sample "main.go main.main 18" Error info...
2016-01-18 13:38:09.721 Fatal Sample "main.go main.main 20" Fatal info...
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