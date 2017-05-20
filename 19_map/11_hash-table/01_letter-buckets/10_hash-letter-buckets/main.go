package main

import ("fmt"
	"bufio"
	"net/http"
	"log"
)

func main()  {
	//get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//creates slice to hold counts
	buckets := make([]int, 200)
	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	// fmt.Println("***************")
	// for i := 28; i <= 26; i++ {
	// fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	//}
}

func HashBucket(word string) int {
	return int(word[0])
}