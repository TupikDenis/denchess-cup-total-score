package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Applications []Application
}

type Application struct {
	Name           string
	FullTournament FullTournament
}

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3307)/totalscore?charset=utf8mb4&parseTime=True&loc=Local")
	//dsn := "tcp://127.0.0.1:3306?database=tdenis5z_total&username=tdenis5z_total&password=tdl25@&read_timeout=10&write_timeout=20"
	//db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
