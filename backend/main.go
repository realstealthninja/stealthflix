package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func test(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "<html></html>")
}

func movie(ctx echo.Context) error {
	movie_name := ctx.QueryParam("name")
	return ctx.File("assets/movies/" + movie_name + " ")
}

func main() {
	e := echo.New()

	e.GET("/test", test)

	e.GET("/api/movies/:name", movie)

	// serve angular front end
	e.Static("/", "static")

	e.Logger.Fatal(e.Start(":8080"))
}
