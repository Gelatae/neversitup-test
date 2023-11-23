package services

import (
	"context"
	"go-template/internal/core/domain"
	ports "go-template/internal/core/ports/repositories"
)

type messageService struct {
	messageRepository ports.MessageRepository
}

func NewMessageService(mr ports.MessageRepository) *messageService {
	return &messageService{
		messageRepository: mr,
	}
}

func (s *messageService) GetMessage(ctx context.Context, id int) (domain.Message, error) {
	message, err := s.messageRepository.GetMessage(ctx, id)
	if err != nil {
		return domain.Message{}, err
	}
	return message, nil
}
