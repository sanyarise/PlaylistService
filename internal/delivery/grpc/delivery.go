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
	usecase  usecases.ISongUsecase
	playlist usecases.IPlaylistUsecase
	logger   *zap.SugaredLogger
	playlist.UnimplementedPlaylistServer
}

func NewDelivery(usecase usecases.ISongUsecase, playlist usecases.IPlaylistUsecase, logger *zap.SugaredLogger) *Delivery {
	logger.Debug("Enter in NewDelivery()")
	return &Delivery{
		usecase:  usecase,
		playlist: playlist,
		logger:   logger,
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
			Title:    req.Title,
			Duration: req.Duration,
		}
		id, err := d.usecase.CreateSong(ctx, song)
		if err != nil {
			d.logger.Errorf("error on create song: %v", err)
			return nil, status.Errorf(codes.Internal, "error on create song: %v", err)
		}
		d.logger.Infof("song with id %v created success", id)
		return &playlist.CreateSongResponse{
			Id: id.String(),
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
		id, err := uuid.Parse(req.Id)
		if err != nil {
			d.logger.Errorf("error on get song: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "error on get song: %v", err)
		}
		modelsSong, err := d.usecase.GetSong(ctx, id)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Errorf("error on get song: song with id: %v not found", id)
			return nil, status.Errorf(codes.NotFound, "error on get song: song with id: %v not found", id)
		}
		if err != nil {
			d.logger.Errorf("error on get song: %v", err)
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
		id, err := uuid.Parse(req.Id)
		if err != nil {
			d.logger.Errorf("error on update song: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "error on update song: %v", err)
		}
		updatedSong := &models.Song{
			Id:       id,
			Title:    req.Title,
			Duration: req.Duration,
		}
		err = d.usecase.UpdateSong(ctx, updatedSong)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Errorf("error on update song: song with id: %v not found", id)
			return nil, status.Errorf(codes.NotFound, "error on update song: song with id: %v not found", id)
		}
		if err != nil {
			d.logger.Errorf("error on update song: %v", err)
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
		id, err := uuid.Parse(req.Id)
		if err != nil {
			d.logger.Errorf("error on delete song: id %v is not correct. error: %v", req.Id, err)
			return nil, status.Errorf(codes.InvalidArgument, "error on delete song: %v", err)
		}
		currentSongId, isPlaying := d.playlist.GetStatus(ctx)
		if currentSongId == id && isPlaying {
			d.logger.Errorf("error on delete song with id: %v. Song is playing now", req.Id)
			return nil, status.Errorf(codes.Aborted, "error on delete song with id: %v. Song is playing now", req.Id)
		}
		err = d.usecase.DeleteSong(ctx, id)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Errorf("error on delete song: song with id: %v not found", id)
			return nil, status.Errorf(codes.NotFound, "error on delete song: song with id: %v not found", id)
		}
		if err != nil {
			d.logger.Errorf("error on delete song: %v", err)
			return nil, status.Errorf(codes.Internal, "error on delete song: %v", err)
		}
		return &playlist.DeleteSongResponse{}, nil
	}
}

// AddSong calls usecase method for add song into playlist
func (d *Delivery) AddSong(ctx context.Context, req *playlist.AddSongRequest) (*playlist.AddSongResponse, error) {
	d.logger.Debug("Enter in delivery AddSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		id, err := uuid.Parse(req.Id)
		if err != nil {
			d.logger.Errorf("error on add song: id %v is not correct. error: %v", req.Id, err)
			return nil, status.Errorf(codes.InvalidArgument, "error on add song: %v", err)
		}

		song, err := d.usecase.GetSong(ctx, id)
		if err != nil && errors.Is(models.ErrorNotFound{}, err) {
			d.logger.Errorf("error on add song with id: %s - song not found", req.Id)
			return nil, status.Errorf(codes.NotFound, "error on add song: song with id: %s not found", req.Id)
		}
		if err != nil {
			d.logger.Errorf("error on add song: %v", err)
			return nil, status.Errorf(codes.Internal, "error on add song: %v", err)
		}
		d.playlist.AddSong(ctx, song)
		d.logger.Infof("Song with id: %s add in playlist success", req.Id)
		return &playlist.AddSongResponse{}, nil
	}
}

// PlaySong calls usecase method to play
func (d *Delivery) PlaySong(ctx context.Context, req *playlist.PlaySongRequest) (*playlist.PlaySongResponse, error) {
	d.logger.Debug("Enter in delivery PlaySong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		err := d.playlist.Play(ctx)
		if err != nil {
			d.logger.Error(err.Error())
			return nil, status.Errorf(codes.Internal, "error on play: %v", err)
		}
		d.logger.Info("Play song success")
		return &playlist.PlaySongResponse{}, nil
	}
}

// PauseSong calls usecase method for pause
func (d *Delivery) PauseSong(ctx context.Context, req *playlist.PauseSongRequest) (*playlist.PauseSongResponse, error) {
	d.logger.Debug("Enter in delivery PauseSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		err := d.playlist.Pause(ctx)
		if err != nil {
			d.logger.Error(err.Error())
			return nil, status.Errorf(codes.Internal, "error on paused: %v", err)
		}
		d.logger.Info("Pause song success")
		return &playlist.PauseSongResponse{}, nil
	}
}

// NextSong calls usecase method to switch to the next song
func (d *Delivery) NextSong(ctx context.Context, req *playlist.NextSongRequest) (*playlist.NextSongResponse, error) {
	d.logger.Debug("Enter in delivery NextSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		err := d.playlist.Next(ctx)
		if err != nil {
			d.logger.Error(err.Error())
			return nil, status.Errorf(codes.Internal, "error on next song: %v", err)
		}
		d.logger.Info("Swithing to the next song success")
		return &playlist.NextSongResponse{}, nil
	}
}

// PrevSong calls usecase method to switch to the prev song
func (d *Delivery) PrevSong(ctx context.Context, req *playlist.PrevSongRequest) (*playlist.PrevSongResponse, error) {
	d.logger.Debug("Enter in delivery PrevSong()")
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "context is canceled")
	default:
		err := d.playlist.Prev(ctx)
		if err != nil {
			d.logger.Error(err.Error())
			return nil, status.Errorf(codes.Internal, "error on prev song: %v", err)
		}
		d.logger.Info("Swithting to the previous song success")
		return &playlist.PrevSongResponse{}, nil
	}
}
