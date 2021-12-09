package repository

import (
	"context"

	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/domain/model"
)

type PlaylistRepository interface {
	// Create(playlist *model.Playlist) (*model.Playlist, error)
	Get(ctx context.Context) (model.Playlist, error)
}
