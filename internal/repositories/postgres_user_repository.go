package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/utkarshkrsingh/goparty/internal/db"
)

type PostgresUserRepository struct {
	DB *sqlx.DB
}

func (r *PostgresUserRepository) FindByEmail(email string) (*db.Users, error) {
	var user db.Users
	query := `SELECT * FROM users WHERE email = $1`
	if err := r.DB.Get(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}
