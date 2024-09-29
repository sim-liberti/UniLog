package database

import (
	"database/sql"
	"teststack/types"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	db *sql.DB
}

func NewSqliteDB() (*SqliteDB, error) {
	const file string = "./sqlite.db"
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SqliteDB{db: db}, nil
}

func (db *SqliteDB) Init() error {
	_, err := db.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)

	return err
}

func (db *SqliteDB) GetUsers(query_params string) []*types.User {
	query := "SELECT * FROM users"

	if query_params != "" {
		query += " " + query_params
	}

	rows, err := db.db.Query(query)
	if err != nil {
		return nil
	}

	users := []*types.User{}
	for rows.Next() {
		user, err := types.LoadUser(rows)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

func (db *SqliteDB) GetUserByID(id int) *types.User {
	rows, err := db.db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil
	}

	user := &types.User{}
	if err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	); err == sql.ErrNoRows {
		return nil
	}

	return user
}
