package main

import (
	"log"
	"net/http"
	"net/url"

	"stealthflix-backend/cmd"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MediaValidator struct {
	validator *validator.Validate
}

func (mv *MediaValidator) Validate(i interface{}) error {
	if err := mv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func test(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "<html></html>")
}

func movies(ctx echo.Context) error {

	movies := cmd.GetMovieList()

	return ctx.JSON(http.StatusOK, movies)

}

func getMovie(ctx echo.Context) error {
	movieName, err := url.PathUnescape(ctx.Param("name"))

	if err != nil {
		log.Fatal(err)
	}
	log.Println(movieName)
	movies := cmd.GetMovies(movieName)

	return ctx.JSON(http.StatusOK, movies)
}

func moviesPost(ctx echo.Context) error {
	movie := new(cmd.Media)
	var err error
	if err = ctx.Bind(movie); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if err = ctx.Validate(movie); err != nil {
		return err
	}

	sources := cmd.GetSources(*movie)

	return ctx.JSON(http.StatusOK, sources)

}

func main() {
	e := echo.New()
	e.Validator = &MediaValidator{validator: validator.New()}

	e.GET("/test", test)

	e.GET("/api/movies/", movies)
	e.POST("/api/movies/", moviesPost)

	e.GET("/api/movies/:name", getMovie)

	// serve angular front end
	e.Static("/", "static")

	cmd.InitScraper()
	cmd.ReloadMovies()

	e.Logger.Fatal(e.Start(":8080"))
}
