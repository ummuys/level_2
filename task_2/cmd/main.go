package main

import "fmt"

func test() (x int) { // Именованное возвращение
	defer func() { // Специальная команда, которая откладывает выполнение нужной команды после выхода с функции
		x++
	}()
	x = 1
	return // Так как return возвращает ИМЕННО x, то defer сможет внести изменения в результат
}

func anotherTest() int { // Безымянное возвращение
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x // Тут функция забирает именно значения с x, поэтому в defer
	// хоть и происходит измения x, но вернется 1
}

func main() {
	fmt.Println(test())        // 2
	fmt.Println(anotherTest()) // 1
}
