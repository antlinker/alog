package store

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/antlinker/alog/log"
)

type _FileConfig struct {
	Size     int64
	Path     string
	NameTmpl *template.Template
	TimeTmpl *template.Template
	MsgTmpl  *template.Template
}

// NewFileStore 创建新的FileStore实例
func NewFileStore(config log.FileConfig) log.LogStore {
	var (
		size     = config.FileSize
		fpath    = config.FilePath
		filename = config.FileNameTmpl
		timeTmpl = config.Item.TimeTmpl
		msgTmpl  = config.Item.Tmpl
	)
	if size == 0 {
		size = log.DefaultFileSize
	}
	if fpath == "" {
		fpath = log.DefaultFilePath
	}
	if !filepath.IsAbs(fpath) {
		fpath, _ = filepath.Abs(fpath)
	}
	if filename == "" {
		filename = log.DefaultFileNameTmpl
	}
	if timeTmpl == "" {
		timeTmpl = log.DefaultTimeTmpl
	}
	if msgTmpl == "" {
		msgTmpl = log.DefaultMsgTmpl
	}

	if l := len(fpath); l > 0 && fpath[l-1] == '/' {
		fpath = fpath[:l-1]
	}

	cfg := &_FileConfig{
		Size:     size * 1024,
		Path:     fpath,
		NameTmpl: template.Must(template.New("").Parse(filename)),
		TimeTmpl: template.Must(template.New("").Parse(timeTmpl)),
		MsgTmpl:  template.Must(template.New("").Parse(msgTmpl)),
	}
	fs := &FileStore{config: cfg}

	// 创建日志目录
	if err := fs.createFolder(); err != nil {
		panic("创建目录发生错误：" + err.Error())
	}

	return fs
}

// FileStore 提供文件日志存储
type FileStore struct {
	config   *_FileConfig
	fileName string
	file     *os.File
}

func (fs *FileStore) formatName(name string, num int) string {
	if num > 0 {
		ns := strconv.Itoa(num)
		return fmt.Sprintf("%s-%s%d", name, strings.Repeat("0", 4-len(ns)), num)
	} else {
		return fmt.Sprintf("%s", name)
	}
}

func (fs *FileStore) createFolder() error {
	folder := fs.config.Path
	_, err := os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (fs *FileStore) changeName(name, ext string) string {
	var (
		number int
	)

	prefix := fs.config.Path + "/" + name
	err := filepath.Walk(fs.config.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasPrefix(path, prefix) {
			return nil
		}
		number++
		return nil
	})
	if err != nil {
		fmt.Println("FileStore Error:", err.Error())
		return ""
	}

	return fmt.Sprintf("%s-%d%s", name, number, ext)
}

func (fs *FileStore) getFile(item *log.LogItem) (file *os.File, err error) {
	fileName := log.ParseName(fs.config.NameTmpl, item)
	if fileName == "" {
		fileName = fmt.Sprintf("unknown.%s.log", item.Time.Format("20060102"))
	}

	ext := filepath.Ext(fileName)
	prefix := strings.TrimSuffix(fileName, ext)
	if strings.HasPrefix(fs.fileName, prefix) {
		if fs.file != nil {
			file = fs.file
			return
		} else {
			fs.fileName = fs.changeName(prefix, ext)
		}
	} else {
		if fs.file != nil {
			fs.file.Close()
			fs.file = nil
		}
		fs.fileName = fileName
	}

	if fs.file == nil {
		file, err = os.OpenFile(fmt.Sprintf("%s/%s", fs.config.Path, fs.fileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			return
		}
		fs.file = file
		return
	}

	file = fs.file
	return
}

func (fs *FileStore) Store(item *log.LogItem) (err error) {

	file, err := fs.getFile(item)
	if err != nil {
		return
	}

	logInfo := log.ParseLogItem(fs.config.MsgTmpl, fs.config.TimeTmpl, item)
	_, err = file.WriteString(logInfo)
	if err != nil {
		return
	}

	finfo, err := file.Stat()
	if err != nil {
		return
	} else if finfo.Size() >= fs.config.Size {
		fs.file.Close()
		fs.file = nil
	}

	return
}
