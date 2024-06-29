package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Media struct {
	Name string `json:"Name" validate:"required"`
	Link string `json:"Link" validate:"required"`
}

var movies []Media

var c = colly.NewCollector()

func InitScraper() {
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.AllowURLRevisit = true
}

func ReloadMovies() {

	c.OnHTML(".directory-entry", func(element *colly.HTMLElement) {
		if element.Text == "Parent Directory" {
			return
		}
		movies = append(movies, Media{element.Text, "https://vadapav.mov" + element.Attr("href")})
	})

	c.Visit("https://vadapav.mov/f36be06f-8edd-4173-99df-77bc4c7c2626/")
	c.OnHTMLDetach(".directory-entry")

	fmt.Println("movies: ", len(movies))
}

func GetMovieList() []Media {
	return movies
}

func GetMovies(name string) []Media {
	var _movies []Media
	c.OnHTML(".directory-entry", func(element *colly.HTMLElement) {
		if element.Text == "Parent Directory" {
			return
		}

		log.Println(element.Text)

		if strings.Contains(strings.ReplaceAll(strings.ToLower(element.Text), ",", ""), strings.ReplaceAll(strings.ToLower(name), ",", ",")) {
			_movies = append(_movies, Media{element.Text, "https://vadapav.mov" + element.Attr("href")})
		}

	})
	c.Visit("https://vadapav.mov/f36be06f-8edd-4173-99df-77bc4c7c2626/")
	c.OnHTMLDetach(".directory-entry")

	return _movies
}

func GetSources(movie Media) []Media {
	var sources []Media
	c.OnHTML(".file-entry", func(element *colly.HTMLElement) {
		if element.Text == "Parent Directory" {
			return
		}

		sources = append(sources, Media{movie.Name, "https://vadapav.mov" + element.Attr("href")})
	})
	c.Visit(movie.Link)
	c.OnHTMLDetach(".directory-entry")
	return sources
}

func SaveMovie(source Media) string {
	c.OnResponse(func(r *colly.Response) {
		log.Println("Downloading...")
		r.Save("assets/movies/" + source.Name)
		log.Println("Download Finfished!")
	})
	c.Visit(source.Link)

	return "assets/movies/" + source.Name
}
