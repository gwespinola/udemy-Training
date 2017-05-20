package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main() {
	// Get the book Adventures of Sherlock Holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
}
