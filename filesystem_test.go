package lua

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestSetFileSystem(t *testing.T) {
	var err error
	var f *os.File
	var dir = "."
	var filename = "test"

	// make temp file
	if f, err = ioutil.TempFile(dir, filename); err != nil{
		t.Fatalf("can not make temp file %v: %v", filename, err)
	}
	filename = f.Name()
	defer func(){
		f.Close()
		os.Remove(dir + string(filepath.Separator) + f.Name())
	}()

	// write lua content
	if _, err = f.WriteString(`print("hello")`); err != nil {
		t.Fatalf("can not write temp file %v: %v", filename, err)
	}

	// set the file system
	SetFileSystem(http.Dir(dir))
	L := NewState()
	if err := L.DoFile(filename); err != nil {
		t.Fatalf("can not do file %v: %v", filename, err)
	}

}