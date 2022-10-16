package DiceChecking

func DiceCheckNil(diceMaxNum *int) {

	if *diceMaxNum <= 0 {
		panic("Dice sides must be > 0")
	}
}
