package helper

import (
	"net/http"

	"github.com/labstack/echo"
)

// JSONWapper format response structure
func JSONWapper(f func(c echo.Context) (int, string, interface{})) func(echo.Context) error {
	return func(c echo.Context) error {
		code, msg, data := f(c)

		msg = "success"
		if code != http.StatusOK {
			msg = "error"
		}

		var resp = struct {
			Code int         `json:"code"`
			Msg  string      `json:"msg"`
			Data interface{} `json:"data"`
		}{
			Code: code,
			Msg:  msg,
			Data: data,
		}

		return c.JSON(http.StatusOK, resp)
	}
}

// Health health check for echo
func Health(e *echo.Echo) {
	e.GET("/health", JSONWapper(func(c echo.Context) (int, string, interface{}) {
		return http.StatusOK, "", nil
	}))
}
