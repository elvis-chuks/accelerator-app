package repository

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"inventory/domain"
	"time"
)

type userRepository struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func (u userRepository) Create(user domain.User) (*domain.User, error) {
	user.Id = domain.UUID{UUID: uuid.NewV4()}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.Db.Exec("INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)", user.Fullname, user.Email, user.Id, user.Password, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	row := u.Db.QueryRow("SELECT id, email, password, fullname, created_at, updated_at FROM users WHERE email = $1", email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Fullname, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u userRepository) GetById(id string) (*domain.User, error) {
	var user domain.User

	row := u.Db.QueryRow("SELECT id, email, password FROM users WHERE id = $1", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&user)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(db *sql.DB, logger *zap.Logger) domain.UserRepository {
	return &userRepository{
		Db:     db,
		Logger: logger,
	}
}
