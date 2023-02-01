package albums

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

func (h RequestHandler) CreateAlbum(c *gin.Context) {
	var vAlbumDataInput AlbumDataInput

	if err := c.Bind(&vAlbumDataInput); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid payload"})
		return
	}

	data := entity.Album{Name: vAlbumDataInput.Name, Year: vAlbumDataInput.Year}

	h.DB.Create(&data)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: data})
}

func (h RequestHandler) GetAlbums(c *gin.Context) {
	var albums []entity.Album

	h.DB.Find(&albums)

	var data []Albumitem
	for _, album := range albums {
		data = append(data, Albumitem{
			ID:        album.ID,
			Name:      album.Name,
			Year:      album.Year,
			CreatedAt: album.CreatedAt,
			UpdatedAt: album.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: data})
}

func (h RequestHandler) GetAlbumsDetail(c *gin.Context) {
	var vAlbum entity.Album

	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&vAlbum).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "album not found"})
		return
	}

	h.DB.Preload("Songs").First(&vAlbum, id)

	data := AlbumData{
		ID:        vAlbum.ID,
		Name:      vAlbum.Name,
		Songs:     NewSongList(vAlbum.Songs),
		Year:      vAlbum.Year,
		CreatedAt: vAlbum.CreatedAt,
		UpdatedAt: vAlbum.UpdatedAt,
	}

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: data})
}

func (h RequestHandler) UpdateAlbum(c *gin.Context) {
	var vAlbum entity.Album

	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&vAlbum).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "Record not found!"})
		return
	}

	c.BindJSON(&vAlbum)

	h.DB.Save(&vAlbum)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: vAlbum})
}

func (h RequestHandler) DeleteAlbum(c *gin.Context) {
	var a entity.Album

	id := c.Params.ByName("id")

	if err := h.DB.Where("id = ?", id).First(&a).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: "invalid id"})
		return
	}

	h.DB.Delete(a, id)

	c.JSON(http.StatusOK, dto.Response{Message: "success", Data: true})
}
