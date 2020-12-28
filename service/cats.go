package service

import (
	"CatsShop/protocol"
	"CatsShopServer/model"
	"context"
)

func (c CatsShop) CreateCat(_ context.Context, prm *protocol.Catparams) (*protocol.Errmsg, error) {
	err:=c.rps.CreateCat(model.CatParams{CatName: prm.CatName, CatAge: int(prm.CatAge), CatGender: prm.CatGender, CatBreed: prm.CatBreed})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) DeleteCat(_ context.Context, prm *protocol.Catparams) (*protocol.Errmsg, error) {
	err:=c.rps.DeleteCat(int(prm.ID))
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) UpdateCat(_ context.Context, prm *protocol.Catparams) (*protocol.Errmsg, error) {
	err:=c.rps.UpdateCat(&model.CatParams{ID: int(prm.ID), CatAge: int(prm.CatAge)})
	if err!=nil{
		return &protocol.Errmsg{Error: err.Error()},nil
	}
	return nil, nil
}

func (c CatsShop) GetCat(_ context.Context, prm *protocol.Catparams) (*protocol.Catparams, error) {
	cat,err:=c.rps.GetCat(int(prm.ID))
	if err!=nil{
		return &protocol.Catparams{Error: err.Error()},nil
	}
	return &protocol.Catparams{ID: int32(cat.ID),CatName: cat.CatName,CatAge: int32(cat.CatAge),CatGender: cat.CatGender,CatBreed: cat.CatBreed}, nil
}

func (c CatsShop) ListCats(_ context.Context, _ *protocol.Catparams) (*protocol.Json, error) {
	json,err:=c.rps.ListCats()
	if err!=nil{
		return &protocol.Json{Error: err.Error()},nil
	}
	return &protocol.Json{Json: json}, nil
}

