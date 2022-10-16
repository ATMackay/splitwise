package main

// transaction represent a transfer of value between two parties
//
// transaction[0] - sender
// transaction[1] - receiver
// transaction[2] - value
//
// For simplicity sake the sender and reciver are assigned a unique numerical value
type transaction [3]int64

func scores(transactions []transaction) (scores map[int64]int64) {

	scores = make(map[int64]int64)

	for _, tx := range transactions {
		if tx[0] == tx[1] {
			continue
		}
		scores[tx[0]] -= tx[2]
		scores[tx[1]] += tx[2]
	}
	return
}
