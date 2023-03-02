package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/sanyarise/playlist/internal/models"
)

type ISongUsecase interface {
	CreateSong(ctx context.Context, song *models.Song) (uuid.UUID, error)
	GetSong(ctx context.Context, id uuid.UUID) (*models.Song, error)
	UpdateSong(ctx context.Context, song *models.Song) error
	DeleteSong(ctx context.Context, id uuid.UUID) error
}

type IPlaylistUsecase interface {
	AddSong(ctx context.Context, song *models.Song)
	DeleteSong(ctx context.Context, id uuid.UUID) error
	Play(ctx context.Context) error
	Pause(ctx context.Context) error
	Next(ctx context.Context) error
	Prev(ctx context.Context) error
}
