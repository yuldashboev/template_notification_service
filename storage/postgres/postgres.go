package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.udevs.io/template/template_notification_service/config"
	"gitlab.udevs.io/template/template_notification_service/pkg/logger"
	"gitlab.udevs.io/template/template_notification_service/storage"
)

type Store struct {
	db           *pgxpool.Pool
	log          logger.Logger
	notification storage.NotificationStorageI
}

func NewPostgres(ctx context.Context, cfg config.Config, log logger.Logger) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(cfg.MaxPgxPoolConnections)

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:  pool,
		log: log,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Notification() storage.NotificationStorageI {
	if s.notification == nil {
		s.notification = NewNotificationRepo(s.db)
	}
	return s.notification
}
