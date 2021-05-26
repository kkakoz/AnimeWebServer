package main

import (
	"flag"
	"github.com/labstack/echo"
	"log"
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
	err := engine.Start(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
}