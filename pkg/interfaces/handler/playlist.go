package handler

import (
	"net/http"

	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type PlaylistHandler interface {
	Get(c *gin.Context)
}

type playlistHandler struct {
	playlistUsecase usecase.PlaylistUsecase
}

func NewPlaylistHandler(playlistUsecase usecase.PlaylistUsecase) PlaylistHandler {
	return &playlistHandler{
		playlistUsecase: playlistUsecase,
	}
}

type responsePlaylist struct {
	Id        int    `json:"id"`
	SourceUrl string `json:"source_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (ph *playlistHandler) Get(c *gin.Context) {
	playlist, err := ph.playlistUsecase.Get(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := responsePlaylist{
		Id:        playlist.Id,
		SourceUrl: playlist.SourceUrl,
		CreatedAt: playlist.CreatedAt,
		UpdatedAt: playlist.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}
