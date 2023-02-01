package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rlrion/moladin-learning-week-be/entity"
	"github.com/rlrion/moladin-learning-week-be/modules/albums"
	"github.com/rlrion/moladin-learning-week-be/modules/songs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	host := gin.Default()

	dbHost := "root@tcp(localhost:3306)/db_moladin_project_learning_week?charset=latin1&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dbHost), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logger.Info)),
	})
	if err != nil {
		panic("failed connection")
	}

	var album entity.Album
	var song entity.Song

	err = db.AutoMigrate(album, song)
	if err != nil {
		log.Println(fmt.Errorf("db.AutoMigrate: %w", err))
		return
	}

	userRequestHandler := albums.RequestHandler{
		DB: db,
	}
	productRequestHandler := songs.RequestHandler{
		DB: db,
	}

	host.POST("/albums", userRequestHandler.CreateAlbum)
	host.GET("/albums", userRequestHandler.GetAlbums)
	host.GET("/albums/:id", userRequestHandler.GetAlbumsDetail)
	host.PUT("/albums/:id", userRequestHandler.UpdateAlbum)
	host.DELETE("/albums/:id", userRequestHandler.DeleteAlbum)

	host.POST("/songs", productRequestHandler.CreateSong)
	host.GET("/songs", productRequestHandler.GetSongs)
	host.GET("/songs/:id", productRequestHandler.GetSongDetail)
	host.PUT("/songs/:id", productRequestHandler.UpdateSong)
	host.DELETE("/songs/:id", productRequestHandler.DeleteSong)

	host.Run()
}
