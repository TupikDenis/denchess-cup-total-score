package main

import "github.com/jinzhu/gorm"

type Total struct {
	gorm.Model
	Username          string `json:"username"`
	TotalFirstPlace   int    `json:"total_first_place"`
	TotalSecondtPlace int    `json:"total_second_place"`
	TotalThirdPlace   int    `json:"total_third_place"`
	TotalTOP10        int    `json:"total_top_10"`
	TotalParticipate  int    `json:"total_paticipate"`
	TotalPoints       int    `json:"total_points"`
	TotalScore        int    `json:"total_score"`
}

type TransformedTotal struct {
	Id                uint   `json:"id"`
	Rank              int    `json:"rank"`
	Username          string `json:"username"`
	TotalFirstPlace   int    `json:"total_first_place"`
	TotalSecondtPlace int    `json:"total_second_place"`
	TotalThirdPlace   int    `json:"total_third_place"`
	TotalTOP10        int    `json:"total_top_10"`
	TotalParticipate  int    `json:"total_paticipate"`
	TotalPoints       int    `json:"total_points"`
	TotalScore        int    `json:"total_score"`
}
