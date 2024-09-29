package database

import "teststack/types"

type PostgresDB struct{}

func (db *PostgresDB) GetUser(id int) *types.User {
	return &types.User{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "email@email.com",
		Password:  "password",
	}
}
