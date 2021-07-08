package repository

import (
	"gopkg.in/mgo.v2"
	"microservices/db"
)

const UsersCollection = "users"

type UsersRepository interface {}

type usersRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) *usersRepository {
	return &usersRepository{c: conn.DB().C(UsersCollection)}
}

func (u *usersRepository) Save() {}
