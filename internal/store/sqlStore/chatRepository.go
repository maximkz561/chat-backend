package sqlstore

import (
	"chat-backend/internal/domain"
	"chat-backend/internal/store"
	"github.com/google/uuid"
)

// ChatRepositorySql ...
type ChatRepositorySql struct {
	store *UnitOfWorkSql
}

// Create ...
func (r *ChatRepositorySql) Create(c *domain.Chat) error {

	if _, err := r.store.transaction.Exec(
		"INSERT INTO chat (id, user_id, title, created_at) VALUES ($1, $2, $3, $4)",
		c.Id,
		c.UserId,
		c.Title,
		c.CreatedAt,
	); err != nil {
		return err
	}

	// TODO: refactor this to use a bulk insert
	for _, msg := range c.Messages {
		_, err := r.store.transaction.Exec(
			"INSERT INTO message (id, content, chat_id, created_at, author_type) VALUES ($1, $2, $3, $4, $5)",
			msg.Id,
			msg.Content,
			msg.ChatID,
			msg.CreatedAt,
			msg.Author,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// Find ...
func (r *ChatRepositorySql) Find(id uuid.UUID) (*domain.Chat, error) {
	rows, err := r.store.transaction.Query(`
		SELECT 
			c.id, c.user_id, c.title, c.created_at, 
			m.id, m.content, m.chat_id, m.created_at, m.author_type
		FROM chat c
		LEFT JOIN message m ON c.id = m.chat_id
		WHERE c.id = $1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var c *domain.Chat
	var msg domain.Message

	for rows.Next() {
		err := rows.Scan(
			&c.Id, &c.UserId, &c.Title, &c.CreatedAt,
			&msg.Id, &msg.Content, &msg.ChatID, &msg.CreatedAt, &msg.Author,
		)
		if err != nil {
			return nil, err
		}

		// initialize chat and append first message
		if c == nil {
			c = &domain.Chat{
				Id:        c.Id,
				UserId:    c.UserId,
				Title:     c.Title,
				CreatedAt: c.CreatedAt,
				Messages:  []domain.Message{msg},
			}
		} else {
			// add subsequent messages to the chat
			c.Messages = append(c.Messages, msg)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if c == nil {
		return nil, store.ErrRecordNotFound
	}

	return c, nil
}

// FindByUser ...
func (r *ChatRepositorySql) FindByUser(userId uuid.UUID) ([]*domain.Chat, error) {
	rows, err := r.store.transaction.Query(`
		SELECT 
			c.id, c.user_id, c.title, c.created_at, 
			m.id, m.content, m.chat_id, m.created_at, m.author_type
		FROM chat c
		LEFT JOIN message m ON c.id = m.chat_id
		WHERE c.user_id = $1
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*domain.Chat
	var msg domain.Message

	for rows.Next() {
		var c *domain.Chat
		err := rows.Scan(
			&c.Id, &c.UserId, &c.Title, &c.CreatedAt,
			&msg.Id, &msg.Content, &msg.ChatID, &msg.CreatedAt, &msg.Author,
		)
		if err != nil {
			return nil, err
		}

		// initialize chat and append first message
		if c == nil {
			c = &domain.Chat{
				Id:        c.Id,
				UserId:    c.UserId,
				Title:     c.Title,
				CreatedAt: c.CreatedAt,
				Messages:  []domain.Message{msg},
			}
		} else {
			// add subsequent messages to the chat
			c.Messages = append(c.Messages, msg)
		}

		chats = append(chats, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if chats == nil {
		return nil, store.ErrRecordNotFound
	}

	return chats, nil
}
