package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("app:app@tcp(127.0.0.1:3308)/app?parseTime=true&loc=Local"))
	db = db.Debug()

	a, _ := gormadapter.NewAdapterByDB(db)
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	e.LoadPolicy()
}
