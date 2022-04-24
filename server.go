package main

import (
  "github.com/labstack/echo/v4"
  "net/http"
)

func main()  {
  e := echo.New()

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  e.GET("/show", show)
  e.GET("/user", user)

  e.Logger.Fatal(e.Start(":1323"))
}

// curl -i "http://localhost:1323/show?team=x-men&member=wolverine"
func show(c echo.Context) error {
  team := c.QueryParam("team")
  member := c.QueryParam("member")
  return c.String(http.StatusOK, "team: " + team + ", member: " + member)
}

type User struct {
  Name string `json:"name" form:"name" query:"name"`
  Email string `json:"email" form:"email" query:"email"`
}

// curl -i "http://localhost:1323/user?user=logx&email=logx@foo.com"
func user(c echo.Context) error {
  u := new(User)

  if err := c.Bind(u); err != nil {
    return  err
  }
  return c.JSON(http.StatusCreated, u)
}
