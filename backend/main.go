package main

import (
	"log"
	"net/http"
	"net/url"

	"stealthflix-backend/cmd"

	"github.com/labstack/echo/v4"
)

func list(ctx echo.Context) error {

	movies := cmd.GetMovieList()

	return ctx.JSON(http.StatusOK, movies)

}

func getMovie(ctx echo.Context) error {
	name, err := url.PathUnescape(ctx.QueryParam("name"))

	if err != nil {

		log.Fatal(err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	movies := cmd.GetMovies(name)

	return ctx.JSON(http.StatusOK, movies)
}

func sources(ctx echo.Context) error {
	link, err := url.PathUnescape(ctx.QueryParam("link"))
	name, err2 := url.PathUnescape(ctx.QueryParam("name"))

	if err != nil || err2 != nil {

		log.Fatal(err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	sources := cmd.GetSources(cmd.Media{Name: name, Link: link})

	return ctx.JSON(http.StatusOK, sources)

}

func main() {
	e := echo.New()

	e.GET("/api/movies/list", list)

	e.GET("/api/movies/sources", sources)

	e.GET("/api/movies/get", getMovie)

	// serve angular front end
	e.Static("/", "static")

	cmd.InitScraper()
	cmd.ReloadMovies()

	e.Logger.Fatal(e.Start(":8080"))
}
