package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println("splitwise algorithm in go")

	start := time.Now()
	transactions := readFile("../test_data/input.csv")
	fmt.Printf("completed csv load execution in %v microseconds\n", time.Since(start).Microseconds())

	minTxs := minTransfers(transactions)
	fmt.Printf("completed minTransfers execution in %v microseconds\n", time.Since(start).Microseconds())

	fmt.Printf("minimum number of transactions was %d\n", minTxs)
}

func readFile(fileName string) [][3]int {

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	reader := csv.NewReader(f)

	var transactions [][3]int

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		t, err := strconv.Atoi(records[0])
		if err != nil {
			log.Fatal(err)
		}
		f, err := strconv.Atoi(records[1])
		if err != nil {
			log.Fatal(err)
		}
		a, err := strconv.Atoi(records[2])
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, [3]int{t, f, a})
	}

	return transactions
}

func minTransfers(transactions [][3]int) int {

	scores := make(map[int]int)

	for _, tx := range transactions {
		if tx[0] == tx[1] {
			continue
		}
		scores[tx[0]] -= tx[2]
		scores[tx[1]] += tx[2]
	}

	var positives, negatives []int
	for i := 0; i < len(scores)+1; i++ {
		if scores[i+1] < 0 {
			negatives = append(negatives, scores[i+1])
		} else if scores[i+1] > 0 {
			positives = append(positives, scores[i+1])
		}
	}

	fmt.Printf("SCORES: %v\n", scores)
	fmt.Printf("POSITIVES: %v\n", positives)
	fmt.Printf("NEGATIVES: %v\n", negatives)

	return recurse(positives, negatives)
}

func recurse(positives, negatives []int) int {

	if len(positives)+len(negatives) == 0 {
		return 0
	}

	if len(negatives) == 0 {
		// error
		return math.MaxInt - 1
	}

	//fmt.Printf("len(positives) %v\n", len(positives))
	//fmt.Printf("len(negatives) %v\n", len(negatives))

	negative := negatives[0]

	fmt.Printf("negative %v\n", negative)

	count := math.MaxInt - 1
	for i := range positives {
		var new_positives []int
		var new_negatives []int
		copy(new_positives, positives)
		copy(new_negatives, negatives)

		new_positives = remove(new_positives, positives[i])
		new_negatives = remove(new_negatives, negative)

		if positives[i] == -negative {
			// do nothing
		} else if positives[i] > -negative {
			new_positives = append(new_positives, positives[i]+negative)
		} else {
			new_negatives = append(new_negatives, positives[i]+negative)
		}

		count = min(count, recurse(new_positives, new_negatives))

	}
	return count + 1

}

func remove(array []int, element int) (updatedArray []int) {
	for i := range array {
		if element != array[i] {
			updatedArray = append(updatedArray, array[i])
		}
	}
	return updatedArray
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
