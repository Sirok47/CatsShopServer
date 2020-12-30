package main

import (
	"context"
	"fmt"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/service"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
)

func main(){
	db, err := pgx.Connect(context.Background(), "postgres://postgres:glazirovanniisirok@localhost:5432/catsshop")
	if err!= nil{
		fmt.Print(err)
		return
	}
	srvobj:=service.NewService(db)
	srv:=grpc.NewServer()
	protocol.RegisterCatsShopServer(srv,srvobj)
	l,_:=net.Listen("tcp",":8080")
	srv.Serve(l)
}
