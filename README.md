# echo-helper
封装的一些 echo 框架常用的函数

## 安装
```
go get github.com/JinesD/echo-helper
```
## 使用
```
import (
    "github.com/labstack/echo"
    ehelper "github.com/JinesD/echo-helper"
)

func main () {
    e := echo.New()

    e.GET("/", ehelper.JSONWapper(func (c echo.Context) (int, string, interface{}) {
        return 200, "success", nil
    }))

    e.Start(":8080")
}
```