package filex

import (
	"io"
	"os"
)

func FileUpload(file io.Reader, path, name string) (string, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	dst, err := os.Create(path + name)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(dst, file)
	return path + name, err
}


