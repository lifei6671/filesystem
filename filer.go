package filesystem

import (
	"io"
	"time"
	"os"
)

type Values map[string]interface{}

type Filer interface {
	Exist(path string) error
	Read(path string) ([]byte,error)
	ReadStream(path string) (io.Reader,error)
	ListContents(dir string,recursive bool) ([]os.FileInfo,error)
	GetMetadata(path string) (os.FileInfo,error)
	GetFileSize(path string)(int64,error)
	GetTimestamp(path string)(time.Time,error)
	GetFileMode(path string)(os.FileMode)
	Write(path string,contents []byte,config Values)(error)
	WriteStream(path string,resource io.Reader)(error)
	Rename(path,newPath string)(error)
	Copy(src,dst string)(error)
	Delete(path string)(error)
	DeleteDir(path string)(error)
	CreateDir(path string,config Values)(error)
	SetVisibility(path string,mode os.FileMode)(error)
	Put(path string,contents []byte,config Values)(error)
	PutStream(path string,resource io.Reader,config Values)(error)
	ReadAndDelete(path string)([]byte,error)
}


func (c Values) Get(key string) (interface{},error) {
	if v,ok := c[key]; ok {
		return v,nil
	}
	return nil,ErrValueNoExist
}

func (c Values) Set(key string,value interface{}) {
	c[key] = value
}