package repository

import (
	"context"
	"database/sql"
	"ocean-pos/internal/model"
)

type CategoryRespository interface {
	InsertCategory(ctx context.Context, tx *sql.Tx, category model.Category) (*model.Category, error)
	FindCategoryByName(ctx context.Context, tx *sql.Tx, categoryName string) (*model.Category, error)
}

type CategoryRespositoryImpl struct{}

func NewCategoryRepository() CategoryRespository {
	return &CategoryRespositoryImpl{}
}

func (repository *CategoryRespositoryImpl) InsertCategory(ctx context.Context, tx *sql.Tx, category model.Category) (*model.Category, error) {
	SQL := "INSERT INTO category (business_id, name, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, category.BusinessId, category.Name, category.CreatedAt, category.CreatedBy, category.UpdatedAt, category.UpdatedBy)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	category.Id = int(id)
	return &category, nil
}

func (repository *CategoryRespositoryImpl) FindCategoryByName(ctx context.Context, tx *sql.Tx, categoryName string) (*model.Category, error) {
	SQL := "SELECT id, business_id, name, deactivated_at, created_at, created_by, updated_at, updated_by  FROM category c WHERE name = ? "
	rows, err := tx.QueryContext(ctx, SQL, categoryName)
	if err != nil {
		return nil, err
	}

	category := &model.Category{}

	if rows.Next() {
		rows.Scan(
			&category.Id,
			&category.BusinessId,
			&category.Name,
			&category.DeactivatedAt,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.UpdatedAt,
			&category.UpdatedBy,
		)
		return category, nil
	} else {
		return nil, sql.ErrNoRows
	}
}
