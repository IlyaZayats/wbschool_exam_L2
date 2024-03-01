package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func ParseFFlagValue(fFlagValue string) ([]int, int, error) {
	index := strings.Index(fFlagValue, "-")
	if index != -1 {
		from, to := 0, -1
		var err error
		if index-1 >= 0 {
			from, err = strconv.Atoi(fFlagValue[:index])
			from -= 1
			if err != nil {
				return nil, -1, err
			}
		}
		if index+1 < len(fFlagValue) {
			to, err = strconv.Atoi(fFlagValue[index+1:])
			if err != nil {
				return nil, -1, err
			}
		}
		return []int{from, to}, 2, nil
	}

	index = strings.Index(fFlagValue, ",")
	if index != -1 {
		var output []int
		before, after := "", fFlagValue
		for {
			before, after, _ = strings.Cut(after, ",")
			if before == "" {
				break
			}
			n, err := strconv.Atoi(before)
			if err != nil {
				return nil, 3, err
			}
			output = append(output, n-1)
		}
		return output, 3, nil
	}

	n, err := strconv.Atoi(fFlagValue)
	if err != nil {
		return nil, 1, err
	}
	return []int{n - 1}, 1, nil
}

func Cut(linesInput []string, delimiter string, isNeededDelimiter bool, neededFields []int, typeOfFlag int) []string {
	var outputLines []string
	for _, line := range linesInput {
		lineSplit := strings.Split(line, delimiter)

		if !(isNeededDelimiter && len(lineSplit) == 1) {
			switch typeOfFlag {
			case 1:
				{
					if neededFields[0] < len(lineSplit) {
						outputLines = append(outputLines, lineSplit[neededFields[0]])
					}
					break
				}
			case 2:
				{
					from, to := neededFields[0], neededFields[1]
					if neededFields[1] == -1 {
						to = len(lineSplit) - 1
					}
					if to > from {
						outputLines = append(outputLines, strings.Join(lineSplit[from:to], delimiter))
					}
					break
				}
			default:
				{
					var output []string
					for _, n := range neededFields {
						if n > -1 && n < len(lineSplit) {
							output = append(output, lineSplit[n])
						}
					}
					outputLines = append(outputLines, strings.Join(output, delimiter))
				}
			}
		}
	}
	return outputLines
}

func main() {
	fFlag := flag.String("f", "1", "fields")
	dFlag := flag.String("d", " ", "delimiter")
	sFlag := flag.Bool("s", false, "separated")

	flag.Parse()

	fileName := flag.Arg(0)
	//fileName := "input.txt"

	linesInput, err := ReadLines(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	neededFields, typeOfFlag, err := ParseFFlagValue(*fFlag)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	outputLines := Cut(linesInput, *dFlag, *sFlag, neededFields, typeOfFlag)

	for _, l := range outputLines {
		fmt.Println(l)
	}

}
