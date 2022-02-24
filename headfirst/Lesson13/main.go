package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)
	usrl := []string{
		"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc",
	}
	for _, url := range usrl {
		go responseSize(url, pages)
	}
	for i := 0; i < len(usrl); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{
		URL:  url,
		Size: len(body),
	}
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}

}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}
