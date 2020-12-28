package repository

import (
	"CatsShopServer/model"
	"context"
	"encoding/json"
)

func (r CatsShop) CreateCat(g model.CatParams) error {
		_, err := r.db.Exec(context.Background(),"insert into cats (CatName,CatAge,CatGender,CatBreed) values ($1,$2,$3,$4)",g.CatName, g.CatAge, g.CatGender, g.CatBreed)
		return err
}

func (r CatsShop) GetCat(id int) (*model.CatParams,error) {
	c := &model.CatParams{ID: id}
	result, err := r.db.Query(context.Background(),"select * from cats where ID = $1", c.ID)
	defer result.Close()
	if err != nil {
		return nil, err
	}
	for result.Next() {
		err = result.Scan(&c.CatName, &c.ID, &c.CatAge, &c.CatGender, &c.CatBreed)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (r CatsShop) UpdateCat(c *model.CatParams) error {
	_, err := r.db.Exec(context.Background(),"update cats set CatAge = $1 where ID = $2",c.CatAge,c.ID)
	return err
}

func (r CatsShop) DeleteCat(id int) error {
	_, err := r.db.Exec(context.Background(),"delete from cats where ID = $1", id)
	return err
}

func (r CatsShop) ListCats() (string,error) {
	c := &model.CatParams{}
	m := map[int]string{}
	result, err := r.db.Query(context.Background(),"select ID,CatName from cats")
	defer result.Close()
	if err != nil {
		return "", err
	}
	for result.Next() {
		err = result.Scan(&c.ID, &c.CatName)
		if err != nil {
			return "", err
		}
		m[c.ID]=c.CatName
	}
	list,err:=json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(list), nil
}