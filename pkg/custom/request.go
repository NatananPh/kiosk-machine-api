package custom

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetParamInt(ctx echo.Context, key string) (int, error) {
	param := ctx.Param(key)
	return strconv.Atoi(param)
}