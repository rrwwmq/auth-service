package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rrwwmq/auth-service/internal/domain"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user domain.User) error {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)		
	`

	_, err := r.db.Exec(context.Background(), query, user.Email, user.PasswordHash)

	return err
}

func (r *UserRepo) GetByEmail(email string) (domain.User, error) {
	query := `
		SELECT * FROM users
		WHERE email = $1
	`

	var user domain.User
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
