package repository

import (
	"context"
	"github.com/sanyarise/playlist/internal/models"

	"github.com/google/uuid"
)

type SongStore interface {
	CreateSong(ctx context.Context, song *models.Song) (uuid.UUID, error)
	GetSong(ctx context.Context, id uuid.UUID) (*models.Song, error)
	UpdateSong(ctx context.Context, song *models.Song) error
	DeleteSong(ctx context.Context, id uuid.UUID) error
}
