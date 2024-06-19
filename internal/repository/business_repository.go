package repository

import (
	"context"
	"database/sql"
	"ocean-pos/internal/model"
)

type BusinessRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, business model.Business) (*model.Business, error)
	FindBusinessById(ctx context.Context, tx *sql.Tx, businessId int) (*model.Business, error)
}

type BusinessRepositoryImpl struct{}

func NewBusinessRepository() BusinessRepository {
	return &BusinessRepositoryImpl{}
}

func (repository *BusinessRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, business model.Business) (*model.Business, error) {
	SQL := "INSERT INTO business (owner_user_id, email, phone_number, name, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, business.OwnerUserId, business.Email, business.PhoneNumber, business.Name, business.CreatedAt, business.CreatedBy, business.UpdatedAt, business.UpdatedBy)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	business.Id = int(id)
	return &business, err
}

func (repository *BusinessRepositoryImpl) FindBusinessById(ctx context.Context, tx *sql.Tx, businessId int) (*model.Business, error) {
	SQL := "SELECT id, owner_user_id, email, phone_number, name, verified_at, deactivated_at, created_at, created_by, updated_at, updated_by FROM business WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, businessId)
	if err != nil {
		return nil, err
	}

	business := &model.Business{}
	if rows.Next() {
		rows.Scan(
			&business.Id,
			&business.OwnerUserId,
			&business.Email,
			&business.PhoneNumber,
			&business.Name,
			&business.VerifiedAt,
			&business.DeactivatedAt,
			&business.CreatedAt,
			&business.CreatedBy,
			&business.UpdatedAt,
			&business.UpdatedBy,
		)
		return business, nil
	} else {
		return nil, sql.ErrNoRows
	}
}
