package service

import (
	"context"
	"github.com/Sirok47/CatsShopClient/protocol"
	"github.com/Sirok47/CatsShopServer/model"
)

func (c CatsShop) CreateCat(ctx context.Context, prm *protocol.Catparams) (*protocol.Errmsg, error) {
	err:=c.rps.CreateCat(ctx,model.CatParams{CatName: prm.CatName, CatAge: int(prm.CatAge), CatGender: prm.CatGender, CatBreed: prm.CatBreed})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) DeleteCat(ctx context.Context, prm *protocol.Catparams) (*protocol.Errmsg, error) {
	err:=c.rps.DeleteCat(ctx,int(prm.ID))
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) UpdateCat(ctx context.Context, prm *protocol.Catparams) (*protocol.Errmsg, error) {
	err:=c.rps.UpdateCat(ctx,&model.CatParams{ID: int(prm.ID), CatAge: int(prm.CatAge)})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) GetCat(ctx context.Context, prm *protocol.Catparams) (*protocol.Catparams, error) {
	cat,err:=c.rps.GetCat(ctx,int(prm.ID))
	if err!=nil{
		return &protocol.Catparams{Error: err.Error()},nil
	}
	return &protocol.Catparams{ID: int32(cat.ID),CatName: cat.CatName,CatAge: int32(cat.CatAge),CatGender: cat.CatGender,CatBreed: cat.CatBreed}, nil
}

func (c CatsShop) ListCats(ctx context.Context, _ *protocol.Catparams) (*protocol.Json, error) {
	json,err:=c.rps.ListCats(ctx)
	if err!=nil{
		return &protocol.Json{Error: err.Error()},nil
	}
	return &protocol.Json{Json: json}, nil
}

