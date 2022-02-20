package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func tournamentArray(filename string) []FullTournament {
	var jsonContent string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		jsonContent += fileScanner.Text()
	}

	//log.Println(jsonContent)
	file.Close()

	tournament := []FullTournament{}

	err = json.Unmarshal([]byte(jsonContent), &tournament)
	if err != nil {
		fmt.Println(err)
	}

	return tournament
}

func changeTotal(tournament []FullTournament) {
	players := getAllPlayers()

	for i := 0; i < len(tournament); i++ {
		flag := false
		points := 0
		score := tournament[i].Score
		username := tournament[i].Username
		perfomance := tournament[i].Perfomance

		switch i {
		case 0:
			points = 25
		case 1:
			points = 19
		case 2:
			points = 14
		case 3:
			points = 10
		case 4:
			points = 7
		case 5:
			points = 5
		case 6:
			points = 4
		case 7:
			points = 3
		case 8:
			points = 2
		case 9:
			points = 1
		}

		if players != nil {
			for k := 0; k < len(players); k++ {
				if username == players[k].Username {
					flag = true
					break
				}
			}
		}

		if flag {
			changePlayer(username, i, perfomance, points, score)
		} else {
			addPlayer(username, score, i, perfomance, points)
		}
	}
}

func getAllPlayers() []TransformedTotal {
	var players []Total
	var _players []TransformedTotal

	db := Database()
	db.Order("total_points desc").Order("total_first_place desc").Order("total_secondt_place desc").Order("total_third_place desc").Order("total_top10 desc").Order("total_participate desc").Order("total_score desc").Find(&players)

	if len(players) <= 0 {
		return nil
	}

	//transforms the todos for building a good response
	i := 1
	for _, item := range players {
		_players = append(
			_players, TransformedTotal{
				Id:                item.ID,
				Rank:              i,
				Username:          item.Username,
				TotalFirstPlace:   item.TotalFirstPlace,
				TotalSecondtPlace: item.TotalSecondtPlace,
				TotalThirdPlace:   item.TotalThirdPlace,
				TotalTOP10:        item.TotalTOP10,
				TotalParticipate:  item.TotalParticipate,
				TotalPoints:       item.TotalPoints,
				TotalScore:        item.TotalScore,
			})

		i++
	}

	return _players
}

func getPlayer(username string) Total {
	var players []Total
	var _players []TransformedTotal

	db := Database()
	db.Find(&players)

	//transforms the todos for building a good response
	for _, item := range players {
		if item.Username == username {
			_players = append(
				_players, TransformedTotal{
					Id:                item.ID,
					Username:          item.Username,
					TotalFirstPlace:   item.TotalFirstPlace,
					TotalSecondtPlace: item.TotalSecondtPlace,
					TotalThirdPlace:   item.TotalThirdPlace,
					TotalTOP10:        item.TotalTOP10,
					TotalParticipate:  item.TotalParticipate,
					TotalPoints:       item.TotalPoints,
					TotalScore:        item.TotalScore,
				})
		}

	}

	_player := _players[len(_players)-1]

	player := Total{
		Model:             gorm.Model{},
		Username:          _player.Username,
		TotalFirstPlace:   _player.TotalFirstPlace,
		TotalSecondtPlace: _player.TotalSecondtPlace,
		TotalThirdPlace:   _player.TotalThirdPlace,
		TotalTOP10:        _player.TotalTOP10,
		TotalParticipate:  _player.TotalParticipate,
		TotalPoints:       _player.TotalPoints,
		TotalScore:        _player.TotalScore,
	}

	return player
}

func changePlayer(username string, index int, perfomance int, points int, score int) {
	player := getPlayer(username)

	db := Database()

	db.First(&player, player.ID)

	db.Model(&Total{}).Where("username = ?", username).Update("total_points", player.TotalPoints+points)
	db.Model(&Total{}).Where("username = ?", username).Update("total_participate", player.TotalParticipate+1)
	db.Model(&Total{}).Where("username = ?", username).Update("total_score", player.TotalScore+score)

	switch index {
	case 0:
		db.Model(&Total{}).Where("username = ?", username).Update("total_first_place", player.TotalFirstPlace+1)
	case 1:
		db.Model(&Total{}).Where("username = ?", username).Update("total_second_place", player.TotalSecondtPlace+1)
	case 2:
		db.Model(&Total{}).Where("username = ?", username).Update("total_third_place", player.TotalThirdPlace+1)
	}

	if index < 10 {
		db.Model(&Total{}).Where("username = ?", username).Update("total_top10", player.TotalTOP10+1)
	}
	//db.Close()
}

func addPlayer(username string, score int, index int, perfomance int, points int) {
	var firstPlace int
	var secondPlace int
	var thirdPlace int
	var top10 int

	switch index {
	case 0:
		firstPlace = 1
	case 1:
		secondPlace = 1
	case 2:
		thirdPlace = 1
	}

	if index < 10 {
		top10 = 1
	}

	total := Total{
		Model:             gorm.Model{},
		Username:          username,
		TotalFirstPlace:   firstPlace,
		TotalSecondtPlace: secondPlace,
		TotalThirdPlace:   thirdPlace,
		TotalTOP10:        top10,
		TotalParticipate:  1,
		TotalPoints:       points,
		TotalScore:        score,
	}

	db := Database()
	db.AutoMigrate(&Total{})
	db.Create(&total)
	//db.Close()
}
