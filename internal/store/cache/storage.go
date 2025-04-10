package cache

import (
	"context"

	"github.com/edzhabs/social/internal/store"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	User interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
	}
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return Storage{
		User: &UserStore{rdb: rdb},
	}
}
