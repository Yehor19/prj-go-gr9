package main

import (
	"fmt"
	"math/rand"
	"prj-go/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 10
	pointsPerQuestion = 5
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаю в моїй грі!")
	time.Sleep(1 * time.Second)

	var users []domain.User

	for {
		menu()

		point := ""
		fmt.Scan(&point)

		switch point {
		case "1":
			u := play()
			users = append(users, u)
		case "2":
			for i, user := range users {
				fmt.Printf(
					"i: %v, id: %v, name: %s, time: %v\n",
					i, user.Id, user.Name, user.Time,
				)
			}
		case "3":
			return
		default:
			fmt.Println("Зробіть коректний вибір")
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	start := time.Now()
	for myPoints < totalPoints {
		x := rand.Intn(100)
		y := rand.Intn(100)

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				fmt.Printf(
					"Чудово! Ти набрав %v балів!\nЗалишилось набрати %v\n",
					myPoints, totalPoints-myPoints,
				)
			} else {
				fmt.Println("Пощастить наступного разу!")
			}
		}
	}
	end := time.Now()
	timeSpent := end.Sub(start)

	fmt.Printf("Вітаю! Ти пройшов гру за %v!\n", timeSpent)
	fmt.Println("Введіть своє ім'я")

	name := ""
	fmt.Scan(&name)

	// var user domain.User
	// user.Id = id
	// user.Name = name
	// user.Time = timeSpent

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

}
