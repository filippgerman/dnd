package main

import (
	chk "dnd/DiceChecking"
	dth "dnd/GetThrowResult"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostDice(c *gin.Context) {
	diceSideNumber := c.Query("DiceRange")
	diceMaxNum, err := strconv.Atoi(diceSideNumber)
	if err != nil {
		c.JSON(400, "")
	}
	chk.DiceCheckNil(&diceMaxNum)
	c.JSON(200, gin.H{
		"Result": dth.Get_DiceThrowResult(diceMaxNum),
	})
}
func main() {
	r := gin.Default()
	r.POST("/dice", PostDice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
