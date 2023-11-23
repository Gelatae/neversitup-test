package ports

import (
	"context"
	"go-template/internal/core/domain"
)

type MessageRepository interface {
	GetMessage(ctx context.Context, id int) (domain.Message, error)
}
