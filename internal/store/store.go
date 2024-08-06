package store

import (
	"context"
	"time"
)

// Store описывает абстрактное хранилище сообщений пользователей
type Store interface {
	// FindRecepient возвращает внутренний идентификатор пользователя по человекопонятному имени
	FindRecepient(ctx context.Context, username string) (userID string, err error)
	// ListMessages возвращает список всех сообщений при определенном получаетеле
	ListMessages(ctx context.Context, userID string) ([]Message, error)
	// GetMessage возвращает сообщение с определенным ID
	GetMessage(ctx context.Context, id int64) (*Message, error)
	// SaveMessage сохраняет новое сообщение
	SaveMessage(ctx context.Context, userID string, msg Message) error
}

// Message описывает объект общения
type Message struct {
	ID      int64     // внутренний идентификатор общения
	Sender  string    // отправитель
	Time    time.Time // время отправления
	Payload string    // текст сообщения
}
