package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func EqualPatterns(pat1, pat2 map[rune]int) bool {
	for k := range pat1 {
		if value, ok := pat2[k]; ok {
			if value != pat1[k] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func GetPattern(runes []rune) map[rune]int {
	pattern := map[rune]int{}
	for _, r := range runes {
		if pattern[r] != 0 {
			pattern[r] += 1
		} else {
			pattern[r] = 1
		}
	}
	return pattern
}

func GetSets(input *[]string) *map[string][]string {
	runes := make([][]rune, len(*input))
	for i, s := range *input {
		runes[i] = []rune(strings.ToLower(s))
	}
	output := map[string][]string{}
	for _, s := range runes {
		curPattern := GetPattern(s)
		counter := 0
		for k := range output {
			if EqualPatterns(GetPattern([]rune(k)), curPattern) {
				output[k] = append(output[k], string(s))
				break
			} else {
				counter++
			}
		}
		if counter == len(output) {
			output[string(s)] = []string{string(s)}
		}
	}
	outputMap := map[string][]string{}
	for k := range output {
		if len(output[k]) != 1 {
			sort.Strings(output[k])
			outputMap[k] = output[k]
		}
	}
	return &outputMap
}

func GetSetsOptimized(input *[]string) *map[string][]string {
	runes := make([][]rune, len(*input))
	for i, s := range *input {
		runes[i] = []rune(strings.ToLower(s))
	}
	output := map[string][]string{}
	for i := range runes {
		if len(runes[i]) != 0 {
			key := string(runes[i])
			output[key] = []string{key}
			currentPattern := GetPattern(runes[i])
			for j := i + 1; j < len(runes); j++ {
				if len(runes[j]) != 0 {
					nextPattern := GetPattern(runes[j])
					if EqualPatterns(nextPattern, currentPattern) {
						output[key] = append(output[key], string(runes[j]))
						runes[j] = []rune{}
					}

				}
			}
			runes[i] = []rune{}
		}
	}
	outputMap := map[string][]string{}
	for k := range output {
		if len(output[k]) != 1 {
			sort.Strings(output[k])
			outputMap[k] = output[k]
		}
	}
	return &outputMap
}

func main() {
	input := strings.Split("Пиво иВоп мясО вопИ соЯм едА", " ")
	//for i, _ := range input {
	//	input[i] = ""
	//}
	//fmt.Println(input)
	sets := GetSetsOptimized(&input)
	for k, v := range *sets {
		fmt.Println("key: ", k)
		fmt.Println(strings.Join(v, " "))
	}
}
