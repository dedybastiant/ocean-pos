package repository

import (
	"context"
	"database/sql"
	"fmt"
	"ocean-pos/internal/model"
)

type MenuRepository interface {
	InsertMenu(ctx context.Context, tx *sql.Tx, menu model.Menu) (*model.Menu, error)
	FindMenuByName(ctx context.Context, tx *sql.Tx, menuName string) (*model.Menu, error)
}

type MenuRepositoryImpl struct{}

func NewMenuRepository() MenuRepository {
	return &MenuRepositoryImpl{}
}

func (repository *MenuRepositoryImpl) InsertMenu(ctx context.Context, tx *sql.Tx, menu model.Menu) (*model.Menu, error) {
	SQL := "INSERT INTO menu (category_id, name, default_price, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, menu.CategoryId, menu.Name, menu.DefaultPrice, menu.CreatedAt, menu.CreatedBy, menu.UpdatedAt, menu.UpdatedBy)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	menu.Id = int(id)
	return &menu, nil
}

func (repository *MenuRepositoryImpl) FindMenuByName(ctx context.Context, tx *sql.Tx, menuName string) (*model.Menu, error) {
	fmt.Println(menuName)
	SQL := "SELECT id, category_id, name, default_price, deactivated_at, created_at, created_by, updated_at, updated_by FROM menu WHERE name = ?"
	rows, err := tx.QueryContext(ctx, SQL, menuName)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows)

	menu := &model.Menu{}
	if rows.Next() {
		rows.Scan(
			&menu.Id,
			&menu.CategoryId,
			&menu.Name,
			&menu.DefaultPrice,
			&menu.DeactivatedAt,
			&menu.CreatedAt,
			&menu.CreatedBy,
			&menu.UpdatedAt,
			&menu.UpdatedBy,
		)
		return menu, nil
	} else {
		return nil, sql.ErrNoRows
	}
}
