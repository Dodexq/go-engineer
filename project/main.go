package main

import (
	"bufio"
	"fmt"
	"linkedlist/mypackage"
	"os"
	"strconv"
)

func main() {
	l := &mypackage.LinkedList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите число для добавления в список (или 'exit' для выхода): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите целое число.")
			continue
		}

		fmt.Print("Введите 'begin' для вставки в начало, 'end' для вставки в конец: ")
		scanner.Scan()
		position := scanner.Text()

		switch position {
		case "begin":
			l.InsertAtPrev(num)
		case "end":
			l.InsertAtBack(num)
		default:
			fmt.Println("Некорректное значение. Введите 'begin' или 'end'.")
			continue
		}

		fmt.Print("Список после добавления: ")
		l.ShowAll()
	}
}
