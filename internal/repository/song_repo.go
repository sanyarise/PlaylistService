package repository

import (
	"context"
	"fmt"
	"playlist/internal/models"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var _ SongStore = (*songRepo)(nil)

type songRepo struct {
	db     *PGres
	logger *zap.SugaredLogger
}

func NewSongRepo(db *PGres, logger *zap.SugaredLogger) SongStore {
	logger.Debug("Enter in NewSongRepo()")
	return &songRepo{
		db:     db,
		logger: logger,
	}
}

// CreateSong add new song in database
func (repo *songRepo) CreateSong(ctx context.Context, song *models.Song) (uuid.UUID, error) {
	repo.logger.Debugf("Enter in repository CreateSong() with args: ctx, song: %v", song)
	select {
	case <-ctx.Done():
		return uuid.Nil, fmt.Errorf("context closed")
	default:
		pool := repo.db.GetPool()

		// Recording operations need transaction
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			repo.logger.Errorf("Can't create transaction: %s", err)
			return uuid.Nil, fmt.Errorf("can't create transaction: %w", err)
		}
		repo.logger.Debug("Transaction begin success")
		defer func() {
			if err != nil {
				repo.logger.Errorf("Transaction rolled back")
				if err = tx.Rollback(ctx); err != nil {
					repo.logger.Errorf("Can't rollback %s", err)
				}

			} else {
				repo.logger.Info("Transaction commited")
				if err != tx.Commit(ctx) {
					repo.logger.Errorf("Can't commit %s", err)
				}
			}
		}()
		var id uuid.UUID
		row := tx.QueryRow(ctx, `INSERT INTO songs(title, duration) values ($1, $2) RETURNING id`,
			song.Title,
			song.Duration,
		)
		err = row.Scan(&id)
		if err != nil {
			repo.logger.Errorf("can't create song %s", err)
			return uuid.Nil, fmt.Errorf("can't create song %w", err)
		}
		repo.logger.Info("Song create success")
		repo.logger.Debugf("id is %v\n", id)
		return id, nil
	}
}

// GetSong returns song by id
func (repo *songRepo) GetSong(ctx context.Context, id uuid.UUID) (*models.Song, error) {
	repo.logger.Debug("Enter in repository GetSong() with args: ctx, id: %v", id)
	select {
	case <-ctx.Done():
		return &models.Song{}, fmt.Errorf("context closed")
	default:
		pool := repo.db.GetPool()

		song := models.Song{}
		row := pool.QueryRow(ctx, `
	SELECT id, title, duration FROM songs WHERE id=$1
	`, id)
		err := row.Scan(
			&song.Id,
			&song.Title,
			&song.Duration,
		)
		if err != nil && strings.Contains(err.Error(), "no rows in result set") {
			repo.logger.Errorf("Error in rows scan get song by id: %s", err)
			return &models.Song{}, models.ErrorNotFound{}
		} else if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
			repo.logger.Errorf("Error in rows scan get song by id: %s", err)
			return &models.Song{}, fmt.Errorf("error in rows scan get song by id: %w", err)
		}
		repo.logger.Info("Get song success")
		return &song, nil
	}
}

// UpdateSong updated song in database
func (repo *songRepo) UpdateSong(ctx context.Context, song *models.Song) error {
	repo.logger.Debugf("Enter in repository UpdateSong() with args: ctx, song: %v", song)
	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	default:
		pool := repo.db.GetPool()
		// Recording operations need transaction
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			repo.logger.Errorf("Can't create transaction: %s", err)
			return fmt.Errorf("can't create transaction: %w", err)
		}
		repo.logger.Debug("Transaction begin success")
		defer func() {
			if err != nil {
				repo.logger.Errorf("Transaction rolled back")
				if err = tx.Rollback(ctx); err != nil {
					repo.logger.Errorf("Can't rollback %s", err)
				}

			} else {
				repo.logger.Info("Transaction commited")
				if err != tx.Commit(ctx) {
					repo.logger.Errorf("Can't commit %s", err)
				}
			}
		}()

		_, err = tx.Exec(ctx, `UPDATE songs SET title=$1, duration=$2 WHERE id=$3`,
			song.Title,
			song.Duration,
			song.Id,
		)
		if err != nil && strings.Contains(err.Error(), "no rows in result set") {
			repo.logger.Errorf("Error on update song %s: %s", song.Id, err)
			return models.ErrorNotFound{}
		}
		if err != nil {
			repo.logger.Errorf("Error on update song %s: %s", song.Id, err)
			return fmt.Errorf("error on update song %s: %w", song.Id, err)
		}
		repo.logger.Infof("Song with id %s successfully updated", song.Id)
		return nil
	}
}

// DeleteSong deleted song from database by id
func (repo *songRepo) DeleteSong(ctx context.Context, id uuid.UUID) error {
	repo.logger.Debug("Enter in repository cart DeleteSong() with args: ctx, id: %v", id)
	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	default:
		pool := repo.db.GetPool()
		// Removal operation is carried out in transaction
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			repo.logger.Errorf("Can't create transaction: %s", err)
			return fmt.Errorf("can't create transaction: %w", err)
		}
		repo.logger.Debug("Transaction begin success")
		defer func() {
			if err != nil {
				repo.logger.Errorf("Transaction rolled back")
				if err = tx.Rollback(ctx); err != nil {
					repo.logger.Errorf("Can't rollback %s", err)
				}

			} else {
				repo.logger.Info("Transaction commited")
				if err != tx.Commit(ctx) {
					repo.logger.Errorf("Can't commit %s", err)
				}
			}
		}()
		_, err = tx.Exec(ctx, `DELETE FROM songs WHERE id=$1`, id)
		if err != nil && strings.Contains(err.Error(), "no rows in result set") {
			repo.logger.Errorf("Error on delete song %s: %s", id, err)
			return models.ErrorNotFound{}
		}
		if err != nil {
			repo.logger.Errorf("Error on delete song %s: %s", id, err)
			return fmt.Errorf("error on delete song %s: %w", id, err)
		}
		repo.logger.Infof("Song with id: %s successfully deleted from database", id)
		return nil
	}
}
