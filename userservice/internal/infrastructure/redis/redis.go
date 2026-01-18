package myredis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
	ttl    *time.Duration
}

func NewRedis(client *redis.Client, ttl *time.Duration) *Redis {
	return &Redis{
		client: client,
		ttl:    ttl,
	}
}

func (r *Redis) Save(ctx context.Context, sessionId string, userId uint32) error {
	return r.client.Set(ctx, sessionId, userId, *r.ttl).Err()
}
