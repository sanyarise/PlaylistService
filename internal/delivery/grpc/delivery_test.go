package grpc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sanyarise/playlist/internal/delivery/grpc"
	playlist "github.com/sanyarise/playlist/internal/delivery/grpc/interface"
	"github.com/sanyarise/playlist/internal/models"
	"github.com/sanyarise/playlist/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()

	testSong := &models.Song{
		Title:    "test",
		Duration: 0,
	}
	testReq := &playlist.CreateSongRequest{
		Title:    "test",
		Duration: 0,
	}

	testId := uuid.New()
	testResp := &playlist.CreateSongResponse{
		Id: testId.String(),
	}

	res, err := delivery.CreateSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().CreateSong(ctx, testSong).Return(uuid.Nil, fmt.Errorf("error"))
	res, err = delivery.CreateSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Internal, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().CreateSong(ctx, testSong).Return(testId, nil)
	res, err = delivery.CreateSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestGetSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}
	testReq := &playlist.GetSongRequest{
		Id: testId.String(),
	}
	testReqInvalidId := &playlist.GetSongRequest{
		Id: "adfasfasf",
	}
	testResp := &playlist.GetSongResponse{
		Song: &playlist.Song{
			Id:       testId.String(),
			Title:    testSong.Title,
			Duration: testSong.Duration,
		},
	}

	res, err := delivery.GetSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	res, err = delivery.GetSong(ctx, testReqInvalidId)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().GetSong(ctx, testId).Return(nil, models.ErrorNotFound{})
	res, err = delivery.GetSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.NotFound, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().GetSong(ctx, testId).Return(nil, fmt.Errorf("error"))
	res, err = delivery.GetSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Internal, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().GetSong(ctx, testId).Return(testSong, nil)
	res, err = delivery.GetSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestUpdateSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testId := uuid.New()
	testReq := &playlist.UpdateSongRequest{
		Id:       testId.String(),
		Title:    "test",
		Duration: 0,
	}
	testReqInvalidId := &playlist.UpdateSongRequest{
		Id:       testId.String() + "a",
		Title:    "test",
		Duration: 0,
	}
	testResp := &playlist.UpdateSongResponse{}

	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}

	res, err := delivery.UpdateSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	res, err = delivery.UpdateSong(ctx, testReqInvalidId)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().UpdateSong(ctx, testSong).Return(models.ErrorNotFound{})
	res, err = delivery.UpdateSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.NotFound, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().UpdateSong(ctx, testSong).Return(fmt.Errorf("error"))
	res, err = delivery.UpdateSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Internal, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().UpdateSong(ctx, testSong).Return(nil)
	res, err = delivery.UpdateSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestDeleteSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testId := uuid.New()
	testReq := &playlist.DeleteSongRequest{
		Id: testId.String(),
	}
	testReqInvalidId := &playlist.DeleteSongRequest{
		Id: "adfasfasf",
	}
	testResp := &playlist.DeleteSongResponse{}

	res, err := delivery.DeleteSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	res, err = delivery.DeleteSong(ctx, testReqInvalidId)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().GetStatus(ctx).Return(testId, true)
	res, err = delivery.DeleteSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unavailable, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().GetStatus(ctx).Return(testId, false)
	usecase.EXPECT().DeleteSong(ctx, testId).Return(models.ErrorNotFound{})
	res, err = delivery.DeleteSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.NotFound, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().GetStatus(ctx).Return(testId, false)
	usecase.EXPECT().DeleteSong(ctx, testId).Return(fmt.Errorf("error"))
	res, err = delivery.DeleteSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Internal, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().GetStatus(ctx).Return(testId, false)
	usecase.EXPECT().DeleteSong(ctx, testId).Return(nil)
	res, err = delivery.DeleteSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestAddSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()

	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}
	testReq := &playlist.AddSongRequest{
		Id: testId.String(),
	}
	testReqInvalidId := &playlist.AddSongRequest{
		Id: testId.String() + "a",
	}
	testResp := &playlist.AddSongResponse{}

	res, err := delivery.AddSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	res, err = delivery.AddSong(ctx, testReqInvalidId)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().GetSong(ctx, testId).Return(nil, models.ErrorNotFound{})
	res, err = delivery.AddSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.NotFound, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().GetSong(ctx, testId).Return(nil, fmt.Errorf("error"))
	res, err = delivery.AddSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Internal, st.Code())
	assert.True(t, ok)

	usecase.EXPECT().GetSong(ctx, testId).Return(testSong, nil)
	playlist_usecase.EXPECT().AddSong(ctx, testSong)
	res, err = delivery.AddSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestPlaySong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testReq := &playlist.PlaySongRequest{}
	testResp := &playlist.PlaySongResponse{}

	res, err := delivery.PlaySong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Play(ctx).Return(fmt.Errorf("error"))
	res, err = delivery.PlaySong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unavailable, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Play(ctx).Return(nil)
	res, err = delivery.PlaySong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestPauseSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testReq := &playlist.PauseSongRequest{}
	testResp := &playlist.PauseSongResponse{}

	res, err := delivery.PauseSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Pause(ctx).Return(fmt.Errorf("error"))
	res, err = delivery.PauseSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unavailable, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Pause(ctx).Return(nil)
	res, err = delivery.PauseSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestNextSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testReq := &playlist.NextSongRequest{}
	testResp := &playlist.NextSongResponse{}

	res, err := delivery.NextSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Next(ctx).Return(fmt.Errorf("error"))
	res, err = delivery.NextSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unavailable, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Next(ctx).Return(nil)
	res, err = delivery.NextSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}

func TestPrevSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecase := mocks.NewMockISongUsecase(ctrl)
	playlist_usecase := mocks.NewMockIPlaylistUsecase(ctrl)
	logger := zap.L().Sugar()
	delivery := grpc.NewDelivery(usecase, playlist_usecase, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testReq := &playlist.PrevSongRequest{}
	testResp := &playlist.PrevSongResponse{}

	res, err := delivery.PrevSong(ctxWithCancel, testReq)
	st, ok := status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Canceled, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Prev(ctx).Return(fmt.Errorf("error"))
	res, err = delivery.PrevSong(ctx, testReq)
	st, ok = status.FromError(err)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unavailable, st.Code())
	assert.True(t, ok)

	playlist_usecase.EXPECT().Prev(ctx).Return(nil)
	res, err = delivery.PrevSong(ctx, testReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testResp, res)
}