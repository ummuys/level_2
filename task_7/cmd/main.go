package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int { // Здесь создаются каналы
	c := make(chan int) // Небуферизированный канал
	go func() {
		for _, v := range vs {
			c <- v // Записывается одно значение и ждем, когда его считают
			// Если его никто читать не будет, то рутина зависнет
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c) // После записи всех чисел канал закрывается
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int) // Создается новый канал, который будет объединять каналы
	go func() {         // Запускаем рутину, чтобы мы могли записывать данные и одновременно вернуть канал с функции
		for { // Цикл обеспечивает постоянное чтение данных
			select { // Select выбирает case с готовым каналом (т.е от кого первее придет ответ)
			// Если одинаково готовы все, то выбор делается псевдослучайно
			case v, ok := <-a:
				if ok { // Если канал закрыт, то ok == false
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil { // Если оба канала закрыты, то закрываем с
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c { // Читаем все значения с канала до тех пор, пока он не закрыт
		fmt.Print(v)
	}
}
