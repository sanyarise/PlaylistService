package usecases

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
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
	logger.Debug("Enter in usecases NewPlaylist()")
	return &Playlist{logger: logger}
}

// AddSong adds a song to the playlist
func (p *Playlist) AddSong(ctx context.Context, song *models.Song) {
	p.logger.Debugf("Enter in usecases AddSong with args: ctx, song: %v", song)
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
	p.logger.Debug("Enter in usecases Play()")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancel")
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		if p.paused {
			p.paused = false
			return nil
		}

		if p.playing {
			p.logger.Info("already playing")
			return fmt.Errorf("already playing")
		}
		if p.current == nil && p.head == nil {
			p.logger.Error("playlist is empty")
			return fmt.Errorf("playlist is empty")
		}
		p.playing = true

		go func() {
			p.logger.Debug("Enter in Play go func()")
			for {
				p.logger.Debug("Enter in for in Play go func()")
				if p.finished {
					p.current = nil
					p.playing = false
					p.finished = false
					return
				}

				if p.current == nil {
					p.current = p.head
				} else {
					if p.playing {
						p.current = p.current.next
					}
				}

				if p.current == nil {
					p.finished = true
					continue
				}
				if !p.playing {
					p.playing = true
				}

				p.playSong(time.Duration(p.current.song.Duration))
			}
		}()
		return nil
	}
}

// Pause pauses playback
func (p *Playlist) Pause(ctx context.Context) error {
	p.logger.Debug("Enter in usecases Pause()")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancel")
	default:
		p.mutex.Lock()
		defer p.mutex.Unlock()

		if !p.playing {
			return fmt.Errorf("not playing")
		}
		if p.paused {
			return fmt.Errorf("already paused")
		}

		p.paused = true
		return nil
	}
}

// Next switches to the next song
func (p *Playlist) Next(ctx context.Context) error {
	p.logger.Debug("Enter in usecases Next()")
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

		p.playing = false
		p.paused = false
		p.finished = false
		p.current = p.current.next
		p.logger.Debugf("p.current after next: %v", p.current.song.Title)
		return nil
	}
}

// Prev swithces to the previous song
func (p *Playlist) Prev(ctx context.Context) error {
	p.logger.Debug("Enter in usecases Prev()")
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

		p.playing = false
		p.paused = false
		p.finished = false
		p.current = p.current.prev
		p.logger.Debugf("p.current after prev: %v", p.current.song.Title)
		return nil
	}
}

// playSong simulates playback
func (p *Playlist) playSong(d time.Duration) {
	p.logger.Debugf("Enter in usecases playSong() with args: ctx, d: %v", d)
	p.logger.Infof("song with id: %s is playing. Title: %s, Duration: %s", p.current.song.Id, p.current.song.Title, p.current.song.Duration)
	finish := time.Now().Add(d * time.Minute)
	count := 0
	for {
		if !p.paused {
			if count == 0 {
				if time.Now().String() >= finish.String() {
					break
				}
			} else {
				finish = time.Now().Add(d)
				count = 0
			}
		} else {
			if count == 0 {
				d = time.Until(finish)
				count++
			}
		}
		if !p.playing {
			p.logger.Infof("song with id: %v is stopped", p.current.song.Id)
			return
		}
		if p.playing && !p.paused {
			fmt.Printf("Play song:%v\n", p.current.song.Title)
			time.Sleep(1 * time.Second)
		}
		if p.paused {
			fmt.Printf("Pause song%v\n", p.current.song.Title)
			time.Sleep(1 * time.Second)
		}
	}
	p.logger.Infof("song %s ends playback", p.current.song.Title)
}

func (p *Playlist) GetStatus(ctx context.Context) (uuid.UUID, bool) {
	p.logger.Debug("Enter in playlist GetStatus()")
	select {
	case <-ctx.Done():
		p.logger.Error(ctx.Err().Error())
		return uuid.Nil, false
	default:
		return p.current.song.Id, p.playing
	}
}
