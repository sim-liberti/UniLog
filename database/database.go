package database

import "teststack/types"

type Database interface {
	// CRUD operations for all types
	GetUsers(string) []*types.User
	GetUserByID(int) *types.User
	// CreateUser()
	// UpdateUser()
	// DeleteUser()
}
