package main

import (
	"fmt"
	"time"

	"github.com/ATMackay/splitwise/go/splitwise"
)

func main() {

	fmt.Println("splitwise algorithm in go")

	// Demonstrate splitwise package APIs
	start := time.Now()
	// Read csv file
	transactions := splitwise.ReadFile("../test_data/input.csv")
	fmt.Printf("completed csv load execution in %v\n", time.Since(start))

	restart := time.Now()
	// Simplify debts
	txs := splitwise.SimplifyDebts(transactions)
	fmt.Printf("completed SimplifyDebts execution in %v, settlement transactions %v\n", time.Since(restart), len(txs))

	fmt.Printf("transactions: %v\n", txs)
}
