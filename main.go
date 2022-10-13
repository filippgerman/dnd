package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func DiceCheckNil(diceMaxNum *int) {

	if *diceMaxNum <= 0 {
		panic("Dice sides must be > 0")
	}
}

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
	DiceCheckNil(&diceMaxNum)
	c.JSON(200, gin.H{
		"Result": Get_DiceThrowResult(diceMaxNum),
	})
}
func main() {
	r := gin.Default()
	r.POST("/dice", PostDice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
