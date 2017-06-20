package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"net/http"
	"io"
	"io/ioutil"
)


func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)		// Start a goroutine
	}
	for range os.Args[1:]{
		fmt.Println(<- ch)		// Receive from channel ch
	}
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Cannot create file",err)
	}
	defer file.Close()
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())	// Prints result to output file
	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)		// Send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()			// Don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/* Find a website that produces a large amount of data. Investigate catching by running fetchall twice in succession
to see whether the reported time changes much. Do you get the same content each time? Modify fetchall to to print its
output to a file so it can be examined
*/
