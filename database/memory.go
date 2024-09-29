package database

import "teststack/types"

type MemoryDB struct{}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{}
}

func (db *MemoryDB) GetUsers(query_params string) *[]types.User {
	return &[]types.User{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "email@email.com",
			Password:  "password",
		},
		{
			ID:        2,
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "email@email.com",
			Password:  "password",
		},
	}
}

func (db *MemoryDB) GetUserByID(id int) *types.User {
	return &types.User{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "email@email.com",
		Password:  "password",
	}
}
