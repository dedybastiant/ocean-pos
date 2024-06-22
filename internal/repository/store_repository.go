package repository

import (
	"context"
	"database/sql"
	"ocean-pos/internal/model"
)

type StoreRepository interface {
	InsertStore(ctx context.Context, tx *sql.Tx, store model.Store) (*model.Store, error)
}

type StoreRepositoryImpl struct{}

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (repository *StoreRepositoryImpl) InsertStore(ctx context.Context, tx *sql.Tx, store model.Store) (*model.Store, error) {
	SQL := "INSERT INTO store (business_id, name, location, description, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?) "
	result, err := tx.ExecContext(ctx, SQL, store.BusinessId, store.Name, store.Location, store.Description, store.CreatedAt, store.CreatedBy, store.UpdatedAt, store.UpdatedBy)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	store.Id = int(id)

	return &store, nil
}
