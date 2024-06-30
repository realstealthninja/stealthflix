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

var cDownload *colly.Collector

var cSourceFinder *colly.Collector

func InitScraper() {

	c.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.AllowURLRevisit = true

	cSourceFinder = c.Clone()

	cDownload = c.Clone()

	//                      GB *  MB  *  KB  * B = 30 GB
	cDownload.MaxBodySize = 30 * 1024 * 1024 * 1024
}

func ReloadMovies() {
	log.Println("COLLY: reloading movie list")
	c.OnHTML(".directory-entry", func(element *colly.HTMLElement) {
		if element.Text == "Parent Directory" {
			return
		}
		movies = append(movies, Media{element.Text, "https://vadapav.mov" + element.Attr("href")})
	})

	c.Visit("https://vadapav.mov/f36be06f-8edd-4173-99df-77bc4c7c2626/")

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
	cSourceFinder.OnHTML(".file-entry", func(element *colly.HTMLElement) {
		if element.Text == "Parent Directory" {
			return
		}

		exts := strings.Split(element.Text, ".")
		ext := exts[len(exts)-1]
		sources = append(sources, Media{movie.Name + ext, "https://vadapav.mov" + element.Attr("href")})
	})
	cSourceFinder.Visit(movie.Link)
	return sources
}

func SaveMovie(source Media) string {

	cDownload.OnResponse(func(r *colly.Response) {
		log.Println("Downloading...")
		err := r.Save("assets/movies/" + source.Name)
		if err != nil {
			log.Println(err)
		}

		log.Println("Download Finfished!")
	})
	cDownload.Visit(source.Link)

	return "assets/movies/" + source.Name
}
