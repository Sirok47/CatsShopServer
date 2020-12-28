package service

import (
	"CatsShop/protocol"
	"CatsShopServer/model"
	"context"
)

func (c CatsShop) CreateUser(_ context.Context, prm *protocol.Userparams) (*protocol.Errmsg, error) {
	err:=c.rps.CreateUser(&model.UserParams{NickName: prm.NickName, Admin: prm.Admin, Password: prm.Password, Address: prm.Address})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) DeleteUser(_ context.Context, prm *protocol.Userparams) (*protocol.Errmsg, error) {
	err:=c.rps.DeleteUser(model.UserParams{NickName: prm.NickName, Password: prm.Password})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) UpdateUser(_ context.Context, prm *protocol.Userparams) (*protocol.Errmsg, error) {
	err:=c.rps.UpdateUser(model.UserParams{NickName: prm.NickName, Password: prm.Password, Address: prm.Address})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) Login(_ context.Context, prm *protocol.Userparams) (*protocol.Jwtoken, error) {
	token,err:=c.rps.Login(prm.NickName,prm.Password)
	if err!=nil{
		return &protocol.Jwtoken{Error: err.Error()},nil
	}
	return &protocol.Jwtoken{Token: token}, nil
}

func (c CatsShop) ListUsers(_ context.Context, _ *protocol.Userparams) (*protocol.Json, error) {
	json,err:=c.rps.ListUsers()
	if err!=nil{
		return &protocol.Json{Error: err.Error()},nil
	}
	return &protocol.Json{Json: json}, nil
}

