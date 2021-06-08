package echox

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func PostInt(ctx echo.Context, name string) (int, error) {
	value := ctx.FormValue(name)
	return strconv.Atoi(value)
}

type res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ToRes(ctx echo.Context, data interface{}) error {
	res := &res{
		Code: 200,
		Data: data,
	}
	err := ctx.JSON(http.StatusOK, res)
	return err
}

func ToErr(ctx echo.Context, data interface{}) error {
	res := &res{
		Code: 500,
		Data: data,
	}
	err := ctx.JSON(http.StatusOK, res)
	return err
}