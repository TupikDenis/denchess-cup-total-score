package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func funcHandle() {
	router := gin.Default()
	router.LoadHTMLGlob("*")

	router.GET("/", home)
	router.POST("/add", addRound)

	router.Run(":8080")
}

func home(c *gin.Context) {
	total := getAllPlayers()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Зачёт",
		"total": total,
	})
}

func addRound(c *gin.Context) {
	id := c.PostForm("id")
	filenameJSON := id + ".json"

	data := &Data{
		URL:      "https://lichess.org/api/tournament/" + id + "/results",
		filename: filenameJSON,
	}

	data.download()
	rewriteJSON(filenameJSON)

	tournament := tournamentArray(filenameJSON)
	changeTotal(tournament)

	c.Redirect(http.StatusFound, "/")
}
