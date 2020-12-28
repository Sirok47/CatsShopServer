package main

import (
	"CatsShop/protocol"
	"CatsShopServer/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main(){
	db, err := pgx.Connect(context.Background(), os.Getenv("postgres://postgres:glazirovanniisirok@localhost:5432/CatsShop"))
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
