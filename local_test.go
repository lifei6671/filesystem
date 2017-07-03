package filesystem

import (
	"testing"
	"path/filepath"
	"os"
)

func TestLocalFile_ListContents(t *testing.T) {
	p := filepath.Dir(os.Args[0])

	t.Log(filepath.Base(os.Args[0]))

	storage := LocalFile{}

	files,err := storage.ListContents(p,false)

	if err != nil {
		t.Fatal(err)
	}
	for _,f := range files {
		t.Log(f.Name())
	}

}

func TestLocalFile_PutStream(t *testing.T) {
	p := filepath.Dir(os.Args[0])
	t.Log(p)
	storage := LocalFile{}
	v := Values{}
	v.Set("FileMode",os.ModePerm)

	filePath := p + string(os.PathSeparator) + "a.txt"

	err := storage.Write(filePath,[]byte("测试"),v)

	if err != nil {
		t.Fatal(err)
	}
	f,err := storage.ReadStream(filePath)
	if err != nil {
		t.Fatal(err)
	}

	err = storage.PutStream(filePath,f,Values{})

	if err != nil {
		t.Fatal(err)
	}

}