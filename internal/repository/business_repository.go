package repository

import (
	"context"
	"database/sql"
	"ocean-pos/internal/model"
)

type BusinessRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, business model.Business) (*model.Business, error)
}

type BusinessRepositoryImpl struct{}

func NewBusinessRepository() BusinessRepository {
	return &BusinessRepositoryImpl{}
}

func (repository *BusinessRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, business model.Business) (*model.Business, error) {
	SQL := "INSERT INTO business (owner_user_id, email, phone_number, name, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, business.OwnerUserId, business.Email, business.PhoneNumber, business.Name, business.CreatedAt, business.CreatedBy, business.UpdatedAt, business.UpdatedBy)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	business.Id = int(id)
	return &business, err
}
