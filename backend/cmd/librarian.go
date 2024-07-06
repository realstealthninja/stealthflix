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
	Download int64
	Size     int64
}

func InitDb() {
	db, err := gorm.Open(sqlite.Open("db/media.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&Download{})
}

func Insert(media Media, path string, download int64, size int64) {
	db.Create(&Download{Media: media, Path: path, Download: download, Size: size})
}

func DownloadStatus(media Media) Download {
	var download Download
	db.Where("media_name = ? AND media_link = ?", media.Name, media.Link).First(&download)
	return download
}

func SetDownload(media Media, sizeDownloaded int64) {
	var download Download
	db.Where("media_name = ? AND media_link = ?", media.Name, media.Link).Find(&download)
	db.Model(&download).Update("Download", sizeDownloaded)
}
