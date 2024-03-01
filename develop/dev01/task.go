package main

import (
	"fmt"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	m := map[string]int{"balls": 2}
	fmt.Println(m["dick"])
	//currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	//
	//if err != nil {
	//	fmt.Fprint(os.Stderr, err.Error())
	//	os.Exit(1)
	//}
	//
	//fmt.Println(currentTime)
}
