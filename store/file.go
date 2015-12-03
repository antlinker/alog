package store

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/alog.v1/log"
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
	cfg := &_FileConfig{
		Size:     size * 1024,
		Path:     fpath,
		NameTmpl: template.Must(template.New("").Parse(filename)),
		TimeTmpl: template.Must(template.New("").Parse(timeTmpl)),
		MsgTmpl:  template.Must(template.New("").Parse(msgTmpl)),
	}
	return &FileStore{config: cfg}
}

// FileStore 提供文件日志存储
type FileStore struct {
	config *_FileConfig
}

func (fs *FileStore) formatName(name string, num int) string {
	if num > 0 {
		return fmt.Sprintf("%s-%d", name, num)
	} else {
		return fmt.Sprintf("%s", name)
	}
}

func (fs *FileStore) fileName(item *log.LogItem) string {
	var (
		number     int
		filterFile []os.FileInfo
	)
	fName := log.ParseFileName(fs.config.NameTmpl, item)
	ext := filepath.Ext(fName)
	fName = strings.TrimSuffix(fName, ext)
	root := fs.config.Path
	prefix := root + "/" + fName
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasPrefix(path, prefix) {
			return nil
		}
		filterFile = append(filterFile, info)
		return nil
	})
	if l := len(filterFile); l > 0 {
		number = l - 1
		for _, file := range filterFile {
			name := fs.formatName(fName, number)
			ffName := file.Name()
			ffName = strings.TrimSuffix(ffName, filepath.Ext(ffName))
			if ffName == name {
				if file.Size() >= fs.config.Size {
					fName = fs.formatName(fName, number+1)
				} else {
					fName = name
				}
				break
			}
		}
	}
	if fName == "" {
		return ""
	}
	if ext == "" {
		ext = ".log"
	}
	return fName + ext
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

func (fs *FileStore) Store(item *log.LogItem) error {
	fs.createFolder()
	fileName := fs.fileName(item)
	if fileName == "" {
		return fmt.Errorf("The file name is invalid.")
	}
	fileName = filepath.Join(fs.config.Path, fileName)
	logInfo := log.ParseLogItem(fs.config.MsgTmpl, fs.config.TimeTmpl, item)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(logInfo)
	return err
}
