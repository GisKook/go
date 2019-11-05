package file

import (
	"log"
	"testing"
)

func TestTempDir(t *testing.T) {
	dir, err := TempDir("", "temp")
	log.Println(dir)
	log.Println(err)

	log.Println(FileExist("not_exist"))
	log.Println(FileExist(dir))
	log.Println(RmDir(dir))
	log.Println(MD5("./file.go"))
}
