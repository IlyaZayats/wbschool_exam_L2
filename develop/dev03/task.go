package main

import (
	"bufio"
	ts "dev03/sort"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки//
-n — сортировать по числовому значению//
-r — сортировать в обратном порядке//
-u — не выводить повторяющиеся строки//

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы//
-c — проверять отсортированы ли данные//
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Equal(a, b []string) error {
	if len(a) != len(b) {
		return errors.New("len of string error")
	}
	for i, v := range a {
		if v != b[i] {
			return fmt.Errorf("disorder: line: %v, value: %s", i+1, v)
		}
	}
	fmt.Println("file in order")
	return nil
}

func splitLines(lines []string) [][]string {
	linesM := make([][]string, len(lines))
	for i, line := range lines {
		linesM[i] = strings.Split(line, " ")
	}
	return linesM
}

func readLines(fileName string) ([]string, error) {
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

func writeLines(fileName string, lines []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		line += "\n"
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}

func main() {

	kFlag := flag.Int("k", 1, "key")
	nFlag := flag.Bool("n", false, "numeric sort")
	rFlag := flag.Bool("r", false, "reverse")
	uFlag := flag.Bool("u", false, "unique")
	cFlag := flag.Bool("c", false, "check")
	bFlag := flag.Bool("b", false, "ignore leading blanks")
	hFlag := flag.Bool("h", false, "human numeric sort")
	MFlag := flag.Bool("M", false, "month sort")

	flag.Parse()

	lines, err := readLines(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *bFlag {
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}
	}

	if *uFlag {
		unique := map[string]struct{}{}
		for i := range lines {
			unique[lines[i]] = struct{}{}
		}
		newLines := make([]string, 0, len(unique))
		for k := range unique {
			newLines = append(newLines, k)
		}
		lines = newLines
	}

	var sorted [][]string
	isNotSorted := true
	linesM := splitLines(lines)
	column := *kFlag

	if column > len(linesM[0]) {
		column = 1
	}

	var numeric []ts.NumericStruct

	switch {
	case *nFlag:
		{
			for i, line := range linesM {
				num, err := strconv.ParseFloat(line[column-1], 64)
				if err != nil {
					sorted = append(sorted, line)
				} else {
					numeric = append(numeric, ts.NumericStruct{Index: i, Num: num})
				}
			}
			break
		}
	case *MFlag:
		{
			for i, line := range linesM {
				n := 0
				switch line[column-1] {
				case `DEC`:
					n = 12
				case `JAN`:
					n = 11
				case `FEB`:
					n = 10
				case `MAR`:
					n = 9
				case `APR`:
					n = 8
				case `MAY`:
					n = 7
				case `JUN`:
					n = 6
				case `JUL`:
					n = 5
				case `AUG`:
					n = 4
				case `SEP`:
					n = 3
				case `OCT`:
					n = 2
				case `NOV`:
					n = 1
				}
				numeric = append(numeric, ts.NumericStruct{Index: i, Num: float64(n)})
			}
			break
		}
	case *hFlag:
		{
			reDigit := regexp.MustCompile(`(?m)[\d]+`)
			reSuffix := regexp.MustCompile(`(?m)[^\d]+`)

			for i, line := range linesM {
				digit := reDigit.FindString(line[column-1])
				suffix := reSuffix.FindString(line[column-1])
				builder := strings.Builder{}
				builder.WriteString(digit)
				builder.WriteString(suffix)
				if len(builder.String()) != len(line[column-1]) {
					sorted = append(sorted, line)
				} else {
					n, _ := strconv.ParseFloat(digit, 64)
					switch strings.ToLower(suffix) {
					case `tb`:
						n *= 1024 * 1024 * 1024
					case `gb`:
						n *= 1024 * 1024
					case `mb`:
						n *= 1024
					case `b`:
						n /= 1024
					}
					numeric = append(numeric, ts.NumericStruct{Index: i, Num: n})
				}
				builder.Reset()
			}
			break
		}
	}

	//numeric sort: -h, -n, -M
	if len(numeric) != 0 {
		if *rFlag {
			sort.Sort(sort.Reverse(ts.NumericSort(numeric)))
		} else {
			sort.Sort(ts.NumericSort(numeric))
		}
		for _, n := range numeric {
			sorted = append(sorted, linesM[n.Index])
		}

		isNotSorted = false
	} else {
		sorted = [][]string{}
	}

	//string sort
	if isNotSorted {
		var strs []ts.StringStruct
		for i, line := range linesM {
			strs = append(strs, ts.StringStruct{Index: i, Str: line[column-1]})
		}

		if *rFlag {
			sort.Sort(sort.Reverse(ts.StringSort(strs)))
		} else {
			sort.Sort(ts.StringSort(strs))
		}
		for _, s := range strs {
			sorted = append(sorted, linesM[s.Index])
		}
	}

	var output []string

	for i := range sorted {
		output = append(output, strings.Join(sorted[i], " "))
	}

	if *cFlag {
		if err := Equal(lines, output); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		for _, line := range output {
			fmt.Println(line)
		}
		if err := writeLines(flag.Arg(1), output); err != nil {
			fmt.Println(err.Error())
		}
	}

}
