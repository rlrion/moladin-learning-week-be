package songs

import (
	"github.com/gin-gonic/gin"
	"github.com/rlrion/moladin-learning-week-be/dto"
	"github.com/rlrion/moladin-learning-week-be/entity"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	DB *gorm.DB
}

func (h RequestHandler) CreateSong(c *gin.Context) {
	var vSongDataInput SongDataInput

	if err := c.Bind(&vSongDataInput); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid payload"})
		return
	}

	data := entity.Song{Title: vSongDataInput.Title, Author: vSongDataInput.Author, AlbumId: vSongDataInput.AlbumId}

	h.DB.Create(&data)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: data})
}

func (h RequestHandler) GetSongs(c *gin.Context) {
	var vSong []entity.Song

	h.DB.Find(&vSong)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: vSong})
}

func (h RequestHandler) GetSongDetail(c *gin.Context) {
	var vSong []entity.Song

	id := c.Params.ByName("id")

	if err := h.DB.Where("album_id = ?", id).First(&vSong).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "song not found"})
		return
	}

	h.DB.Where("album_id = ?", id).Find(&vSong)

	data := make([]SongData, len(vSong))

	for i, song := range vSong {
		data[i] = SongData{
			ID:        song.ID,
			AlbumId:   song.AlbumId,
			Title:     song.Title,
			Author:    song.Author,
			CreatedAt: song.CreatedAt,
			UpdatedAt: song.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: data})
}

func (h RequestHandler) UpdateSong(c *gin.Context) {
	var vSong entity.Song

	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&vSong).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "Record not found"})
		return
	}

	c.BindJSON(&vSong)

	h.DB.Save(&vSong)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: vSong})
}

func (h RequestHandler) DeleteSong(c *gin.Context) {
	var vSong entity.Song

	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&vSong).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "Song not found"})
		return
	}

	h.DB.Delete(vSong, id)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: true})
}
