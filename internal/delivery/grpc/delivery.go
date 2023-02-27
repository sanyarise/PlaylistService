package grpc

import (
	"context"

	playlist "github.com/sanyarise/playlist/internal/delivery/grpc/interface"
	"github.com/sanyarise/playlist/internal/models"
	"github.com/sanyarise/playlist/internal/usecases"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Delivery struct {
	usecase usecases.ISongUsecase
	logger  *zap.SugaredLogger
	playlist.UnimplementedPlaylistServer
}

func NewDelivery(usecase usecases.ISongUsecase, logger *zap.SugaredLogger) *Delivery {
	logger.Debug("Enter in NewDelivery()")
	return &Delivery{
		usecase: usecase,
		logger:  logger,
	}
}

func (d *Delivery) CreateSong(ctx context.Context, req *playlist.CreateSongRequest) (*playlist.CreateSongResponse, error) {
	d.logger.Debug("Enter in delivery CreateSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		song := &models.Song{
			Title:    req.GetTitle(),
			Duration: req.GetDuration(),
		}
		id, err := d.usecase.CreateSong(ctx, song)
		if err != nil {
			d.logger.Error("error on create song: %v", err)
			return nil, status.Errorf(codes.Internal, "error on create song: %v", err)
		}
		d.logger.Info("song with id %v created success", id)
		return &playlist.CreateSongResponse{
			Id: song.Id.String(),
		}, nil
	}
}
