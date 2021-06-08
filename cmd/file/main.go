package main

import (
	"flag"
	"github.com/labstack/echo"
	"log"
	"red-bean-anime-server/pkg/echox"
	"strconv"
)


var port = flag.Int("port", 8080, "web port")

func main() {
	flag.Parse()
	engine := echo.New()
	fileService := NewService()
	engine.POST("/file/image", fileService.UploadImage)
	engine.POST("/file/video", fileService.UploadVideo)
	engine.GET("/file/video/:filename", fileService.GetVideo)
	engine.GET("/file/image/:filename", fileService.GetImage)
	engine.HTTPErrorHandler = echox.HandlerErr()
	err := engine.Start(":" + strconv.Itoa(*port))

	//err := engine.StartTLS(":" + strconv.Itoa(*port),
	//	"./configs/server.crt",
	//	"./configs/server.key",)
	//
	if err != nil {
		log.Fatal(err)
	}
}