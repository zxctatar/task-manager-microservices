package sessionvalidator

import "context"

type SessionValidator interface {
	GetIdBySession(ctx context.Context, sessionId string) (uint32, error)
}
