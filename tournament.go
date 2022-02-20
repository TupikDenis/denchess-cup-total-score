package main

import "github.com/jinzhu/gorm"

type Tournament struct {
	gorm.Model
	JSON string `json:"json"`
}

type TransformedTournament struct {
	Id   uint   `json:"id"`
	JSON string `json:"json"`
}

type FullTournament struct {
	Score      int    `json:"score"`
	Username   string `json:"username"`
	Perfomance int    `json:"perfomance"`
}
