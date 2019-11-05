package file

import (
	"crypto/md5"
	gkbytes "github.com/giskook/go/bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TempDir(dir, prefix string) (string, error) {
	return ioutil.TempDir(dir, prefix)
}

func FileExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func RmDir(dir string) error {
	return os.RemoveAll(dir)
}

func GetFileSize(file_path string) (int64, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	file_stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return file_stat.Size(), nil
}

func GetDir(file string) string {
	return filepath.Dir(file)
}

func MD5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return gkbytes.GetBcdString(h.Sum(nil)), nil
}
