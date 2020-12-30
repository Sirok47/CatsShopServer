package service

import (
	"github.com/Sirok47/CatsShopServer/repository"
	"github.com/jackc/pgx/v4"
)

type CatsShop struct {
	rps *repository.CatsShop
}

func NewService(db *pgx.Conn) *CatsShop {
	return &CatsShop{repository.NewCatsShop(db)}
}
