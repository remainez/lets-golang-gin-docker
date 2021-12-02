package infrastructure

import (
	"context"

    "github.com/nagaihiroya/lets-golang-gin-docker/pkg/domain/model"
	"github.com/nagaihiroya/lets-golang-gin-docker/pkg/domain/repository"

    // "github.com/jinzhu/gorm"
)

type PlaylistRepository struct {
    // Conn *gorm.DB
}

// func NewPlaylistRepository(conn *gorm.DB) repository.PlaylistRepository {
//     return &PlaylistRepository{Conn: conn}
// }

func NewPlaylistRepository() repository.PlaylistRepository {
    return &PlaylistRepository{}
}

// func (pr *PlaylistRepository) Create(playlist *model.Playlist) (*model.Playlist, error) {
//     if err := pr.Conn.Create(&Playlist).Error; err != nil {
//         return nil, err
//     }

//     return Playlist, nil
// }

func (pr *PlaylistRepository) Get(ctx context.Context) (model.Playlist, error) {
    // if err := pr.Conn.Get().Error; err != nil {
        // return nil, err
    // }

    // return Playlist, nil
    music1 := model.Playlist{Id: 1, SourceUrl: "https://example.com/", CreatedAt: "2021-12-12 00:00:00", UpdatedAt: "2021-12-12 00:00:00"}
	// music2 := model.Playlist{Id: 2, SourceUrl: "https://example.com/", CreatedAt: "2021-12-12 00:00:00", UpdatedAt: "2021-12-12 00:00:00"}

	return music1, nil
}