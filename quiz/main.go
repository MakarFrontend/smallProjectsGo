package main

import (
	"fmt"
)

func main() {
	fmt.Println("Добро пожаловать!")
	var quantity int
	for {
		for {
			fmt.Print("Сколько вопросов в Викторине?::: ")
			_, err := fmt.Scan(&quantity)
			if err != nil {
				fmt.Println("\nНеккоректные данные!")
			} else {
				break
			}
		}

		if quantity == 0 {
			break
		} else {
			quiz(quantity)
		}
	}
	var l int
	fmt.Scan(&l)
}
