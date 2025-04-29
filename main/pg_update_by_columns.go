package main

import (
	"xsky.com/ocpf/console/dao/usercenter"
	"xsky.com/ocpf/pkg/pg"
)

func main() {
	pg.Init(&pg.Config{
		Url:     "postgres://console_u:Xsky@123@127.0.0.1:5432/console?sslmode=disable",
		ShowSql: false,
		MaxIdle: 10,
		MaxOpen: 10,
	})
	user, err := usercenter.UserDao.Get(pg.Engine, 100000)
	if err != nil {
		panic(err)
	}
	if user == nil {
		panic("user not found")
	}
	user.Phone = "12345678901"
	user.Email = "linww@xsky.com"
	cols := make([]string, 0)
	cols = append(cols, "phone")
	_, err = usercenter.UserDao.UpdateByColumns(pg.Engine, user, cols...)
	if err != nil {
		panic(err)
	}
}
