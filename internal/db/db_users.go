package db

import (
	"time"

	"github.com/utkarshkrsingh/goparty/internal/initializer"
)

type Users struct {
	ID        string    `db:"id" json:"id"`
	UserName  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

func (u *Users) CreateUser() error {
	dbQuery := `INSERT INTO users (username, email, password_hash)
    VALUES ($1, $2, $3)`
	_, err := initializer.DB.Exec(dbQuery, u.UserName, u.Email, u.Password)
	return err
}

func (u *Users) FindByMail() (*Users, error) {
	var user Users
	query := `SELECT * FROM users WHERE email = $1`
	if err := initializer.DB.Get(&user, query, u.Email); err != nil {
		return nil, err
	}
	return &user, nil
}
