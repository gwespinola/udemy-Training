package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main()  {

	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 12)
	// Loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
