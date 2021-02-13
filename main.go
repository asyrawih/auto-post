package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/hananloser/auto-post/libs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupHandler(){
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func main() {
	setupHandler();
	c := colly.NewCollector(colly.AllowedDomains("komikcast.com"), colly.Async(true))
	lists := libs.GetComics(c)
	for _ , list := range lists{
		fmt.Println(list.Name)
	}
	for {
		fmt.Println("- Sleeping")
		time.Sleep(10 * time.Second)
	}
}
