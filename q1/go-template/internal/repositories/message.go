package repositories

import (
	"context"
	"go-template/internal/core/domain"
	"go-template/internal/repositories/model"

	"github.com/jmoiron/sqlx"
)

type messageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (mr *messageRepository) GetMessage(ctx context.Context, id int) (domain.Message, error) {
	m := model.Message{Id: id, Message: "Hello World"}
	// TODO: implement query here
	// example:
	// sql := `SELECT .....`
	// err = mr.db.SelectContext(ctx, &m, sql)
	return m.ToDomainModel(), nil
}
