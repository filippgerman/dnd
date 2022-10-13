package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)
func DiceCheckNil (diceRng *int){
	if diceRng<= 0 {

		})
		return
	}
}
func PostDice(c *gin.Context) {
	diceSideNumber := c.Query("DiceRange")
	diceRng, err := strconv.Atoi(diceSideNumber)
	if err != nil {

			return
		}
	DiceCheckNil(&diceRng)
	rand.Seed(time.Now().UnixNano())
	diceRng = (rand.Intn(diceRng))


	c.JSON(200, gin.H{
		"Result": diceRng,
	})
}
func main() {
	r := gin.Default()

	r.POST("/dice", PostDice)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
