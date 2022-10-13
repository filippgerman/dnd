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

func DiceThrow(diceRng *int) {
	rand.Seed(time.Now().UnixNano())
	*diceRng = (rand.Intn(*diceRng))
	*diceRng += 1
}

func PostDice(c *gin.Context) {
	diceSideNumber := c.Query("DiceRange")
	diceRng, _ := strconv.Atoi(diceSideNumber)
	DiceCheckNil(&diceRng)
	DiceThrow(&diceRng)
	c.JSON(200, gin.H{
		"Result": diceRng,
	})
}
func main() {
	r := gin.Default()
	r.POST("/dice", PostDice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
