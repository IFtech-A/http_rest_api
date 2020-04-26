package store

import (
	"github.com/IFtech-A/http_rest_api/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
