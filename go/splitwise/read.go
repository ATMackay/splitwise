package splitwise

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// ReadFile opens a target csv file 'fileName'
// and reads the contents into a Transactions struct
// will panic if the csv contains invalid inputs (see ../test_data for valid examples)
func ReadFile(fileName string) Transactions {

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	reader := csv.NewReader(f)

	var transactions Transactions = Transactions{}

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		t, err := strconv.ParseInt(records[0], 10, 64)
		if err != nil {
			panic(err)
		}
		f, err := strconv.ParseInt(records[1], 10, 64)
		if err != nil {
			panic(err)
		}
		a, err := strconv.ParseInt(records[2], 10, 64)
		if err != nil {
			panic(err)
		}
		transactions = Add(transactions, t, f, a)
	}

	return transactions
}
