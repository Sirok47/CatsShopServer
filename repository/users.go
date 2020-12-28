package repository

import (
	"CatsShopServer/model"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx"
	"github.com/labstack/echo"
	"time"
)

func (r CatsShop) CreateUser(u *model.UserParams) error {
	_, err := r.db.Exec(context.Background(),"insert into users (NickName,isAdmin,Password,Address) values ($1,$2,$3,$4)",u.NickName,u.Admin,u.Password,u.Address)
	return err
}

func (r CatsShop) Login(nick string,password string) (string,error) {
	u := model.UserParams{}
	res, err := r.db.Query(context.Background(),"select * from users where NickName = $1", nick)
	defer res.Close()
	if err != nil {
		return "error", err
	}
	for res.Next() {
		err = res.Scan(&u.NickName,&u.Admin,&u.Password,&u.Address)
		if err != nil {
			return "error", err
		}
	}
	if password==u.Password{
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["Nick"] = u.NickName
		claims["Admin"] = u.Admin
		claims["Address"] = u.Address
		claims["exp"] = time.Now().Add(time.Hour).Unix()
		t, err := token.SignedString([]byte("sirok"))
		if err != nil {
			return "error",err
		}
		return t,nil
	}
	return "",echo.ErrUnauthorized
}

func authorization(pass,nick string, db *pgx.Conn) bool {
	var truePass string
	res, err := db.Query(context.Background(),"select Password from users where NickName = $1",nick)
	defer res.Close()
	if err != nil {
		return false
	}
	for res.Next() {
		err = res.Scan(&truePass)
		if err != nil {
			return false
		}
	}
	if pass==truePass{
		return true
	}
	return false
}

func (r CatsShop) UpdateUser(u model.UserParams) error {
	if authorization(u.Password,u.NickName,r.db)==true{
	_, err := r.db.Exec(context.Background(),"update users set Address = $1 where NickName = $2",u.Address,u.NickName)
	return err
}
return echo.ErrUnauthorized
}

func (r CatsShop) DeleteUser(u model.UserParams) error {
	if authorization(u.Password,u.NickName,r.db)==true{
		_, err := r.db.Exec(context.Background(),"delete from users where NickName = $1", u.NickName)
		return err
	}
	return echo.ErrUnauthorized
}

func (r CatsShop) ListUsers() (string,error) {
	u := model.UserParams{}
	m := map[string]bool{}
	result, err := r.db.Query(context.Background(),"select NickName,isadmin from users")
	defer result.Close()
	if err != nil {
		return "", err
	}
	for result.Next() {
		err = result.Scan(&u.NickName, &u.Admin)
		if err != nil {
			return "", err
		}
		m[u.NickName]=u.Admin
	}
	list,err:=json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(list), nil
}