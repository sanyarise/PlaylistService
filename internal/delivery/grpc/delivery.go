package grpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
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

// CreateSong parsed request and calls usecase method for create song
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

// GetSong parsed request and calls usecase method for get song
func (d *Delivery) GetSong(ctx context.Context, req *playlist.GetSongRequest) (*playlist.GetSongResponse, error) {
	d.logger.Debug("Enter in delivery GetSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			d.logger.Error("error on get song: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "error on get song: %v", err)
		}
		modelsSong, err := d.usecase.GetSong(ctx, id)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Error("error on get song: song with id: %v not found", id)
			return nil, status.Errorf(codes.NotFound, "error on get song: song with id: %v not found", id)
		}
		if err != nil {
			d.logger.Error("error on get song: %v", err)
			return nil, status.Errorf(codes.Internal, "error on get song: %v", err)
		}
		song := &playlist.Song{
			Id:       modelsSong.Id.String(),
			Title:    modelsSong.Title,
			Duration: modelsSong.Duration,
		}
		return &playlist.GetSongResponse{
			Song: song,
		}, nil
	}
}

// UpdateSong parsed request and calls usecase method for update song
func (d *Delivery) UpdateSong(ctx context.Context, req *playlist.UpdateSongRequest) (*playlist.UpdateSongResponse, error) {
	d.logger.Debug("Enter in delivery UpdateSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			d.logger.Error("error on update song: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "error on update song: %v", err)
		}
		updatedSong := &models.Song{
			Id:       id,
			Title:    req.GetTitle(),
			Duration: req.GetDuration(),
		}
		err = d.usecase.UpdateSong(ctx, updatedSong)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Error("error on update song: song with id: %v not found", id)
			return nil, status.Errorf(codes.NotFound, "error on update song: song with id: %v not found", id)
		}
		if err != nil {
			d.logger.Error("error on update song: %v", err)
			return nil, status.Errorf(codes.Internal, "error un update song: %v", err)
		}
		return &playlist.UpdateSongResponse{}, nil
	}
}

// DeleteSong parsed request and calls usecase method for delete song
func (d *Delivery) DeleteSong(ctx context.Context, req *playlist.DeleteSongRequest) (*playlist.DeleteSongResponse, error) {
	d.logger.Debug("Enter in delivery DeleteSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			d.logger.Error("error on delete song: id %v is not correct. errror: %v", req.GetId(), err)
			return nil, status.Errorf(codes.InvalidArgument, "error on delete song: %v", err)
		}
		err = d.usecase.DeleteSong(ctx, id)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Error("error on delete song: song with id: %v not found", id)
			return nil, status.Errorf(codes.NotFound, "error on delete song: song with id: %v not found", id)
		}
		if err != nil {
			d.logger.Error("error on delete song: %v", err)
			return nil, status.Errorf(codes.Internal, "error on delete song: %v", err)
		}
		return &playlist.DeleteSongResponse{}, nil
	}
}
