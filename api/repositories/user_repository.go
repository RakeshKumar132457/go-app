package repositories

import (
	"database/sql"
	"go-app/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, username, email, first_name, last_name, date_of_birth, created_at, last_login FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.CreatedAt, &user.LastLogin)

	if err != nil {
		return nil, err
	}
	return user, nil
}
