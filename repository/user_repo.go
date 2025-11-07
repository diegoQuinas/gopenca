
package repository

import (
	"database/sql"
	"github.com/diegoQuinas/gopenca/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, email FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err:= rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) EmailExists(email string) (bool, error) {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (r *UserRepo) Create(u models.User) (*models.User, error) {
	var created models.User
	err := r.DB.QueryRow(
		"INSERT INTO users (email) VALUES ($1) RETURNING id, email", u.Email,
	).Scan(&created.ID, &created.Email)
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (r *UserRepo) Update(id int, u models.User) (models.User, error) {
	_, err := r.DB.Exec(
		"UPDATE users SET email=%1 WHERE id=%2",
		u.Email, id,
	)
	u.ID = id
	return u, err
}

func (r *UserRepo) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id=%1", id)
	return err
}
