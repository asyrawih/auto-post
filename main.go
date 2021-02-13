package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/hananloser/auto-post/Comics"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

	c := colly.NewCollector(colly.AllowedDomains("komikcast.com"), colly.Async(true))

	// Read Input From Console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	fmt.Println("1 . Get Hot Updates")
	fmt.Println("2 . Latest Updates")
	fmt.Println("0 . Exit")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			lists := Comics.GetComics(c)
			for _, list := range lists {
				fmt.Println(list.Name , list.Link)
			}
		}

		// Get Latest Comics
		if strings.Compare("2" , text) == 0 {
			latests := Comics.LastUpdate(c)
			for index, latest := range latests {
				fmt.Println(index + 1 , "." , latest.Name , latest.Link)
			}
		}

		teriminate()
	}
}

func teriminate() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
