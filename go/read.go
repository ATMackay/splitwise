package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func readFile(fileName string) []transaction {

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	reader := csv.NewReader(f)

	var transactions []transaction

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		t, err := strconv.ParseInt(records[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		f, err := strconv.ParseInt(records[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		a, err := strconv.ParseInt(records[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, [3]int64{t, f, a})
	}

	return transactions
}
