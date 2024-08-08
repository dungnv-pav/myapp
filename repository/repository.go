package repository

import (
	"database/sql"
	"myapp/domain/model"
)

type Repository struct {
}

// DeleteUser implements domain.IUserRepository.
func (r *Repository) DeleteUser(db *sql.DB, userID int) error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}

// GetUser implements domain.IUserRepository.
func (r *Repository) GetUser(db *sql.DB, userID int) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Other errors
	}
	return &user, nil
}

// GetUserList implements domain.IUserRepository.
func (r *Repository) GetUserList(db *sql.DB) ([]*model.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// RegisterUser implements domain.IUserRepository.
func (r *Repository) RegisterUser(db *sql.DB, user *model.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser implements domain.IUserRepository.
func (r *Repository) UpdateUser(db *sql.DB, userID int, user *model.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, userID)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository() *Repository {
	return &Repository{}
}
