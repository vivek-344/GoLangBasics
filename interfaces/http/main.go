package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// There a lot going around Reader and Writer Interface
	// Read Documentation :)
	alphabets, errAl := countAlphabets(resp.Body)

	if errAl != nil {
		panic(errAl)
	}

	fmt.Printf("Letters: %v", alphabets)
}

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buff := make([]byte, 1024)

	for {
		n, err := r.Read(buff)
		for _, ch := range buff[:n] {
			if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return 0, err
		}
	}
}
