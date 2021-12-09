package usecase

import (
	"context"

	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/domain/model"
	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/domain/repository"
)

type PlaylistUsecase interface {
	Get(ctx context.Context) (model.Playlist, error)
}

type playlistUsecase struct {
	playlistRepository repository.PlaylistRepository
}

func NewPlaylistUsecase(playlistRepository repository.PlaylistRepository) PlaylistUsecase {
	return &playlistUsecase{playlistRepository: playlistRepository}
}

func (pu playlistUsecase) Get(ctx context.Context) (playlists model.Playlist, err error) {
	playlists, err = pu.playlistRepository.Get(ctx)
	if err != nil {
		return playlists, err
	}
	return playlists, nil
}
