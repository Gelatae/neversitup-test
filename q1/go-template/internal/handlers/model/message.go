package model

import (
	"go-template/internal/core/domain"
)

type Message struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func (m *Message) ToDomainModel() domain.Message {
	msg := domain.Message{
		Id:      m.Id,
		Message: m.Message,
	}
	return msg
}

func NewMessage(m domain.Message) Message {
	return Message{
		Id:      m.Id,
		Message: m.Message,
	}
}
