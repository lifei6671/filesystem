package filesystem

import (
	"os"
	"io/ioutil"
	"io"
	"path/filepath"
	"time"
)

type LocalFile struct {

}

func (c *LocalFile) Exist(path string) error  {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return err
	}
	return err
}

func (c *LocalFile) Read(path string) ([]byte,error) {
	return ioutil.ReadFile(path)
}

func (c *LocalFile) ReadStream(path string) (io.Reader,error) {
	f,err := os.Open(path)
	if err != nil {
		return nil,err
	}
	return f,nil
}

func (c *LocalFile) ListContents(dir string,recursive bool) ([]os.FileInfo,error) {
	f,err := os.Lstat(dir)

	if err != nil {
		return nil,err
	}
	if !f.IsDir() {
		return nil,ErrPathNotDirectory
	}
	files := make([]os.FileInfo, 0)
	if recursive == false {
		dirs, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}


		for _, fi := range dirs {
			files = append(files, fi)
		}
	}else{
		err = filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, fi)
			return nil
		})

		if err != nil {
			return nil,err
		}
	}

	return files, nil
}

func (c *LocalFile) GetMetadata(path string) (os.FileInfo,error) {
	f,err := os.Lstat(path)
	if err != nil {
		return f,err
	}

	return f,nil
}

func (c *LocalFile) GetFileSize(path string)(int64,error) {
	f,err := os.Lstat(path)
	if err != nil {
		return 0,err
	}

	return f.Size(),nil
}

func (c *LocalFile) GetTimestamp(path string)(time.Time,error) {
	f,err := os.Lstat(path)
	if err != nil {
		return time.Time{},err
	}

	return f.ModTime(),nil
}

func (c *LocalFile) GetVisibility(path string)(os.FileMode) {
	f,err := os.Lstat(path)
	if err != nil {
		return 0
	}
	return f.Mode()
}

func (c *LocalFile) Write(path string,contents []byte,config Values) (error) {

	mode := os.ModePerm
	mode1,err := config.Get("FileMode")

	if err == nil {
		if m,ok := mode1.(os.FileMode); ok {
			mode = m
		}
	}

	f,err := os.OpenFile(path,os.O_CREATE|os.O_APPEND,mode)
	if err != nil {
		return err
	}
	defer f.Close()

	_,err = f.Write(contents)
	if err != nil {
		return err
	}

	return nil
}

func (c *LocalFile) WriteStream(path string,resource io.Reader)(error) {
	f,err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte,1024)
	for{
		n,err := resource.Read(buf)
		if err != nil && err != io.EOF{
			return err
		}
		if 0 == n { break }

		f.Write(buf[:n])
	}
	return nil
}

func (c *LocalFile) Rename(oldpath,newpath string)(error) {
	return os.Rename(oldpath,newpath)
}

func (c *LocalFile) Copy(src,dst string)(error) {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()
	dstf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dstf.Close()
	_,err = io.Copy(dstf, srcf)

	return err
}

func (c *LocalFile) Delete(path string)(error) {
	return os.Remove(path)
}

func (c *LocalFile) DeleteDir(path string)(error) {
	return  os.RemoveAll(path)
}

func (c *LocalFile) CreateDir(path string,config Values)(error) {
	mode := os.ModePerm
	mode1,err := config.Get("FileMode")

	if err == nil {
		if m,ok := mode1.(os.FileMode); ok {
			mode = m
		}
	}

	return os.MkdirAll(path,os.FileMode(mode))
}

func (c *LocalFile) SetVisibility(path string,mode os.FileMode)(error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return f.Chmod(os.FileMode(mode))
}

func (c *LocalFile) Put(path string,contents []byte,config Values)(error) {
	if c.Exist(path) == nil{
		os.Remove(path)
	}
	mode := os.ModePerm
	mode1,err := config.Get("FileMode")

	if err == nil {
		if m,ok := mode1.(os.FileMode); ok {
			mode = m
		}
	}

	f,err := os.Open(path)

	if err != nil {
		return err
	}
	defer f.Close()
	f.Chmod(os.FileMode(mode))
	_,err = f.Write(contents)

	return err
}

func (c *LocalFile) PutStream(path string,resource io.Reader,config Values)(error) {
	if c.Exist(path) == nil {
		os.Remove(path)
	}
	err := c.WriteStream(path,resource)

	if err != nil {
		return err
	}
	mode := os.ModePerm
	mode1,err := config.Get("FileMode")

	if err == nil {
		if m,ok := mode1.(os.FileMode); ok {
			mode = m
		}
	}else{
		return nil
	}
	f,err := os.Open(path)

	if err != nil {
		return err
	}
	defer f.Close()
	return f.Chmod(os.FileMode(mode))
}

func (c *LocalFile) ReadAndDelete(path string)([]byte,error) {
	defer os.Remove(path)
	return ioutil.ReadFile(path)
}









