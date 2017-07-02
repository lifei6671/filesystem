package filesystem

import (
	"io"
	"time"
)

type Values map[string]interface{}

type Filer interface {
	Exist(path string) error
	Read(path string) ([]byte,error)
	ReadStream(path string) (io.Reader,error)
	ListContents(dir string,recursive bool) ([]FileInfo,error)
	GetMetadata(path string) (FileInfo,error)
	GetFileSize(path string)(int64,error)
	GetMimeType(path string)(string,error)
	GetTimestamp(path string)(time.Time,error)
	GetVisibility(path string)(error)
	Write(path string,contents []byte,config Values)(error)
	WriteStream(path string,resource io.Writer)(error)
	Rename(path,newPath string)(error)
	Copy(src,dst string)(error)
	Delete(path string)(error)
	DeleteDir(path string)(error)
	CreateDir(path string,config Values)(error)
	SetVisibility(path string,visibility string)(error)
	Put(path string,contents []byte,config Values)(error)
	PutStream(path string,resource io.Writer,config Values)(error)
	ReadAndDelete(path string)([]byte,error)
}


type FileInfo struct {
	Type string
	Path string
	Size string
	Timestamp time.Time
	Metadata map[string]string
}

func (c *Values) Get(key string) (interface{},error) {
	if v,ok := c[key]; ok {
		return v,nil
	}
	return nil,ErrValueNoExist
}