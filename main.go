package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Введіть відстань в м:")
	var distance uint64
	fmt.Scan(&distance)

	fmt.Println("Введіть час в хв:")
	var t float64
	fmt.Scan(&t)

	velocity := float64(distance) / 1000 / (t / 60)
	fmt.Printf("Швидкість руху об'єкту %v км/год", velocity)
	time.Sleep(10 * time.Second)
}
