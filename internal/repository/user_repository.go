package repository

import (
	"context"
	"database/sql"
	"ocean-pos/internal/model"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user model.User) (*model.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error)
	FindByPhoneNumber(ctx context.Context, tx *sql.Tx, phoneNumber string) (*model.User, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user model.User) (*model.User, error) {
	SQL := "INSERT INTO user (email, password, name, phone_number, created_at, created_by, updated_at, updated_by) VALUES (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, user.Name, user.PhoneNumber, user.CreatedAt, user.CreatedBy, user.UpdatedAt, user.UpdatedBy)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = int(id)

	return &user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error) {
	user := model.User{}
	SQL := "SELECT id,email,password,name,phone_number,is_email_verified,email_verified_at,is_phone_number_verified,phone_number_verified_at,deactivated_at,last_login,created_at,created_by,updated_at,updated_by FROM user WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.PhoneNumber,
			&user.IsEmailVerified,
			&user.EmailVerifiedAt,
			&user.IsPhoneNumberVerified,
			&user.PhoneNumberVerifiedAt,
			&user.DeactivatedAt,
			&user.LastLogin,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
		)
		return &user, nil
	} else {
		return nil, sql.ErrNoRows
	}
}

func (repository *UserRepositoryImpl) FindByPhoneNumber(ctx context.Context, tx *sql.Tx, phoneNumber string) (*model.User, error) {
	user := model.User{}
	SQL := "SELECT id,email,password,name,phone_number,is_email_verified,email_verified_at,is_phone_number_verified,phone_number_verified_at,deactivated_at,last_login,created_at,created_by,updated_at,updated_by FROM user WHERE phone_number = ?"
	rows, err := tx.QueryContext(ctx, SQL, phoneNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Email,
			&user.Name,
			&user.PhoneNumber,
			&user.IsEmailVerified,
			&user.EmailVerifiedAt,
			&user.IsPhoneNumberVerified,
			&user.PhoneNumberVerifiedAt,
			&user.DeactivatedAt,
			&user.LastLogin,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
		)
		return &user, nil
	} else {
		return nil, sql.ErrNoRows
	}
}
