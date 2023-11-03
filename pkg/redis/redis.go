package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Rds *redis.Client
}

func (r *Redis) Close() {
	if r.Rds != nil {
		r.Rds.Close()
	}
}

func New(ctx context.Context, addr string, password string, minIdleConns int, db int) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		MinIdleConns: minIdleConns,
		DB:           db,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	rds := &Redis{
		Rds: client,
	}

	return rds, nil
}
