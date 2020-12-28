package service

import (
	"CatsShop/protocol"
	"CatsShopServer/model"
	"context"
)

func (c CatsShop) AddOperation(_ context.Context, prm *protocol.Operationparams) (*protocol.Errmsg, error) {
	err:=c.rps.AddOperation(&model.OperationParams{NewOwnerNick: prm.NewOwnersNick, CatID: int(prm.CatID), CatName: prm.CatName, Status: prm.Status})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) EditOperation(_ context.Context, prm *protocol.Operationparams) (*protocol.Errmsg, error) {
	err:=c.rps.EditOperation(&model.OperationParams{CatID: int(prm.CatID), Status: prm.Status})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) ListOperations(_ context.Context, prm *protocol.Userparams) (*protocol.Json, error) {
	json,err:=c.rps.ListOperations(prm.NickName,prm.Admin)
	if err!=nil{
		return &protocol.Json{Error: err.Error()},nil
	}
	return &protocol.Json{Json: json}, nil
}

func (c CatsShop) GetOperation(_ context.Context, prm *protocol.Operationparams) (*protocol.Operationparams, error) {
	opr,err:=c.rps.GetOperation(int(prm.CatID))
	if err!=nil{
		return &protocol.Operationparams{Error: err.Error()},nil
	}
	return &protocol.Operationparams{NewOwnersNick: opr.NewOwnerNick,CatID: int32(opr.CatID),CatName: opr.CatName,Date: opr.PurchaseDate,Status: opr.Status}, nil
}

