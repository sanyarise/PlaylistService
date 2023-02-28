package usecases

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sanyarise/playlist/internal/models"
	"go.uber.org/zap"
)

var _ IPlaylistUsecase = &Playlist{}

type Playlist struct {
	head     *node
	tail     *node
	current  *node
	mutex    sync.RWMutex
	playing  bool
	paused   bool
	finished bool
	logger   *zap.SugaredLogger
}

type node struct {
	song *models.Song
	prev *node
	next *node
}

func NewPlaylist(logger *zap.SugaredLogger) IPlaylistUsecase {
	logger.Debug("Enter in models NewPlaylist")
	return &Playlist{logger: logger}
}

// AddSong adds a song to the playlist
func (p *Playlist) AddSong(ctx context.Context, song *models.Song) {
	p.logger.Debug("Enter in models AddSong with args: ctx, song: %v", song)
	select {
	case <-ctx.Done():
		p.logger.Info("Context cancel")
		return
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		node := &node{
			song: song,
		}

		if p.head == nil {
			p.head = node
			p.tail = node
			return
		}

		p.tail.next = node
		node.prev = p.tail
		p.tail = node
	}
}

// Play starts playback
func (p *Playlist) Play(ctx context.Context) error {
	p.logger.Debug("Enter in models Play()")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancel")
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		if p.playing {
			p.logger.Info("already playing")
			return fmt.Errorf("already playing")
		}

		p.playing = true

		go func() {
			for {
				if p.paused {
					time.Sleep(1 * time.Second)
					continue
				}

				if p.finished {
					p.current = nil
					p.playing = false
					p.finished = false
					return
				}

				if p.current == nil {
					p.current = p.head
				} else {
					p.current = p.current.next
				}

				if p.current == nil {
					p.finished = true
					continue
				}

				p.playSong(ctx, time.Duration(p.current.song.Duration))
			}
		}()
		return nil
	}
}

// Pause pauses playback
func (p *Playlist) Pause(ctx context.Context) error {
	p.logger.Debug("Enter in models Pause()")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancel")
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		if !p.playing {
			return fmt.Errorf("not playing")
		}

		p.paused = !p.paused
		return nil
	}
}

// Next switches to the next song
func (p *Playlist) Next(ctx context.Context) error {
	p.logger.Debug("Enter in models Next()")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancel")
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		if !p.playing {
			return fmt.Errorf("not playing")
		}
		if p.current == nil {
			return fmt.Errorf("no current song")
		}
		if p.current.next == nil {
			return fmt.Errorf("no next song")
		}

		p.current = p.current.next
		p.paused = false
		p.finished = false
		return nil
	}
}

// Prev swithces to the previous song
func (p *Playlist) Prev(ctx context.Context) error {
	p.logger.Debug("Enter in models Prev()")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancel")
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		if !p.playing {
			return fmt.Errorf("not playing")
		}
		if p.current == nil {
			return fmt.Errorf("no current song")
		}
		if p.current.prev == nil {
			return fmt.Errorf("no previous song")
		}

		p.current = p.current.prev
		p.paused = false
		p.finished = false
		return nil
	}
}

// playSong simulates playback
func (p *Playlist) playSong(ctx context.Context, d time.Duration) {
	p.logger.Debug("Enter in models playSong() with args: ctx, d: %v", d)
	select {
	case <-ctx.Done():
		p.logger.Info("context cancel")
		return
	default:
		time.Sleep(d)
	}
}
