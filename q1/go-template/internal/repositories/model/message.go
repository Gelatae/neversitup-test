package model

import (
	"go-template/internal/core/domain"
)

// ///////////////////////// MESSAGE //////////////////////////
type Message struct {
	Id      int    `db:"id"`
	Message string `db:"message"`
}

func NewMessage(m domain.Message) (message Message, err error) {
	message = Message{
		Id:      m.Id,
		Message: m.Message,
	}
	return
}

func (m *Message) ToDomainModel() (message domain.Message) {
	message = domain.Message{
		Id:      m.Id,
		Message: m.Message,
	}
	return
}
