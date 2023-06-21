package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"http_md5/request"
)

var parallel = flag.Int("parallel", 10, "Maximum number of parallel requests")

func main() {
	flag.Parse()

	urls :=  flag.Args()
	taskQueue := make(chan string, len(urls))
	for _, url := range urls {
		taskQueue <- url
	}
	close(taskQueue)

	threadNumber := *parallel
	if len(urls) < threadNumber {
		threadNumber = len(urls)
	}

	for i := 0; i < threadNumber; i++ {
		go func() {
			for url := range taskQueue {
				res, err := request.HttpToMD5(url)
				if err == nil {
					fmt.Printf("%s %s\n", res.Url, res.Hash)
				}
			}
		}()
	}

	// only stopping on ^C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
