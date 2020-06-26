//пример работы с каналами
//Я НЕ СМОГЛА ПОНЯТЬ, ПОЧЕМУ ДАННЫЕ НЕ ПЕРЕДАЮТСЯ В ВЫВОДЕ НА КОНСОЛИ, А ИМЕННО БУКВЫ A, B, C, D
//(в функции writeToChan это второй параметр)

package main

import (
	"fmt"
	"time"
)

var (
	readFormat  = "Получено %-8v: %v\n"
	writeFormat = "Отправлено %-8v: %v\n"
)

//отправить строку по одному символу каналу
//в функцию передается канал, в который записываются данные; сами данные и номер канала(first, second)
func writeToChan(channel chan string, text string, number string) {
	for symbol := range text { //считывается по символьно

		fmt.Printf(writeFormat, number, string(symbol))
		channel <- string(symbol) //запись в канал
	}
}

func main() {

	//создаем два канала
	first := make(chan string)
	second := make(chan string)

	//после закрыть каналы
	defer close(first)
	defer close(second)

	//отправляем строки в канал
	go writeToChan(first, "ABCD", "FIRST")
	go writeToChan(second, "abcd", "SECOND")

	//вывод того, что в канале есть, каждую секунду
	for {
		time.Sleep(1 * time.Second)
		select {
		case firstData := <-first:
			fmt.Println(readFormat, "FIRST", firstData)
		case secondData := <-second:
			fmt.Println(readFormat, "SECOND", secondData)
		default: //если данных для получения нет в каналах
			fmt.Println("Default data")
			return
		}
	}
}
