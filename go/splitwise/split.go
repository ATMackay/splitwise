package splitwise

// SimplifyDebts takes a slice of transactions between parties and
// returns the minimum number of payments between parties necessary
// to settle the outstanding debts
func SimplifyDebts(transactions Transactions) (simplifiedDebts Transactions) {

	// compute total balances
	scores := Scores(transactions)

	if !isZeroSum(scores) {
		panic("invalid scores, must be zero sum")
	}

	return greedy(scores, simplifiedDebts)
}

func isZeroSum(scores map[int64]int64) bool {
	var v int64
	for i := range scores {
		v += scores[i]
	}
	return int(v) == 0
}

// greedy is a recursive algorithm that simplifies debts
// within the group given the current set of outstanding balances.
// The algorithm is as follows
//
// Step 1)
//
//	Work out maximum creditor and debtor minimum and maximum scores
//
//		    mD - minimum score, the largest debtor in the group
//		    mC - maximum score, the largest creditor in the group
//
//	        maxDebtor - index for debtor with score mD
//	        maxCreditor - index for creditor with score mC
//
// Step 2)
//
// if |mD| == |mC| then create a transaction from maxDebtor to maxCreditor of mC and eliminate both maxDebtor and maxCreditor from the group
// if |mD| <  |mC| then create a transaction from maxDebtor to maxCreditor of |mD| and eliminate maxDebtor from the group
// if |mD| >  |mC| then create a transaction from maxDebtor to maxCreditor of mC and eliminate maxCreditor from the group
//
// Step 3)
//
//	Repeat Step 1) - 2) until the group is eliminated and return
//
// WARNING: Does not protect against malicious inputs,
// To avoid an infinite loop use isZeroSum() to check if the input scores are valid
// before executing this function
func greedy(scores map[int64]int64, existingDebts Transactions) (debts Transactions) {

	if len(scores) == 0 {
		return existingDebts
	}

	maxCreditor, mC := maxEntry(scores)

	maxDebtor, mD := minEntry(scores)

	if mC == -mD {
		debts = Add(existingDebts, maxDebtor, maxCreditor, mC)
		delete(scores, maxDebtor)
		delete(scores, maxCreditor)
	} else if mC > -mD {
		debts = Add(existingDebts, maxDebtor, maxCreditor, -mD)
		delete(scores, maxDebtor)
		scores[maxCreditor] += mD
	} else if mC < -mD {
		debts = Add(existingDebts, maxDebtor, maxCreditor, mC)
		delete(scores, maxCreditor)
		scores[maxDebtor] += mC
	}

	return greedy(scores, debts)
}

func maxEntry(m map[int64]int64) (index, value int64) {
	// find first element
	for i := range m {
		index = i
		value = m[index]
		break
	}
	// obtain max element
	for x := range m {
		if m[x] > value {
			index, value = x, m[x]
		}
	}
	return
}

func minEntry(m map[int64]int64) (index, value int64) {
	// find first element
	for i := range m {
		index = i
		value = m[index]
		break
	}
	// obtain min element
	for x := range m {
		if m[x] < value {
			index, value = x, m[x]
		}
	}
	return
}
