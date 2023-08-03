package repository

import "database/sql"

type user struct {
	db *sql.DB
}