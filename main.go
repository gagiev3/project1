package main

import (
	"fmt"
	"math/rand"
	"project_1/domain"
	"sort"
	"strconv"
	"time"
)

type ByTime []domain.User

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time < a[j].Time }

var id uint64 = 1

const (
	pointsPerQuestion = 5
	requiredPoints    = 10
)

func menu() {
	fmt.Println("1.Почати гру")
	fmt.Println("2.Переглянути рейтинг")
	fmt.Println("3.Завершити гру")

}

func play() domain.User {
	fmt.Println("Вітаємо у найгіршій грі всесвіту!!!! :(")
	time.Sleep(3 * time.Second)

	fmt.Println("Гра починається через...")
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	points := requiredPoints

	now := time.Now()
	for points > 0 {
		var res, x, y int

		operation := rand.Intn(3)

		switch operation {
		case 0: // Додавання
			x, y = rand.Intn(100), rand.Intn(100)
			fmt.Printf("%v + %v = ", x, y)
			res = x + y
		case 1: // Віднімання
			x, y = rand.Intn(100), rand.Intn(100)
			fmt.Printf("%v - %v = ", x, y)
			res = x - y
		case 2: // Множення
			x, y = rand.Intn(10)+1, rand.Intn(10)+1 // Генеруємо числа від 1 до 10
			fmt.Printf("%v * %v = ", x, y)
			res = x * y
		}

		var answer string
		fmt.Scan(&answer)

		ansInt, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			if res == ansInt {
				points -= pointsPerQuestion
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно :( Залишилось набрати %v \n", points)
				fmt.Printf("У вас зараз %v балів. \n", myPoints)
			} else {
				fmt.Println("Ха-Ха-Ха! Неправильно!")
			}
		}

	}
	then := time.Now()
	timeSpent := then.Sub(now).Seconds()

	fmt.Printf("Чьйорт! Ти впорався за %v", timeSpent)
	fmt.Println("Введіть ім'я гравця:")

	name := ""
	fmt.Scan(&name)

	if name == "" {
		fmt.Println("Некоректне ім'я, дані не будуть збережені")
		return domain.User{}
	}

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func main() {
	menu()
	fmt.Println("Вітаємо у найгіршій грі всесвіту!!!! :(")
	time.Sleep(2 * time.Second)

	var users []domain.User

	for {
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			if u.Id != 0 {
				users = append(users, u)
			}
		case "2":
			sort.Sort(ByTime(users)) // Сортуємо гравців за часом
			for _, u := range users {
				fmt.Printf("id: %v, name: %s, time: %v\n",
					u.Id,
					u.Name,
					u.Time,
				)
			}
		case "3":
			return
		default:
			fmt.Println("Зробіть правильний вибір")
		}
	}
}
