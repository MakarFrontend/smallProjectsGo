package main

import "fmt"

var (
	toEnrypt string //Текст для шифровки, расшифровки
	result   string //Результат
	whatDo   int    //Что делать?
)

func main() {
	fmt.Println("Шифровальщик")

	fmt.Println("Введи текст: ")
	fmt.Scan(&toEnrypt)
	fmt.Println("Что делать? 1-шифровать, 2-расшифровать")
	fmt.Scan(&whatDo)

	switch whatDo {
	case 1:
		result = encrypt(toEnrypt)
		fmt.Printf("Результат: %s", result)
	case 2:
		result = deencrypt(toEnrypt)
		fmt.Printf("Результат: %s", result)
	default:
		fmt.Println("Неизвестное действие!")
	}

	var g int
	fmt.Scan(&g)
}
