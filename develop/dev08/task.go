package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func RemovePlaceholders(str string, replacements map[int]string) string {
	re := regexp.MustCompile(`<not-pipe>\d+</not-pipe>`)
	if re.MatchString(str) {
		for _, item := range re.FindAllString(str, -1) {
			key := strings.TrimLeft(item, `<not-pipe>`)
			key = strings.Trim(key, `</not-pipe>`)
			keyInt, _ := strconv.Atoi(key)
			str = strings.Replace(str, item, replacements[keyInt], -1)
		}
	}
	return str
}

func Parse(command string) string {
	re := regexp.MustCompile(`\$\((.*?)\)`)
	submatchall := re.FindAllString(command, -1)
	if len(submatchall) != 0 {
		for _, item := range submatchall {
			element := strings.Trim(item, `$(`)
			element = strings.Trim(element, `)`)
			result := Parse(element)
			command = strings.Replace(command, item, result, -1)
		}
	}
	replacements := map[int]string{}
	re = regexp.MustCompile(`"(.*?)\|(.*?)"`)
	submatchall = re.FindAllString(command, -1)
	if len(submatchall) != 0 {
		for i, item := range submatchall {
			replacements[i] = item
			command = strings.Replace(command, item, fmt.Sprintf("<not-pipe>%v</not-pipe>", i), -1)
		}
	}
	output := ""

	if strings.Contains(command, `|`) {
		pipeArgs := ``
		for _, pipe := range strings.Split(command, `|`) {
			pipe = RemovePlaceholders(pipe, replacements)
			pipe = strings.TrimSpace(pipe)
			currentCommand := strings.Split(pipe, ` `)
			if strings.Contains(currentCommand[0], `pwd`) {
				pipeArgs, _ = os.Getwd()
			} else if strings.Contains(currentCommand[0], `echo`) {
				pipeArgs = strings.Join(currentCommand[1:], ` `)
				output = pipeArgs
				//fmt.Println(strings.Join(currentCommand[1:], ` `))
			} else if strings.Contains(currentCommand[0], `ps`) {
				pipeArgs = UtilPs()
			} else {
				currentCommand = append(currentCommand, strings.Split(pipeArgs, ` `)...)
				pipeArgs = UtilExec(currentCommand[0], currentCommand[1:]...)
				output = pipeArgs
			}
		}
	} else {
		command = RemovePlaceholders(command, replacements)
		currentCommand := strings.Split(command, ` `)
		if strings.Contains(currentCommand[0], `kill`) {
			pid, _ := strconv.Atoi(currentCommand[1])
			UtilKill(pid)
		} else if strings.Contains(currentCommand[0], `cd`) {
			UtilCd(currentCommand[1])
		} else if strings.Contains(currentCommand[0], `pwd`) {
			output = UtilPwd()
		} else if strings.Contains(currentCommand[0], `echo`) {
			output = strings.Join(currentCommand[1:], ` `)
		} else if strings.Contains(currentCommand[0], `ps`) {
			output = UtilPs()
		} else if strings.Contains(currentCommand[0], `exit`) {
			os.Exit(1)
		} else {
			output = UtilExec(currentCommand[0], currentCommand[1:]...)
		}

	}
	return output
}

func UtilCd(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Println(err.Error())
	}
}

func UtilKill(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := process.Kill(); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func UtilPs() string {
	processList, err := ps.Processes()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var builder strings.Builder
	for x := range processList {
		process := processList[x]
		builder.WriteString(fmt.Sprintf("%d\t%s\n", process.Pid(), process.Executable()))
	}
	return builder.String()
}

func UtilPwd() string {
	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return curDir
}

func UtilExec(command string, inputArgs ...string) string {
	cmd := exec.Command(command, inputArgs...)
	cmdOut, _ := cmd.StdoutPipe()
	cmdErr, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	bytesOut, _ := io.ReadAll(cmdOut)
	bytesErr, _ := io.ReadAll(cmdErr)
	if err := cmd.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	result, errorMsg := string(bytesOut), string(bytesErr)
	if result == "" && errorMsg != "" {
		fmt.Println(errorMsg)
	}
	cmdOut.Close()
	cmdErr.Close()
	return result
}

func main() {

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		curDir := UtilPwd()
		_, after, _ := strings.Cut(curDir, currentUser.HomeDir)
		fmt.Printf("\n%s@%s:~%s$ ", currentUser.Name, host, after)
		scanner.Scan()
		result := Parse(scanner.Text())
		if result != `` {
			fmt.Print(result)
		}
	}
}
