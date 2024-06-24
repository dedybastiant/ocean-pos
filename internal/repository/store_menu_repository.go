package repository

import (
	"context"
	"database/sql"
	"ocean-pos/internal/model"
)

type StoreMenuRepository interface {
	InsertStoreMenu(ctx context.Context, tx *sql.Tx, storeMenu model.StoreMenu) (*model.StoreMenu, error)
	FindStoreMenuByStoreAndMenuId(ctx context.Context, tx *sql.Tx, storeId int, menuId int) (*model.StoreMenu, error)
}

type StoreMenuRepositoryImpl struct{}

func NewStoreMenuRepository() StoreMenuRepository {
	return &StoreMenuRepositoryImpl{}
}

func (repository *StoreMenuRepositoryImpl) InsertStoreMenu(ctx context.Context, tx *sql.Tx, storeMenu model.StoreMenu) (*model.StoreMenu, error) {
	SQL := "INSERT INTO store_menu (store_id, menu_id, store_price, is_available, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, storeMenu.StoreId, storeMenu.MenuId, storeMenu.StorePrice, storeMenu.IsAvailable, storeMenu.CreatedAt, storeMenu.CreatedBy, storeMenu.UpdatedAt, storeMenu.UpdatedBy)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	storeMenu.Id = int(id)
	return &storeMenu, nil
}

func (repository *StoreMenuRepositoryImpl) FindStoreMenuByStoreAndMenuId(ctx context.Context, tx *sql.Tx, storeId int, menuId int) (*model.StoreMenu, error) {
	SQL := "SELECT id, store_id, menu_id, store_price, is_available, deactivated_at, created_at, created_by, updated_at, updated_by FROM store_menu WHERE store_id = ? and menu_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, storeId, menuId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	storeMenu := &model.StoreMenu{}
	if rows.Next() {
		rows.Scan(
			&storeMenu.Id,
			&storeMenu.StoreId,
			&storeMenu.MenuId,
			&storeMenu.StorePrice,
			&storeMenu.IsAvailable,
			&storeMenu.DeactivatedAt,
			&storeMenu.CreatedAt,
			&storeMenu.CreatedBy,
			&storeMenu.UpdatedAt,
			&storeMenu.UpdatedBy,
		)
		return storeMenu, nil
	} else {
		return nil, sql.ErrNoRows
	}
}
