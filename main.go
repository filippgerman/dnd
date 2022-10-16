package main

import (
	chk "github.com/filippgerman/dnd/blob/0c3a82ba0925aed9f2f47afc307b94f46713eee3/DiceChecking"
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
