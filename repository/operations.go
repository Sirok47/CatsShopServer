package repository

import (
	"context"
	"encoding/json"
	"github.com/Sirok47/CatsShopServer/model"
	"github.com/jackc/pgx/v4"
	"time"
)

func (r CatsShop) AddOperation(ctx context.Context,o *model.OperationParams) error {
	_, err := r.db.Exec(ctx,"insert into operations (newOwnersNick,CatID,PurchaseDate,Status) values ($1,$2,$3,$4)", o.NewOwnerNick, o.CatID, time.Now(), o.Status)
	return err
}

func (r CatsShop) EditOperation(ctx context.Context,o *model.OperationParams) error {
	_, err := r.db.Exec(ctx,"update operations set Status = $1 where CatID = $2",o.Status,o.CatID)
	return err
}

func (r CatsShop) GetOperation(ctx context.Context,id int) (*model.OperationParams,error) {
	o := &model.OperationParams{CatID: id}
	result, err := r.db.Query(ctx,"select * from operations where CatID = $1", id)
	defer result.Close()
	if err != nil {
		return nil, err
	}
	for result.Next() {
		err = result.Scan(&o.NewOwnerNick,&o.CatID,&o.PurchaseDate,&o.Status)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func (r CatsShop) ListOperations(ctx context.Context,nick string, admin bool) (string,error) {
	var(
		o = model.OperationParams{}
		result pgx.Rows
		err error
	)
if admin==true {
	result, err = r.db.Query(ctx,"select CatID,Status from operations")
} else {
	result, err = r.db.Query(ctx,"select CatID,Status from operations where NewOwnersNick = $1", nick)
	}
defer result.Close()
if err != nil {
	return "", err
}
m:=make(map[int]string,10)
for result.Next() {
	err = result.Scan(&o.CatID,&o.Status)
if err != nil {
return "", err
}
m[o.CatID]=o.Status
}
list,err:=json.Marshal(m)
if err != nil {
return "", err
}
return string(list), nil
}