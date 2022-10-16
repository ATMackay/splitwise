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

	/*
		restart := time.Now()
		minTxs := minTransfers(transactions)
		fmt.Printf("completed minTransfers execution in %v microseconds\n", time.Since(restart).Microseconds())

		fmt.Printf("minimum number of transactions was %d\n", minTxs)
	*/

	restart_again := time.Now()

	txs := simplifyDebts(transactions)
	fmt.Printf("completed simplifyDebts execution in %v microseconds, min transactions %v\n", time.Since(restart_again).Microseconds(), len(txs))

	fmt.Printf("transactions: %v\n", txs)
}
