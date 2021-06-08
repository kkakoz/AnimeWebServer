package service

import (
	"github.com/google/wire"
	"github.com/labstack/echo"
	"log"
	"red-bean-anime-server/internal/app/file/domain"
	"red-bean-anime-server/internal/app/file/pkg/consts"
	"red-bean-anime-server/internal/app/file/usecase"
	"red-bean-anime-server/pkg/echox"
	"red-bean-anime-server/pkg/filex"
	"red-bean-anime-server/pkg/gerrors"
)

type FileService struct {
	fileUsecase domain.IFileUsecase
}

func NewFileService(fileUsecase domain.IFileUsecase) *FileService {
	return &FileService{fileUsecase: fileUsecase}
}

func (f *FileService) UploadVideo(ctx echo.Context) error  {
	header, err := ctx.FormFile("video")
	if err != nil {
		return err
	}
	file, err := header.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	fileType := filex.ParseFileType(header.Filename)
	fileName, err := f.fileUsecase.UploadFile(file, consts.VideoFilePath,  fileType)
	if err != nil {
		return err
	}
	return echox.ToRes(ctx, "/file/video/" + fileName);
}

func (f *FileService) UploadImage(ctx echo.Context) error {
	header, err := ctx.FormFile("image")
	if err != nil {
		return err
	}
	file, err := header.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	fileType := filex.ParseFileType(header.Filename)
	isImage := filex.IsImageFile(fileType)
	if !isImage {
		return gerrors.NewBusErr("无法识别的图片类型")
	}
	fileName, err := f.fileUsecase.UploadFile(file, consts.ImageFilePath, fileType)
	if err != nil {
		return err
	}
	return echox.ToRes(ctx, "/file/image/" + fileName);
}

func (f *FileService) GetVideo(ctx echo.Context) error {
	filename := ctx.Param("filename")
	if filename == "" {
		return gerrors.NewBusErr("请输入视频名称")
	}
	log.Println("filename = ", consts.VideoFilePath + filename)
	return ctx.File(consts.VideoFilePath + filename)
}

func (f *FileService) GetImage(ctx echo.Context) error {
	filename := ctx.Param("filename")
	if filename == "" {
		return gerrors.NewBusErr("请输入图片名称")
	}
	return ctx.File(consts.ImageFilePath + filename)
}

var ProviderSet = wire.NewSet(NewFileService, usecase.NewFileUsecase)