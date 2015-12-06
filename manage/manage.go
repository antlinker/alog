package manage

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"

	"gopkg.in/alog.v1/buffer"
	"gopkg.in/alog.v1/log"
	"gopkg.in/alog.v1/store"
)

var (
	// 日志项标识
	_GLOGID uint64
)

// NewLogManage 创建新的LogManage实例
func NewLogManage(config *log.LogConfig) log.LogManage {
	cfg := *config
	manage := &_LogManage{
		Config: config,
	}
	manage.Template = map[log.TmplKey]*template.Template{
		log.TmplConsole:     template.Must(template.New("").Parse(cfg.Console.Item.Tmpl)),
		log.TmplConsoleTime: template.Must(template.New("").Parse(cfg.Console.Item.TimeTmpl)),
	}
	switch cfg.Global.Buffer.Engine {
	case log.REDIS_BUFFER:
		var redisConfig log.RedisConfig
		redisStore := cfg.Store.Redis
		if redisStore != nil {
			if v, ok := redisStore[cfg.Global.Buffer.TargetStore]; ok {
				redisConfig = v
			}
		}
		manage.Buffer = buffer.NewRedisBuffer(redisConfig)
	default:
		manage.Buffer = buffer.NewMemoryBuffer()
	}

	manageStore := make(map[string]log.LogStore)
	if fileStore := cfg.Store.File; fileStore != nil {
		for fk, fv := range fileStore {
			manageStore[fk] = store.NewFileStore(fv)
		}
	}
	manage.Store = manageStore
	go manage.execStore()

	return manage
}

// _LogManage
type _LogManage struct {
	locker   sync.RWMutex
	total    int64
	Config   *log.LogConfig
	Template map[log.TmplKey]*template.Template
	Buffer   log.LogBuffer
	Store    map[string]log.LogStore
}

func (lm *_LogManage) Write(level log.LogLevel, tag log.LogTag, v ...interface{}) {
	msg := fmt.Sprint(v...)
	lm.writeMsg(level, tag, msg)
}

func (lm *_LogManage) Writef(level log.LogLevel, tag log.LogTag, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	lm.writeMsg(level, tag, msg)
}

func (lm *_LogManage) Console(level log.LogLevel, tag log.LogTag, v ...interface{}) {
	msg := fmt.Sprint(v...)
	lm.writeConsole(level, tag, msg)
}

func (lm *_LogManage) Consolef(level log.LogLevel, tag log.LogTag, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	lm.writeConsole(level, tag, msg)
}

func (lm *_LogManage) writeConsole(level log.LogLevel, tag log.LogTag, msg string) {
	item := lm.logItem(level, tag, msg)
	lm.console(&item)
}

func (lm *_LogManage) TotalNum() int64 {
	lm.locker.RLock()
	defer lm.locker.RUnlock()
	return lm.total
}

func (lm *_LogManage) writeMsg(level log.LogLevel, tag log.LogTag, msg string) {
	item := lm.logItem(level, tag, msg)
	item.ID = atomic.AddUint64(&_GLOGID, 1)
	if lm.Config.Global.ShowFile > 0 {
		item.File = lm.file()
	}
	if lm.isPrint(&item) {
		lm.console(&item)
	}
	lm.Buffer.Push(item)
}

func (lm *_LogManage) logItem(level log.LogLevel, tag log.LogTag, msg string) log.LogItem {
	item := log.LogItem{
		Level:   level,
		Time:    time.Now(),
		Tag:     tag,
		Message: msg,
	}
	return item
}

func (lm *_LogManage) file() log.LogFile {
	var logFile log.LogFile
	pc, file, line, ok := runtime.Caller(lm.Config.Global.FileCaller)
	if !ok {
		logFile.Name = "???"
		logFile.FuncName = "???"
		return logFile
	}
	logFile.Name = file
	logFile.Line = line
	logFile.FuncName = runtime.FuncForPC(pc).Name()
	return logFile
}

func (lm *_LogManage) isPrint(item *log.LogItem) bool {
	var isPrint bool
	switch lm.Config.Global.Rule {
	case log.AlwaysRule:
		isPrint = lm.isEitherTrue(item, func(lm *_LogManage, item *log.LogItem) bool {
			return lm.Config.Global.IsPrint > 0
		}, func(lm *_LogManage, item *log.LogItem) bool {
			var v bool
			lm.tags(item, func(t log.TagConfig) bool {
				if t.Config.IsPrint > 0 {
					v = true
					return true
				}
				return false
			})
			return v
		}, func(lm *_LogManage, item *log.LogItem) bool {
			var v bool
			lm.levels(item, func(l log.LevelConfig) bool {
				if l.Config.IsPrint > 0 {
					v = true
					return true
				}
				return false
			})
			return v
		})
	case log.GlobalRule:
		isPrint = lm.Config.Global.IsPrint > 0
	case log.TagRule:
		lm.tags(item, func(t log.TagConfig) bool {
			if t.Config.IsPrint > 0 {
				isPrint = true
				return true
			}
			return false
		})
	case log.LevelRule:
		lm.levels(item, func(l log.LevelConfig) bool {
			if l.Config.IsPrint > 0 {
				isPrint = true
				return true
			}
			return false
		})
	}
	return isPrint
}

func (lm *_LogManage) isEitherTrue(item *log.LogItem, fns ...func(*_LogManage, *log.LogItem) bool) bool {
	for _, fn := range fns {
		if fn(lm, item) {
			return true
		}
	}
	return false
}

func (lm *_LogManage) tags(item *log.LogItem, fn func(log.TagConfig) bool) {
	for _, tag := range lm.Config.Tags {
		for _, tagName := range tag.Names {
			if tagName == (*item).Tag {
				if fn(tag) {
					return
				}
				break
			}
		}
	}
}

func (lm *_LogManage) levels(item *log.LogItem, fn func(log.LevelConfig) bool) {
	for _, lev := range lm.Config.Levels {
		for _, levVal := range lev.Values {
			if (*item).Level == levVal {
				if fn(lev) {
					return
				}
				break
			}
		}
	}
}

func (lm *_LogManage) console(item *log.LogItem) {
	if lm.Config.Console.Level <= (*item).Level {
		lm.stdout(lm.Template[log.TmplConsole], lm.Template[log.TmplConsoleTime], item)
	}
}

func (lm *_LogManage) stdout(tmpl, timetmpl interface{}, item *log.LogItem) {
	buf := log.ParseLogItemToBuffer(tmpl, timetmpl, item)
	fmt.Fprintln(os.Stdout, buf.String())
}

func (lm *_LogManage) execStore() {
	interval := time.Duration(lm.Config.Global.Interval) * time.Second
	time.Sleep(interval)
	lm.afterStore(interval)
}

func (lm *_LogManage) afterStore(interval time.Duration) {
	lm.store()
	time.AfterFunc(interval, func() {
		lm.execStore()
	})
}

func (lm *_LogManage) store() {
	defer func() {
		if err := recover(); err != nil {
			item := lm.logItem(log.FATAL, "SYSTEM", fmt.Sprint(err))
			lm.stdout(log.DefaultSystemTmpl, log.DefaultConsoleTimeTmpl, &item)
		}
	}()
	for {
		item, err := lm.Buffer.Pop()
		if err != nil {
			panic(err)
		}
		if item == nil {
			break
		}
		var errs []error
		targets := lm.storeTargets(item)
		l := len(targets)
		mux := new(sync.Mutex)
		wg := new(sync.WaitGroup)
		wg.Add(l)
		for i, l := 0, len(targets); i < l; i++ {
			go lm.writeStore(wg, mux, &errs, targets[i], item)
		}
		wg.Wait()
		if len(errs) == l {
			panic("Write store error.")
		}
		lm.locker.Lock()
		lm.total++
		lm.locker.Unlock()
	}
}

func (lm *_LogManage) target(target map[string]string, ts string) {
	tsa := strings.Split(ts, ",")
	for i := 0; i < len(tsa); i++ {
		t := tsa[i]
		if t == "" {
			continue
		}
		if _, ok := target[t]; !ok {
			target[t] = t
		}
	}
}

func (lm *_LogManage) storeTargets(item *log.LogItem) (targets []string) {
	target := make(map[string]string)
	rule := lm.Config.Global.Rule

	if rule == log.AlwaysRule || rule == log.GlobalRule {
		lm.target(target, lm.Config.Global.TargetStore)
	}
	if rule == log.AlwaysRule || rule == log.TagRule {
		lm.tags(item, func(tag log.TagConfig) bool {
			if tag.Level <= item.Level {
				lm.target(target, tag.Config.TargetStore)
			}
			return false
		})
	}
	if rule == log.AlwaysRule || rule == log.LevelRule {
		lm.levels(item, func(lev log.LevelConfig) bool {
			lm.target(target, lev.Config.TargetStore)
			return false
		})
	}

	for k := range target {
		targets = append(targets, k)
	}

	return
}

func (lm *_LogManage) writeStore(wg *sync.WaitGroup, mux *sync.Mutex, errs *[]error, target string, item *log.LogItem) {
	defer wg.Done()
	store, ok := lm.Store[target]
	if !ok {
		mux.Lock()
		*errs = append(*errs, fmt.Errorf("Unknown store."))
		mux.Unlock()
		return
	}
	err := store.Store(item)
	if err != nil {
		mux.Lock()
		*errs = append(*errs, err)
		mux.Unlock()
	}
}
