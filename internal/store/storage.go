package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound       = errors.New("record not found")
	ErrConflict       = errors.New("resource already exists")
	QueryTimeDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Update(context.Context, *Post) error
		Delete(context.Context, int64) error
		GetUserFeed(context.Context, int64) ([]PostWithMetadata, error)
	}
	Users interface {
		GetByID(context.Context, int64) (*User, error)
		Create(context.Context, *User) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerID, user_ID int64) error
		Unfollow(ctx context.Context, followerID, user_ID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
