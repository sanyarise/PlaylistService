package usecases

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/sanyarise/playlist/internal/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAddSong(t *testing.T) {
	l := zap.L().Sugar()
	idx := make(map[uuid.UUID]struct{})
	p := new(Playlist)
	p.idx = idx
	p.logger = l
	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}
	testId2 := uuid.New()
	testSong2 := &models.Song{
		Id:       testId2,
		Title:    "test2",
		Duration: 0,
	}
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	cancel()
	p.AddSong(ctxWithTimeout, testSong)

	p.AddSong(ctx, testSong)
	_, ok := p.idx[testId]
	assert.True(t, ok)
	assert.Equal(t, testSong, p.head.song)
	assert.Equal(t, p.tail, p.head)

	p.AddSong(ctx, testSong2)
	_, ok = p.idx[testId2]
	assert.True(t, ok)
	assert.Equal(t, p.tail.prev, p.head)
	assert.Equal(t, p.head.next, p.tail)
}

func TestPlay(t *testing.T) {
	l := zap.L().Sugar()
	idx := make(map[uuid.UUID]struct{})
	p := new(Playlist)
	p.idx = idx
	p.logger = l
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	cancel()

	err := p.Play(ctxWithTimeout)
	assert.Error(t, err)

	p.paused = true
	err = p.Play(ctx)
	assert.NoError(t, err)

	p.paused = false
	p.playing = true
	err = p.Play(ctx)
	assert.Error(t, err)

	p.playing = false
	err = p.Play(ctx)
	assert.Error(t, err)

	p.finished = true
	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}
	p.AddSong(ctx, testSong)
	err = p.Play(ctx)
	assert.NoError(t, err)

	p.finished = false
	p.playing = false
	err = p.Play(ctx)
	assert.NoError(t, err)
	assert.True(t, p.playing)
}

func TestPause(t *testing.T) {
	l := zap.L().Sugar()
	idx := make(map[uuid.UUID]struct{})
	p := new(Playlist)
	p.idx = idx
	p.logger = l
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	cancel()

	err := p.Pause(ctxWithTimeout)
	assert.Error(t, err)

	p.playing = false
	err = p.Pause(ctx)
	assert.Error(t, err)

	p.playing = false
	p.paused = true
	err = p.Pause(ctx)
	assert.Error(t, err)

	p.playing = true
	p.paused = false
	err = p.Pause(ctx)
	assert.NoError(t, err)
	assert.True(t, p.paused)
}

func TestNext(t *testing.T){
	l := zap.L().Sugar()
	idx := make(map[uuid.UUID]struct{})
	p := new(Playlist)
	p.idx = idx
	p.logger = l
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	cancel()

	err := p.Next(ctxWithTimeout)
	assert.Error(t, err)

	p.playing = false
	err = p.Next(ctx)
	assert.Error(t, err)

	p.playing = true
	err = p.Next(ctx)
	assert.Error(t, err)

	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}
	p.AddSong(ctx, testSong)
	p.current = p.head
	p.current.next = nil
	err = p.Next(ctx)
	assert.Error(t, err)

	testId2 := uuid.New()
	testSong2 := &models.Song{
		Id:       testId2,
		Title:    "test2",
		Duration: 0,
	}
	p.AddSong(ctx, testSong2)
	err = p.Next(ctx)
	assert.NoError(t, err)
	assert.False(t, p.playing)
	assert.False(t, p.paused)
	assert.False(t, p.finished)
	assert.Equal(t, "test2", p.current.song.Title)
}

func TestPrev(t *testing.T){
	l := zap.L().Sugar()
	idx := make(map[uuid.UUID]struct{})
	p := new(Playlist)
	p.idx = idx
	p.logger = l
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	cancel()

	err := p.Prev(ctxWithTimeout)
	assert.Error(t, err)

	p.playing = false
	err = p.Prev(ctx)
	assert.Error(t, err)

	p.playing = true
	err = p.Prev(ctx)
	assert.Error(t, err)

	testId := uuid.New()
	testSong := &models.Song{
		Id:       testId,
		Title:    "test",
		Duration: 0,
	}
	p.AddSong(ctx, testSong)
	p.current = p.head
	p.current.prev = nil
	err = p.Prev(ctx)
	assert.Error(t, err)

	testId2 := uuid.New()
	testSong2 := &models.Song{
		Id:       testId2,
		Title:    "test2",
		Duration: 0,
	}
	p.AddSong(ctx, testSong2)
	p.current = p.tail

	err = p.Prev(ctx)
	assert.NoError(t, err)
	assert.False(t, p.playing)
	assert.False(t, p.paused)
	assert.False(t, p.finished)
	assert.Equal(t, "test", p.current.song.Title)
}