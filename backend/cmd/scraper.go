package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Media struct {
	Name string
	Link string
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
		movies = append(movies, Media{element.Text, element.Attr("href")})
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
			_movies = append(_movies, Media{element.Text, element.Attr("href")})
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
		ext := "." + exts[len(exts)-1]
		sources = append(sources, Media{movie.Name + ext, element.Attr("href")})
	})
	cSourceFinder.Visit("https://www.vadapav.mov" + movie.Link)
	return sources
}

func SaveMovie(source Media) string {
	fileName := source.Name

	response, err := http.Get("https://www.vadapav.mov" + source.Link)
	file, err2 := os.Create("assets/movies/" + fileName)

	if err != nil {
		log.Fatalln(err.Error())
	} else if err2 != nil {
		log.Fatalln(err2.Error())
	}

	totalBytes := response.ContentLength
	var downloadedBytes int64 = 0

	//                     3 mb buffer size
	buffer := make([]byte, 3072*1024)

	for {
		n, err := response.Body.Read(buffer)

		if n > 0 {
			_, werr := file.Write(buffer[:n])
			if werr != nil {
				log.Fatalln(werr.Error())
			}
			downloadedBytes += int64(n)

			log.Printf("Downloaded %d %%, (%d mb / %d mb) \n", (100 * (downloadedBytes / totalBytes)), (downloadedBytes / (1024 * 1024)), (totalBytes / (1024 * 1024)))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	log.Println("download completed")
	return "assets/movies/" + source.Name
}
