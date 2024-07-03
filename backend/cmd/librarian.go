package cmd

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Download struct {
	gorm.Model
	Media    Media `gorm:embedded;embeddedPrefix:media_`
	Path     string
	Download uint
	Size     uint
}

func InitDb() {
	db, err := gorm.Open(sqlite.Open("db/media.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&Download{})

}

func insert(media Media, path string, download uint, size uint) {
	db.Create(&Download{Media: media, Path: path, Download: download, Size: size})
}

func download_status(media Media) Download {
	var download Download
	db.Where("media_name = ? AND media_link = ?", media.Name, media.Link).First(&download)
	return download
}

func set_download(media Media, sizeDownloaded uint) {
	var download Download
	db.Where("media_name = ? AND media_link = ?", media.Name, media.Link).Find(&download)
	db.Model(&download).Update("Download", sizeDownloaded)
}
