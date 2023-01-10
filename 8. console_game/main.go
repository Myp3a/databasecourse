package main

import (
	"console_game/funcs"
	"fmt"
)

var hole, health, respect, weight int64 = 10, 100, 20, 30

func main() {
	rat := funcs.New(hole, health, respect, weight)
	day := 0
	for {
		day += 1
		fmt.Printf("++++++++++++ День %v ++++++++++++\n", day)
		rat.Info()
		rat.Day()
		fmt.Printf("++++++++++++ Ночь %v ++++++++++++\n", day)
		rat.Sleep()
		if rat.Hole == 0 || rat.Health == 0 || rat.Respect == 0 || rat.Weight == 0 {
			fmt.Println("Поражение.")
			break
		}
		if rat.Respect > 100 {
			fmt.Println("Победа!!!")
			break
		}
	}
}
