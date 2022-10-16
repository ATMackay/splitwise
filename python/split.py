from collections import defaultdict
from array import *
import numpy as np
import math
import time
"""
This programme takes an input file containing trading data, calculates the implied volatility for each 
trade and ouputs a new csv file.

Requirements: Python 3.6 (or later), numpy, pandas, scipy


Installation of NumPy, SciPy and Pandas using pip --> $ python3 -m pip install --user numpy scipy pandas
"""


"""
Solution

Class containing minimum transaction algorithm and csv decoding 
"""
class Solution: 

    def __init__(self):      
        # do nothing
        return

    def load(self, input_csv):
        return np.genfromtxt(input_csv, delimiter=',')

    def min_transfers(self, transactions) -> int:

        scores = scores(transactions)
        
        positives = [v for v in scores.values() if v > 0]
        negatives = [v for v in scores.values() if v < 0]


        print(f"scores: {scores}")
        print(f"positives {positives}")
        print(f"negatives {negatives}")

        def recurse(positives, negatives):
            if len(positives) + len(negatives) == 0: return 0

            negative = negatives[0]

            count = math.inf
            for positive in positives:

                new_positives = positives.copy()
                new_negatives = negatives.copy()

                new_positives.remove(positive)
                new_negatives.remove(negative)

                if positive == -negative:
                    pass
                elif positive > -negative:
                    new_positives.append(positive+negative)
                else:
                    new_negatives.append(positive+negative)
                
                count = min(count, recurse(new_positives, new_negatives))
        
            return count + 1

        return recurse(positives, negatives)

    def simplifyDebts(self, transactions):
        scores = scores(transactions)

        debts = np.ndarray()
        def greedy(scores, debt_array):
            if len(scores) == 0:
                return debt_array
            
            """
            maxCreditor, maxCredit = maxEntry(scores)

	        maxDebtor, maxDebt = minEntry(scores)

            if maxCredit == -maxDebt:
                debts = append(existingDebts, [3]int64{maxDebtor, maxCreditor, -maxDebt})
                delete(scores, maxDebtor)
                delete(scores, maxCreditor)
            elif maxCredit > -maxDebt:
                debts = append(existingDebts, [3]int64{maxDebtor, maxCreditor, -maxDebt})
                delete(scores, maxDebtor)
                scores[maxCreditor] += maxDebt
            elif maxCredit < -maxDebt:
                debts = append(existingDebts, [3]int64{maxDebtor, maxCreditor, maxCredit})
                delete(scores, maxCreditor)
                scores[maxDebtor] += maxCredit
            """
            
            return greedy(updated_scores, new_debt_array)

        return greedy(transactions, debts)


def scores(transactions):
    scores = defaultdict(int)

    for f, t, a in transactions:
        scores[f] -= a
        scores[t] += a

    return scores

"""
def maxEntry(m):
	# find first element
	for i = range m: 
		index = i
		value = m[index]
		break
	
	# obtain max element
	for x := range m :
		if m[x] > value:
			index, value = x, m[x]
	return


def minEntry(m) :
	# find first element
	for i := range m {
		index = i
		value = m[index]
		break
	}
	# obtain min element
	for x := range m {
		if m[x] < value {
			index, value = x, m[x]
		}
	}
	return
"""


if __name__ == "__main__":
    print("splitwise algorithm in python")

    # instantiate class instance
    solution_instance = Solution

    # Measure csv load time
    start = time.perf_counter_ns()
    data = solution_instance().load("../test_data/input.csv")
    print(f"completed csv load execution in {(time.perf_counter_ns() - start)/1000} microseconds")

    # Measure algorithm time
    min_transfer_start = time.perf_counter_ns()
    m = solution_instance().min_transfers(data)
    print(f"completed min_transfer execution in {(time.perf_counter_ns() -  min_transfer_start)/1000} microseconds")

    print("minimum transfers were", m)