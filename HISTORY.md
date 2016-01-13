# HISTORY

## v1.1.0

* 增加针对`ElasticSearch`的持久化存储
* 增加针对`MongoDB`的持久化存储
* 优化一些实现及错误处理

## v1.0.6

* 针对`Global`增加`Level`配置
* 针对`manage`中的一些实现方式进行调整

## v1.0.5

* 修复一些关于配置`bug`
* 增加`IsEnabled`配置是否启用日志

## v1.0.4

* 针对`ALog`中的日志级别函数进行稍大强度的调整
* 规范`Global`日志级别函数输出
* 增加`NewALog`函数，运行外部创建`ALog`实例

## v1.0.3

* 增加`ALog`结构体，提供统一的处理函数
* 加入`ShortName`(短文件名)日志项模板格式
* 调整`README.md`

## v1.0.2

* `target`可以指向多个store,以`,`分隔
* 增加针对`FileCaller`的配置
* 在`Tags`配置下，增加`level`配置
* 针对范例程序进行更细致的划分，单独抽出`config`