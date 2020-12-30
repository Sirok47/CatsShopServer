package service

import (
	"context"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
)

func (c CatsShop) CreateUser(ctx context.Context, prm *protocol.Userparams) (*protocol.Errmsg, error) {
	err:=c.rps.CreateUser(ctx,&model.UserParams{NickName: prm.NickName, Admin: prm.Admin, Password: prm.Password, Address: prm.Address})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) DeleteUser(ctx context.Context, prm *protocol.Userparams) (*protocol.Errmsg, error) {
	err:=c.rps.DeleteUser(ctx,model.UserParams{NickName: prm.NickName, Password: prm.Password})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) UpdateUser(ctx context.Context, prm *protocol.Userparams) (*protocol.Errmsg, error) {
	err:=c.rps.UpdateUser(ctx,model.UserParams{NickName: prm.NickName, Password: prm.Password, Address: prm.Address})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) Login(ctx context.Context, prm *protocol.Userparams) (*protocol.Jwtoken, error) {
	token,err:=c.rps.Login(ctx,prm.NickName,prm.Password)
	if err!=nil{
		return &protocol.Jwtoken{Error: err.Error()},nil
	}
	return &protocol.Jwtoken{Token: token}, nil
}

func (c CatsShop) ListUsers(ctx context.Context, _ *protocol.Userparams) (*protocol.Json, error) {
	json,err:=c.rps.ListUsers(ctx)
	if err!=nil{
		return &protocol.Json{Error: err.Error()},nil
	}
	return &protocol.Json{Json: json}, nil
}

