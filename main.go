package main

import (
	pst "dnd/PostDice"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/dice", pst.PostDice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
