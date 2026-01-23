package session

import "context"

type SessionRepo interface {
	Save(ctx context.Context, sessionId string, userId uint32) error
	Get(ctx context.Context, sessionId string) (uint32, error)
}
