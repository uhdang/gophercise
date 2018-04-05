package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("problems.csv")
	check(err)
	defer file.Close()

	// os.Open returns a "type File", which implements "Read" method
	// csv.NewReader accepts "io.Reader", or type Reader, which also implements "Read" method
	// Go interfaces are based on capabilities expressed implicitly from object set of methods.
	// Consequently, csv.NewReader also accepts whats been returned from os.Open,
	// since they both implement "Read" Method
	// https://golang.org/pkg/io/#Reader

	r := csv.NewReader(file)
	// NewReader expects io.Reader and returns *Reader

	total := 0
	correct := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		// Read accepts *Reader and returns []string.
		i, err := strconv.Atoi(record[1])
		check(err)
		var input int
		fmt.Printf("Question: " + record[0] + " = ")
		fmt.Scanf("%d", &input)

		//fmt.Println("Check: ", record, record[0], record[1])
		if input == i {
			correct++
		}
		total++

	}

	fmt.Printf("Score: %d/%d\n", correct, total)

}
