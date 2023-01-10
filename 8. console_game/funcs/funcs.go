package funcs

import (
	"fmt"
	"math/rand"
)

// Неплохо бы добавить вывод изменения характеристик...
type rat struct {
	Hole, Health, Respect, Weight int64
}

func New(hole, health, respect, weight int64) *rat {
	return &rat{
		hole,
		health,
		respect,
		weight}
}

func (r *rat) Update(hole, health, respect, weight int64) {
	r.Hole += hole
	r.Health += health
	r.Respect += respect
	r.Weight += weight
	fmt.Print("Изменения:")
	if hole != 0 {
		fmt.Printf(" длина норы: %v", hole)
	}
	if health != 0 {
		fmt.Printf(" здоровье: %v", health)
	}
	if respect != 0 {
		fmt.Printf(" уважение: %v", respect)
	}
	if weight != 0 {
		fmt.Printf(" вес: %v", weight)
	}
	fmt.Println()
}

func (r *rat) Info() {
	// Господи, кто придумал отсутствие printf с новой строкой...
	fmt.Println("Крыса:")
	fmt.Printf("Длина норы \t- %v\n", r.Hole)
	fmt.Printf("Здоровье \t- %v\n", r.Health)
	fmt.Printf("Уважение \t- %v\n", r.Respect)
	fmt.Printf("Вес \t- %v\n", r.Weight)
}

func (r *rat) Day() {
	fmt.Println("Действия:")
	fmt.Println("1 - копать нору")
	fmt.Println("2 - жрат травку")
	fmt.Println("3 - драться")
	fmt.Println("4 - спат")
	fmt.Print("Выбор: ")
	var act int64
	fmt.Scan(&act)
	switch act {
	case 1:
		r.dig()
	case 2:
		r.eat()
	case 3:
		r.fight()
	case 4:
		r.Sleep()
	default:
		// Надо сделать в цикле, иначе второе значение не проверяется
		fmt.Println("Неверный выбор!")
		fmt.Print("Выбор: ")
		fmt.Scan(&act)
	}
}

func (r *rat) dig() {

	fmt.Println("Как копаем?")
	fmt.Println("1 - интенсивно")
	fmt.Println("2 - лениво")
	fmt.Print("Выбор: ")
	var act int64
	fmt.Scan(&act)
	switch act {
	case 1:
		r.Update(5, -30, 0, 0)
	case 2:
		r.Update(2, -10, 0, 0)
	default:
		fmt.Println("Неверный выбор!")
		fmt.Print("Выбор: ")
		fmt.Scan(&act)
	}
}

func (r *rat) eat() {
	fmt.Println("Что кушац?")
	fmt.Println("1 - пожухлую траву")
	fmt.Println("2 - свежую траву")
	fmt.Print("Выбор: ")
	var act int64
	fmt.Scan(&act)
	switch act {
	case 1:
		r.Update(0, 10, 0, 15)
	case 2:
		if r.Respect >= 30 {
			r.Update(0, 30, 0, 30)
		} else {
			r.Update(0, -30, 0, 0)
		}
	}
}

func (r *rat) fight() {
	var enemyWeight int64
	var winP float64
	fmt.Println("Кого отпинаем?")
	fmt.Println("1 - Слабого")
	fmt.Println("2 - Среднего")
	fmt.Println("3 - Сильного")
	fmt.Print("Выбор: ")
	var act int64
	fmt.Scan(&act)
	switch act {
	case 1:
		enemyWeight = 30
	case 2:
		enemyWeight = 50
	case 3:
		enemyWeight = 70
	default:
		fmt.Println("Неверный выбор!")
		fmt.Print("Выбор: ")
		fmt.Scan(&act)
	}
	winP = float64(r.Weight) / float64(r.Weight+enemyWeight)
	randPoint := rand.Float64()
	if 0 < randPoint && randPoint <= winP {
		fmt.Println("Поражение!")
	} else {
		fmt.Println("Победа!")
		if r.Weight < enemyWeight {
			r.Update(0, -40, 40, 0)
		} else if r.Weight == enemyWeight {
			r.Update(0, -20, 20, 0)
		} else {
			r.Update(0, -10, 10, 0)
		}
	}

}

func (r *rat) Sleep() {
	r.Update(-2, 20, -2, -5)
}
