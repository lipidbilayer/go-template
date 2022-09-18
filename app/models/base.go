package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10/orm"
)

type BaseModel struct {
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time `sql:"default:now()"`
	DeletedAt time.Time
}

// BeforeInsert - update createdAt and updatedAt
func (m *BaseModel) BeforeInsert(ctx context.Context, db orm.DB) error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

// BeforeUpdate - update updatedAt
func (m *BaseModel) BeforeUpdate(ctx context.Context, db orm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
