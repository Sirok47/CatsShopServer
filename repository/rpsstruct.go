package repository

import (
	"github.com/jackc/pgx/v4"
)

type CatsShop struct {
	db *pgx.Conn
}

func NewCatsShop(db *pgx.Conn) *CatsShop {
	return &CatsShop{db: db}
}
