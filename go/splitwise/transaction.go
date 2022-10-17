package splitwise

// transaction represent a transfer of value between two parties
//
// Transaction[0] - sender
// Transaction[1] - receiver
// Transaction[2] - value
//
// For simplicity sake the sender and receiver are assigned a unique numerical value
type Transaction [3]int64

type Transactions []Transaction

func Add(t Transactions, sender, receiver, amt int64) Transactions {
	return append(t, Transaction{sender, receiver, amt})
}

// Scores returns a map of balances given a set of transactions between parties
//
// Transaction[0] --> Transaction[1] : Transaction[3]
// results in a map
// scores[Transaction[0]] = - Transaction[3]
// scores[Transaction[1]] = + Transaction[3]
//
// e.g.
//
//	1 --> 2 : 10
//
// scores[1] = -10
// scores[2] = 10
func Scores(transactions Transactions) (scores map[int64]int64) {

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
