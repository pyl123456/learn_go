package service

import (
	"log"
	"database/sql"

	"github.com/google/wire"
)

// UserRepository
type UserRepository interface {
	AddUser()
}

// userRepo UserRepository
type userRepo struct {
	*sql.DB
}

func (u *userRepo) AddUser() {
	log.Println("add user")
}

func NewUserRepo(db *sql.DB) *userRepo{
	return &userRepo{}
}

var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(UserRepository), new(*userRepo)))