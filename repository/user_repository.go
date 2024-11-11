package repository

import (
	"database/sql"
	"test-ordent/model"
)

type UserRepository interface {
	GetById(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
	GetByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) GetById(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, full_name, email, password, phone_number, address, role from mst_users WHERE id=$1`, id).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Address,
		&user.Role,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) Create(payload model.User) (model.User, error) {
	// fmt.Println(payload)
	var user model.User
	err := u.db.QueryRow(`INSERT INTO mst_users (id, full_name, email, password, phone_number, address, role, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id, full_name, email, password, phone_number, address, role, created_at, updated_at`,
		payload.Id,
		payload.FullName,
		payload.Email,
		payload.Password,
		payload.PhoneNumber,
		payload.Address,
		payload.Role,
		payload.CreatedAt,
		payload.UpdatedAt,
	).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Address,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, full_name, email, password, phone_number, address, role from mst_users WHERE email=$1`, email).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Address,
		&user.Role,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
