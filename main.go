package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func DiceCheckNil(diceRng *int) {

	if *diceRng <= 0 {
		panic("Dice sides must be > 0")
	}
}

func DiceThrow(diceRoll int) int {
	rand.Seed(time.Now().UnixNano())
	diceRoll = (rand.Intn(diceRoll))
	diceRoll += 1
	return diceRoll
}

func PostDice(c *gin.Context) {
	diceSideNumber := c.Query("DiceRange")
	diceRng, err := strconv.Atoi(diceSideNumber)
	if err != nil {
		c.JSON(400, "")
	}
	DiceCheckNil(&diceRng)
	c.JSON(200, gin.H{
		"Result": DiceThrow(diceRng),
	})
}
func main() {
	r := gin.Default()
	r.POST("/dice", PostDice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
