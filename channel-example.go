//пример работы с каналами
//Я НЕ СМОГЛА ПОНЯТЬ, ПОЧЕМУ ДАННЫЕ НЕ ПЕРЕДАЮТСЯ В ВЫВОДЕ НА КОНСОЛИ, %v - ДОЛЖНО РАБОТАТЬ, НО ПОКА Я НЕ НАШЛА ОШИБКУ,
//ПОЧЕМУ НЕ ВЫВОДЯТСЯ БУКВЫ A, B, C, D

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
		fmt.Println(writeFormat, number, string(symbol))
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
	go writeToChan(first, "ABCD", "First")
	go writeToChan(second, "abcd", "Second")

	//вывод того, что в канале есть, каждую секунду
	for {
		time.Sleep(1 * time.Second)
		select {
		case firstData := <-first:
			fmt.Println(readFormat, "First", firstData)
		case secondData := <-second:
			fmt.Println(readFormat, "Second", secondData)
		default: //если данных для получения нет в каналах
			fmt.Println("Default data")
			return
		}
	}
}
