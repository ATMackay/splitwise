package main

// simplifyDebts takes a slice of transactions between parties and
// returns the minimum number of payments between parties necessary
// to settle the outstanding debts
func simplifyDebts(transactions []transaction) (debts []transaction) {

	scores := scores(transactions)

	return greedy(scores, debts)
}

func greedy(scores map[int64]int64, existingDebts []transaction) (debts []transaction) {

	if len(scores) == 0 {
		return existingDebts
	}

	maxCreditor, maxCredit := maxEntry(scores)

	maxDebtor, maxDebt := minEntry(scores)

	if maxCredit == -maxDebt {
		debts = append(existingDebts, [3]int64{maxDebtor, maxCreditor, -maxDebt})
		delete(scores, maxDebtor)
		delete(scores, maxCreditor)
	} else if maxCredit > -maxDebt {
		debts = append(existingDebts, [3]int64{maxDebtor, maxCreditor, -maxDebt})
		delete(scores, maxDebtor)
		scores[maxCreditor] += maxDebt
	} else if maxCredit < -maxDebt {
		debts = append(existingDebts, [3]int64{maxDebtor, maxCreditor, maxCredit})
		delete(scores, maxCreditor)
		scores[maxDebtor] += maxCredit
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
