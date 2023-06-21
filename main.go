package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"http_md5/request"
)

var parallel = flag.Int("parallel", 10, "Maximum number of parallel requests")

func main() {
	flag.Parse()

	// Reading initial arguments
	urls := flag.Args()
	taskQueue := make(chan string, len(urls))
	for _, url := range urls {
		taskQueue <- url
	}

	threadNumber := *parallel
	for i := 0; i < threadNumber; i++ {
		go func() {
			for url := range taskQueue {
				res, err := request.HttpToMD5(url)
				if err == nil {
					fmt.Printf("%s %s\n", res.Url, res.Hash)
				} else {
					fmt.Printf("Failed to request \"%s\": %s\n", res.Url, err)
				}
			}
		}()
	}

	// Reading arguments after script started running
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		urls = strings.Fields(strings.Replace(text, "\n", "", -1))
		for _, url := range urls {
			taskQueue <- url
		}
	}
}
