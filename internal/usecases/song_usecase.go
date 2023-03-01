package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sanyarise/playlist/internal/models"
	"github.com/sanyarise/playlist/internal/repository"
	"go.uber.org/zap"
)

var _ ISongUsecase = &SongUsecase{}

type SongUsecase struct {
	songStore repository.SongStore
	logger    *zap.SugaredLogger
}

func NewSongUsecase(songStore repository.SongStore, logger *zap.SugaredLogger) ISongUsecase {
	logger.Debug("Enter in usecase NewSongUsecase()")

	return SongUsecase{
		songStore: songStore,
		logger:    logger,
	}
}

// CreateSong calls database method and returns id of created song or error
func (su SongUsecase) CreateSong(ctx context.Context, song *models.Song) (uuid.UUID, error) {
	su.logger.Debugf("Enter in usecase CreateSong() with args: ctx, song: %v", song)
	select {
	case <-ctx.Done():
		return uuid.Nil, fmt.Errorf("context closed")
	default:
		id, err := su.songStore.CreateSong(ctx, song)
		if err != nil {
			return uuid.Nil, err
		}
		return id, nil
	}
}

// GetSong calls database method and returns song or error
func (su SongUsecase) GetSong(ctx context.Context, id uuid.UUID) (*models.Song, error) {
	su.logger.Debugf("Enter in usecase GetSong() with args: ctx, id: %v", id)
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("context closed")
	default:
		song, err := su.songStore.GetSong(ctx, id)
		if err != nil {
			return nil, err
		}
		return song, nil
	}
}

// UpdateSong calls database method for updating song
func (su SongUsecase) UpdateSong(ctx context.Context, song *models.Song) error {
	su.logger.Debugf("Enter in usecase UpdateSong() with args: ctx, song: %v", song)
	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	default:
		return su.songStore.UpdateSong(ctx, song)
	}
}

// DeleteSong calls database method for deleting song
func (su SongUsecase) DeleteSong(ctx context.Context, id uuid.UUID) error {
	su.logger.Debugf("Enter in usecase DeleteSong() with args: ctx, id: %v", id)
	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	default:
		return su.songStore.DeleteSong(ctx, id)
	}
}
