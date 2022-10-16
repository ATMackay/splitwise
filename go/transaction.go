package main

import (
	"fmt"
	"math"
)

// transaction represent a transfer of value between two parties
//
// transaction[0] - sender
// transaction[1] - receiver
// transaction[2] - value
//
// For simplicity sake the sender and reciver are assigned a unique numerical value
type transaction [3]int64

func minTransfers(transactions []transaction) int64 {

	scores := scores(transactions)

	negatives, positives := splitScores(scores)

	fmt.Printf("SCORES: %v\n", scores)
	fmt.Printf("POSITIVES: %v\n", positives)
	fmt.Printf("NEGATIVES: %v\n", negatives)

	return recurseSlice(positives, negatives)
}

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

func splitScores(scores map[int64]int64) (negatives, positives []int64) {
	for i := int64(0); i < int64(len(scores)+1); i++ {
		if scores[i+1] < 0 {
			negatives = append(negatives, scores[i+1])
		} else if scores[i+1] > 0 {
			positives = append(positives, scores[i+1])
		}
	}
	return
}

func splitScoreMap(scores map[int64]int64) (negatives, positives map[int64]int64) {
	for i := int64(0); i < int64(len(scores)+1); i++ {
		if scores[i+1] < 0 {
			negatives[i+1] = scores[i+1]
		} else if scores[i+1] > 0 {
			positives[i+1] = scores[i+1]
		}
	}
	return
}

func recurseSlice(positives, negatives []int64) int64 {

	if len(positives)+len(negatives) == 0 {
		return 0
	}

	negative := negatives[0]

	count := int64(math.MaxInt - 1)
	for i := range positives {
		var new_positives []int64
		var new_negatives []int64

		new_positives = positives
		new_negatives = negatives

		new_positives = remove(new_positives, positives[i])
		new_negatives = remove(new_negatives, negative)

		if positives[i] == -negative {
			// do nothing
		} else if positives[i] > -negative {
			new_positives = append(new_positives, positives[i]+negative)
		} else {
			new_negatives = append(new_negatives, positives[i]+negative)
		}

		count = min(count, recurseSlice(new_positives, new_negatives))

	}
	return count + 1

}

func remove(array []int64, element int64) (updatedArray []int64) {
	for i := range array {
		if element != array[i] {
			updatedArray = append(updatedArray, array[i])
		}
	}
	return updatedArray
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
