package service

import (
	"context"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
)

func (c CatsShop) AddOperation(ctx context.Context, prm *protocol.Operationparams) (*protocol.Errmsg, error) {
	err:=c.rps.AddOperation(ctx,&model.OperationParams{NewOwnerNick: prm.NewOwnersNick, CatID: int(prm.CatID), CatName: prm.CatName, Status: prm.Status})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) EditOperation(ctx context.Context, prm *protocol.Operationparams) (*protocol.Errmsg, error) {
	err:=c.rps.EditOperation(ctx,&model.OperationParams{CatID: int(prm.CatID), Status: prm.Status})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) ListOperations(ctx context.Context, prm *protocol.Userparams) (*protocol.Json, error) {
	json,err:=c.rps.ListOperations(ctx,prm.NickName,prm.Admin)
	if err!=nil{
		return &protocol.Json{Error: err.Error()},nil
	}
	return &protocol.Json{Json: json}, nil
}

func (c CatsShop) GetOperation(ctx context.Context, prm *protocol.Operationparams) (*protocol.Operationparams, error) {
	opr,err:=c.rps.GetOperation(ctx,int(prm.CatID))
	if err!=nil{
		return &protocol.Operationparams{Error: err.Error()},nil
	}
	return &protocol.Operationparams{NewOwnersNick: opr.NewOwnerNick,CatID: int32(opr.CatID),CatName: opr.CatName,Date: opr.PurchaseDate,Status: opr.Status}, nil
}

