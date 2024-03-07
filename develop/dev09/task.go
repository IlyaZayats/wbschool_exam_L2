package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var visited sync.Map

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r = rand.New(rand.New(rand.NewSource(time.Now().UnixNano())))

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

type DownloadedFile struct {
	Link       string
	FileFormat string
	Body       []byte
}

func DownloadWorker(inputChan <-chan string, saveChan chan<- *DownloadedFile, rChan chan<- *DownloadedFile, m *sync.Mutex, t *time.Ticker) {
	client := http.Client{}
	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	for link := range inputChan {
		t.Stop()
		link, _, _ = strings.Cut(link, `#`)
		if _, ok := visited.Load(link); re.MatchString(link) && !ok {
			visited.Store(link, struct{}{})
			resp, err := client.Get(link)
			if resp.StatusCode == 200 {
				if err != nil {
					fmt.Println(err.Error())
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					_, fileFormat, _ := strings.Cut(resp.Header[`Content-Type`][0], `/`)
					fileFormat = regexp.MustCompile(`\w+`).FindString(fileFormat)
					downloadedFile := &DownloadedFile{FileFormat: fileFormat, Body: body, Link: link}
					saveChan <- downloadedFile
					if err := resp.Body.Close(); err != nil {
						fmt.Println(err.Error())
					}
					m.Lock()
					fmt.Println(`Downloaded file from: `, link)
					fmt.Println(`Content-Type: `, fileFormat)
					m.Unlock()
					if fileFormat == `html` {
						rChan <- downloadedFile
					}
				}
			}
		}
		t.Reset(10 * time.Second)
	}
	fmt.Println("exiting from downloader thread")
}

func WriteWorker(saveChan <-chan *DownloadedFile, dir string, m *sync.Mutex) {
	for downloadedFile := range saveChan {
		var fileName, randomSuffix string
		randomSuffix = RandStringRunes(25)
		if strings.Contains(downloadedFile.FileFormat, `html`) {
			fileName = `index`
		} else {
			if strings.Contains(downloadedFile.FileFormat, "javascript") {
				downloadedFile.FileFormat = `js`
			}
			temp := strings.Split(downloadedFile.Link, `/`)
			fileName, _, _ = strings.Cut(temp[len(temp)-1], `.`)
		}
		savePath := fmt.Sprintf(`%s_%s.%s`, fileName, randomSuffix, downloadedFile.FileFormat)
		if dir != `` {
			savePath = fmt.Sprintf(`%s/%s`, dir, savePath)
		}

		if err := os.WriteFile(savePath, downloadedFile.Body, 0666); err != nil {
			fmt.Println(err.Error())
		}
		m.Lock()
		fmt.Println("Saved at: ", savePath)
		m.Unlock()
	}
	fmt.Println("exiting from writer thread")
}

func ParseWorker(parseChan <-chan *DownloadedFile, dChan chan<- string, m *sync.Mutex, mode bool) {
	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	for downloadedFile := range parseChan {
		for _, item := range regexp.MustCompile(`(src="[\S]+\")`).FindAllString(string(downloadedFile.Body), -1) {
			_, after, found := strings.Cut(item, `src="`)
			if !found {
				_, after, _ = strings.Cut(item, `href="`)
			}
			item = after[:len(after)-1]
			m.Lock()
			fmt.Println("Trying to read: ", item)
			m.Unlock()
			//&&
			linkDownload := re.MatchString(item)
			if mode {
				linkDownload = linkDownload && strings.Contains(item, downloadedFile.Link)
			}
			if linkDownload {
				dChan <- item
			}
			if !strings.Contains(item, `http://`) && !strings.Contains(item, `https://`) {
				b, a, _ := strings.Cut(item, `/`)
				if len(b) == 0 {
					dChan <- fmt.Sprintf(`%s%s`, downloadedFile.Link, a)
				} else {
					dChan <- fmt.Sprintf(`%s%s`, downloadedFile.Link, b)
				}

			}
		}
	}
	fmt.Println("exiting from parser thread")
}

func main() {

	dFlag := flag.String(`d`, `dude`, `storage dir`)
	tFlag := flag.Int(`t`, 10, `timeout (seconds)`)
	dwFlag := flag.Int(`download`, 1, `amount of download workers`)
	wwFlag := flag.Int(`write`, 1, `amount of write workers`)
	pwFlag := flag.Int(`parse`, 1, `amount of parse workers`)
	oFlag := flag.Bool(`only-host`, true, `download mode`)

	flag.Parse()

	url := flag.Arg(0)

	if string(url[len(url)-1]) != `/` {
		url += `/`
	}

	if err := os.Mkdir(*dFlag, 0750); err != nil {
		fmt.Println(err)
	}

	inputChan, saveChan, parseChan := make(chan string, 10), make(chan *DownloadedFile, 10), make(chan *DownloadedFile, 10)

	var m sync.Mutex

	t := time.NewTicker(time.Duration(*tFlag) * time.Second)
	for i := 0; i < *dwFlag; i++ {
		go DownloadWorker(inputChan, saveChan, parseChan, &m, t)
	}
	inputChan <- url
	for i := 0; i < *wwFlag; i++ {
		go WriteWorker(saveChan, *dFlag, &m)
	}
	for i := 0; i < *pwFlag; i++ {
		go ParseWorker(parseChan, inputChan, &m, *oFlag)
	}

	<-t.C
	close(inputChan)
	time.Sleep(1 * time.Second)
	close(saveChan)
	time.Sleep(1 * time.Second)
	close(parseChan)
	time.Sleep(1 * time.Second)
	fmt.Println("exiting from main thread")

}
