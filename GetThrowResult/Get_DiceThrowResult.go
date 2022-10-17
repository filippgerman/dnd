package GetThrowResult

import (
	"math/rand"
	"time"
)

func Get_DiceThrowResult(diceResult int) int {
	rand.Seed(time.Now().UnixNano())
	diceResult = (rand.Intn(diceResult))
	diceResult += 1
	return diceResult
}
