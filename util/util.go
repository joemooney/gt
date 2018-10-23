package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Startswith return true if string startwith substr
func Startswith(s, substr string) bool {
	return strings.HasPrefix(s, substr)
}

// Endswith return true if string startwith substr
func Endswith(s, substr string) bool {
	return strings.HasSuffix(s, substr)
}

func readlinesAsync(path string, scanner *bufio.Scanner, lines chan<- string) {
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	//fmt.Println("Closing channel" + path)
	close(lines)
}

// Readlines returns a channel for reading all the lines in a file
func Readlines(path string) chan string {

	ch := make(chan string)
	//fmt.Printf("Opening channel %v %v\n", bigi, path)
	//runtime.SetFinalizer(&ch, func(c *chan string) { fmt.Println("Channel is dead " + path) })

	// open a file
	if file, err := os.Open(path); err == nil {
		//runtime.SetFinalizer(&file, func(c **os.File) { fmt.Println("File is dead " + path) })

		// make sure it gets closed

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		//runtime.SetFinalizer(&scanner, func(c **bufio.Scanner) { fmt.Println("Sanner is dead " + path) })
		go readlinesAsync(path, scanner, ch)

		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}
	return ch
}
