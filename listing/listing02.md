Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
test() вернет 2
anotherTest() вернет 1
Это связано с тем, что anotherTest() не имеет named return и возвращает старое значение x
defer добавляет отложенные функции в стэк и вызывает их перед выходом из "родительской" функции

```
