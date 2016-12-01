package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/antlinker/alog"
)

const (
	// 写入日志条数
	_LogNum = 1000000
	// 日志信息长度
	_DataLen = 512
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
	alog.GALog.SetLogTag("MEMORY")
	ticker := time.NewTicker(time.Second)
	go output(startTime, ticker)
	go func() {
		logInfo := logData()
		for i := 0; i < _LogNum; i++ {
			alog.Info(logInfo)
		}
	}()
	endTime := <-_GCHComplete
	useSecond := float64(endTime.Sub(startTime))/float64(time.Second) - 1
	fmt.Printf("\n===> 文件日志写入\n===> 总条数：%d,总耗时：%.2fs,每条日志长度：%d,每秒写入日志条数：%d\n",
		_LogNum, useSecond, _DataLen, int64(_LogNum)/int64(useSecond))
}

func output(startTime time.Time, ticker *time.Ticker) {
	for t := range ticker.C {
		totalNum := alog.GALog.GetWriteNum()
		currentSecond := float64(t.Sub(startTime)) / float64(time.Second)
		info := fmt.Sprintf("\r ===> 写入日志条数：%d,用时：%.2fs", totalNum, currentSecond)
		fmt.Print(info)
		if totalNum == int64(_LogNum) {
			ticker.Stop()
			_GCHComplete <- time.Now()
		}
	}
}
