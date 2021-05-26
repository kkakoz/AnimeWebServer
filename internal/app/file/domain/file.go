package domain

import (
	"io"
)


type IFileUsecase interface {
	UploadFile(file io.Reader, path, fileType string) (string, error)
}

