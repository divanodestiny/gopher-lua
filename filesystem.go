package lua

import (
	"net/http"
	"os"
)

type LocalFileSystem struct{}

func (lfs *LocalFileSystem) Open(name string) (http.File, error){
	return os.Open(name)
}

var filesystem http.FileSystem = &LocalFileSystem{}

func SetFileSystem(extFileSystem http.FileSystem){
	filesystem = extFileSystem
}
