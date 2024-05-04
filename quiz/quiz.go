package main

import "fmt"

/*функция для викторины*/
func quiz(quantity int) {
	var allScores int //Общее количество очков
	winners := make(map[string]int)

	for ind := 1; ind <= quantity; ind++ { //Берём победителя на каждый вопрос
		var (
			localWinner string //Победитель конкурса
			complexity  int    //Сложность вопроса
		)
		fmt.Printf("--- Вопрос %v ---\n", ind)

		fmt.Print("Правильный ответ дал(-а): ")
		fmt.Scan(&localWinner)

		for {
			fmt.Print("Сложность вопроса: ")
			_, err := fmt.Scan(&complexity)
			if err == nil {
				break
			}
		}
		winners[localWinner] += complexity //Добавляем команде баллы
		allScores += complexity            //Увеличиваем общее количество очков
	}
	fmt.Println("--------")
	for team, score := range winners {
		fmt.Printf("%v: %.2f %% \n", team, (float64(score)/float64(allScores))*100)
	}

	fmt.Println("Викторина окончена.\n--------\n--------")
}
