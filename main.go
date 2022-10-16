package main

import (
	chk "github.com/filippgerman/dnd/blob/e8e01b736093616d9bae87ab43058e0b492c8f0a/DiceChecking/DiceCheck"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func Get_DiceThrowResult(diceResult int) int {
	rand.Seed(time.Now().UnixNano())
	diceResult = (rand.Intn(diceResult))
	diceResult += 1
	return diceResult
}

func PostDice(c *gin.Context) {
	diceSideNumber := c.Query("DiceRange")
	diceMaxNum, err := strconv.Atoi(diceSideNumber)
	if err != nil {
		c.JSON(400, "")
	}
	chk.DiceCheckNil(&diceMaxNum)
	c.JSON(200, gin.H{
		"Result": Get_DiceThrowResult(diceMaxNum),
	})
}
func main() {
	r := gin.Default()
	r.POST("/dice", PostDice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
