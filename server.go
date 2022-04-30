package main

import (
	"github.com/labstack/echo/v4"
	"github.com/wulianglongrd/echo-example/router"
)

func main() {
	e := echo.New()

	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
