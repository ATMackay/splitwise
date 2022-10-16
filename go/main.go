package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("splitwise algorithm in go")

	start := time.Now()
	transactions := readFile("../test_data/input.csv")
	fmt.Printf("completed csv load execution in %v microseconds\n", time.Since(start).Microseconds())

	restart := time.Now()

	txs := simplifyDebts(transactions)
	fmt.Printf("completed simplifyDebts execution in %v microseconds, min transactions %v\n", time.Since(restart).Microseconds(), len(txs))

	fmt.Printf("transactions: %v\n", txs)
}
