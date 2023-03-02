package usecases_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sanyarise/playlist/internal/models"
	"github.com/sanyarise/playlist/internal/repository/mocks"
	"github.com/sanyarise/playlist/internal/usecases"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockSongStore(ctrl)
	logger := zap.L().Sugar()
	usecase := usecases.NewSongUsecase(store, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testSong := &models.Song{
		Title:    "test",
		Duration: 1,
	}
	testId := uuid.New()

	res, err := usecase.CreateSong(ctxWithCancel, testSong)
	assert.Error(t, err)
	assert.Equal(t, res, uuid.Nil)
	assert.Equal(t, err, fmt.Errorf("context closed"))

	store.EXPECT().CreateSong(ctx, testSong).Return(uuid.Nil, fmt.Errorf("error"))
	res, err = usecase.CreateSong(ctx, testSong)
	assert.Error(t, err)
	assert.Equal(t, res, uuid.Nil)

	store.EXPECT().CreateSong(ctx, testSong).Return(testId, nil)
	res, err = usecase.CreateSong(ctx, testSong)
	assert.NoError(t, err)
	assert.Equal(t, res, testId)
}

func TestGetSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockSongStore(ctrl)
	logger := zap.L().Sugar()
	usecase := usecases.NewSongUsecase(store, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}

	res, err := usecase.GetSong(ctxWithCancel, testId)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, err, fmt.Errorf("context closed"))

	store.EXPECT().GetSong(ctx, testId).Return(nil, fmt.Errorf("error"))
	res, err = usecase.GetSong(ctx, testId)
	assert.Error(t, err)
	assert.Nil(t, res)

	store.EXPECT().GetSong(ctx, testId).Return(testSong, nil)
	res, err = usecase.GetSong(ctx, testId)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res, testSong)
}

func TestUpdateSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockSongStore(ctrl)
	logger := zap.L().Sugar()
	usecase := usecases.NewSongUsecase(store, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}

	err := usecase.UpdateSong(ctxWithCancel, testSong)
	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("context closed"))

	store.EXPECT().UpdateSong(ctx, testSong).Return(fmt.Errorf("error"))
	err = usecase.UpdateSong(ctx, testSong)
	assert.Error(t, err)

	store.EXPECT().UpdateSong(ctx, testSong).Return(nil)
	err = usecase.UpdateSong(ctx, testSong)
	assert.NoError(t, err)
}

func TestDeleteSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockSongStore(ctrl)
	logger := zap.L().Sugar()
	usecase := usecases.NewSongUsecase(store, logger)
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	testId := uuid.New()

	err := usecase.DeleteSong(ctxWithCancel, testId)
	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("context closed"))

	store.EXPECT().DeleteSong(ctx, testId).Return(fmt.Errorf("error"))
	err = usecase.DeleteSong(ctx, testId)
	assert.Error(t, err)

	store.EXPECT().DeleteSong(ctx, testId).Return(nil)
	err = usecase.DeleteSong(ctx, testId)
	assert.NoError(t, err)
}


