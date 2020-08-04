package domain

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type AccountService interface {
	RegisterNew(ctx context.Context, p Profile, password string) (uint64, error)
	Login(ctx context.Context, userName string, password string) error
	Logout(ctx context.Context) error
}

type AccountServiceImplementation struct {
	db *sqlx.DB
}

const insertAccountQuery = `INSERT INTO Profile (Email, PasswordHash, FirstName, LastName, Age, Gender, City, Hobby) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

func (s AccountServiceImplementation) RegisterNew(ctx context.Context, p Profile, password string) (uint64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = s.db.ExecContext(ctx, insertAccountQuery, p.Email, password, p.FirstName, p.LastName, p.Age, p.Gender, p.City, p.Hobby)
	row := s.db.QueryRow("SELECT LAST_INSERT_ID()")
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return id, err
}

func (s AccountServiceImplementation) Login(ctx context.Context, userName string, password string) error {
	panic("implement me")
}


func (s AccountServiceImplementation) Logout(ctx context.Context) error {
	panic("implement me")
}

func NewAccountService(db *sqlx.DB) AccountService {
	return AccountServiceImplementation{
		db,
	}
}
