package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

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

func AFlagProcess(matchedIndex, flagValue int, lines []string, matchedLines map[int]string) map[int]string {
	for j := matchedIndex + flagValue; j >= matchedIndex; j-- {
		if j < len(lines) {
			for k := j; k >= matchedIndex; k-- {
				matchedLines[k] = lines[k]
			}
			break
		}
	}
	return matchedLines
}

func BFlagProcess(matchedIndex, flagValue int, lines []string, matchedLines map[int]string) map[int]string {
	for j := matchedIndex - flagValue; j <= matchedIndex; j++ {
		if j > -1 {
			for k := j; k <= matchedIndex; k++ {
				matchedLines[k] = lines[k]
			}
			break
		}
	}
	return matchedLines
}

func main() {

	AFlag := flag.Int("A", 0, "after content")
	BFlag := flag.Int("B", 0, "before content")
	CFlag := flag.Int("C", 0, "around content")
	cFlag := flag.Bool("c", false, "found lines")
	iFlag := flag.Bool("i", false, "ignore case")
	vFlag := flag.Bool("v", false, "invert")
	FFlag := flag.Bool("F", false, "fixed")
	nFlag := flag.Bool("n", false, "line num")

	flag.Parse()

	pattern, fileName := flag.Arg(0), flag.Arg(1)

	linesInput, err := ReadLines(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	lines := make([]string, len(linesInput))

	if *iFlag {
		for i := range lines {
			lines[i] = strings.ToLower(linesInput[i])
		}
	} else {
		for i := range lines {
			lines[i] = linesInput[i]
		}
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	matchedLines := map[int]string{}
	found := false

	for i, line := range lines {
		found = false

		if *FFlag {
			currentLine := strings.Split(line, " ")
			for _, c := range currentLine {
				if c == pattern {
					found = true
					break
				}
			}
		} else if re.MatchString(line) {
			found = true
		}

		if found {
			if !*cFlag && !*nFlag && !*vFlag {
				if *BFlag > 0 {
					matchedLines = BFlagProcess(i, *BFlag, lines, matchedLines)
				}

				if *AFlag > 0 {
					matchedLines = AFlagProcess(i, *AFlag, lines, matchedLines)
				}

				if *CFlag > 0 {
					matchedLines = BFlagProcess(i, *CFlag, lines, matchedLines)
					matchedLines = AFlagProcess(i, *CFlag, lines, matchedLines)
				}
			}
			if (*BFlag == 0 && *AFlag == 0 && *CFlag == 0) || *cFlag || *nFlag || *vFlag {
				matchedLines[i] = line
			}
		}
	}

	keys := make([]int, len(matchedLines))
	index := 0
	for k := range matchedLines {
		keys[index] = k
		index++
	}
	sort.Ints(keys)

	if *cFlag {
		fmt.Println(len(matchedLines))
		os.Exit(0)
	}

	if *nFlag {
		for _, k := range keys {
			fmt.Printf("%v:%s\n", k+1, linesInput[k])
		}
		os.Exit(0)
	}

	if *vFlag {
		for _, k := range keys {
			for j := 0; j < len(lines); j++ {
				if lines[j] == matchedLines[k] {
					lines[j] = ""
				}
			}
		}
		for i := range lines {
			if lines[i] != "" {
				fmt.Println(linesInput[i])
			}
		}
		os.Exit(0)
	}

	for _, k := range keys {
		fmt.Println(linesInput[k])
	}
}
