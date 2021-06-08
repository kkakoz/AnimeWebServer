package usecase

import (
	"io"
	"red-bean-anime-server/internal/app/file/domain"
	"red-bean-anime-server/internal/app/file/pkg/filex"
	"red-bean-anime-server/pkg/cryption"
	"red-bean-anime-server/pkg/gerrors"
)

type FileUsecase struct {
}

func NewFileUsecase() domain.IFileUsecase {
	return &FileUsecase{}
}

func (v *FileUsecase) UploadFile(file io.Reader, path,  fileType string) (string, error) {
	if fileType == "" {
		return "", gerrors.NewBusErr("请输入正确的文件类型")
	}
	newName := cryption.UUID()
	_, err := filex.FileUpload(file, path, newName+"."+fileType)
	if err != nil {
		return "", err
	}
	return newName+"."+fileType, nil
}

