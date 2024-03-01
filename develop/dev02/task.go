package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func joinRunes(r rune, length int) []rune {
	var sb strings.Builder
	for j := 0; j < length; j++ {
		sb.WriteRune(r)
	}
	return []rune(sb.String())
}

func isNumber(r rune) bool {
	return r >= '1' && r <= '9'
}

func isEscape(r rune) bool {
	return r == '\\'
}

func isSym(r rune) bool {
	return !isEscape(r) && !isNumber(r)
}

var WrongStringError = errors.New("got wrong string")

func UnwrapString(inputStr string) (string, error) {
	if len(inputStr) == 0 {
		return "", nil
	}
	var re = regexp.MustCompile(`(?m)[[:alpha:]][0-9]|[[:alpha:]]|[\\][\\][0-9]|[\\][\\]|[\\][0-9][0-9]|[\\][0-9]`)
	match := re.FindAllString(inputStr, -1)
	sb := strings.Builder{}
	for _, m := range match {
		sb.WriteString(m)
	}
	if sb.String() != inputStr {
		return "", WrongStringError
	}
	sb.Reset()
	for _, m := range match {
		runes := []rune(m)
		var output string
		if len(runes) == 1 {
			output = string(runes[0])
		}
		if len(runes) == 2 {
			if isSym(runes[0]) {
				output = string(joinRunes(runes[0], int(runes[1]-'0')))
			} else {
				output = string(runes[1])
			}
		}
		if len(runes) == 3 {
			output = string(joinRunes(runes[1], int(runes[2]-'0')))
		}
		sb.WriteString(output)
	}
	return sb.String(), nil
}

func main() {
	output, err := UnwrapString(`we2we2`)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(output)
}
