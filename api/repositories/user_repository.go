package repositories

import (
	"database/sql"
	"go-app/internal/models"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	query := "SELECT id, username, email, first_name, last_name, date_of_birth, created_at, last_login FROM users WHERE id = $1"
	user := &models.User{}

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.CreatedAt,
		&user.LastLogin,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, username, email, first_name, last_name, date_of_birth, created_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.CreatedAt); err != nil {
			log.Printf("Error scanning user rows: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after scanning users: %v", err)
		return nil, err
	}

	return users, nil
}


