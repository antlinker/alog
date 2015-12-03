# ALog

> 基于Golang的分布式日志包

## 获取

``` bash
$ go get gopkg.in/alog.v1
```

## 日志测试

``` go
package main

import (
	"bytes"
	"fmt"
	"time"

	"gopkg.in/alog.v1"
)

const (
	// 写入日志条数
	_LogNum = 100000
	// 日志信息长度
	_DataLen = 512
	// 日志测试标签
	_LogTag  = "TEST"
)

var (
	_GCHComplete chan time.Time
)

func logData() string {
	buf := new(bytes.Buffer)
	for i := 0; i < _DataLen; i++ {
		buf.WriteByte('a')
	}
	return buf.String()
}

func main() {
	_GCHComplete = make(chan time.Time, 1)
	startTime := time.Now()
	alog.RegisterAlog("config.yaml")
	ticker := time.NewTicker(time.Second)
	go output(startTime, ticker)
	go func() {
		logInfo := logData()
		for i := 0; i < _LogNum; i++ {
			alog.Debug(_LogTag, logInfo)
		}
	}()
	endTime := <-_GCHComplete
	useSecond := float64(endTime.Sub(startTime))/float64(time.Second) - 1
	fmt.Printf("\n===> 文件日志写入,总条数：%d,总耗时：%.2fs,每条日志长度：%d,每秒写入日志条数：%d\n",
		_LogNum, useSecond, _DataLen, int64(_LogNum)/int64(useSecond))
}

func output(startTime time.Time, ticker *time.Ticker) {
	for t := range ticker.C {
		totalNum := alog.GLogManage.TotalNum()
		currentSecond := float64(t.Sub(startTime)) / float64(time.Second)
		info := fmt.Sprintf("\r ===> 写入日志条数：%d,用时：%.2fs", totalNum, currentSecond)
		fmt.Print(info)
		if totalNum == int64(_LogNum) {
			ticker.Stop()
			_GCHComplete <- time.Now()
		}
	}
}

```

## 输出结果

``` bash
# 内存写入统计
# ===> 文件日志写入,总条数：100000,总耗时：13.00s,每条日志长度：512,每秒写入日志条数：7692
# redis写入统计
# ===> 文件日志写入,总条数：100000,总耗时：72.01s,每条日志长度：512,每秒写入日志条数：1388
```

## 配置文件说明

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
	tags: {
		# 标签名
  		names: ["ALOG","TEST"],
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
	},
	# 级别配置
	levels: {
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
	},
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