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


def greedy(scores, debt_array):
    if len(scores) == 0:
        return debt_array
        
    max_creditor, max_credit = max_entry(scores)
    max_debtor, max_debt = min_entry(scores)

    new_debt_array = debt_array.copy()

    if max_credit == -max_debt:
        new_debt_array.append([max_debtor, max_creditor, -max_debt])
        del scores[max_debtor]
        del scores[max_creditor]
    elif max_credit  > -max_debt:
        new_debt_array.append([max_debtor, max_creditor, -max_debt])
        del scores[max_debtor]
        scores[max_creditor] += max_debt
    elif max_credit  < -max_debt:
        new_debt_array.append([max_debtor, max_creditor, max_credit])
        del scores[max_creditor]
        scores[max_debtor] += max_credit

    updated_scores = scores
    return greedy(updated_scores, new_debt_array)

def scores(transactions):
    scores = defaultdict(int)

    for f, t, a in transactions:
        scores[f] -= a
        scores[t] += a

    return scores


def max_entry(m):
	# find first element
	for i in m: 
		index = i
		value = m[index]
		break
	
	# obtain max element
	for x in m:
		if m[x] > value:
			index, value = x, m[x]
	return index, value


def min_entry(m):
	# find first element
	for i in m :
		index = i
		value = m[index]
		break
	# obtain min element
	for x in m :
		if m[x] < value :
			index, value = x, m[x]
	return index, value



"""
Solution

Class containing minimum transaction algorithm and csv decoding 
"""
class Solution: 

    def __init__(self):      
        # do nothing
        return

    def read_file(self, input_csv):
        return np.genfromtxt(input_csv, delimiter=',')

    def simplify_debts(self, transactions):
        balances = scores(transactions)
        debts = list()
        return greedy(balances, debts)




if __name__ == "__main__":
    print("splitwise algorithm in python")

    # instantiate class instance
    solution_instance = Solution

    # Measure csv load time
    start = time.perf_counter_ns()
    data = solution_instance().read_file("../../../test_data/input.csv")
    print(f"read csv file in {(time.perf_counter_ns() - start)/1000} microseconds")

    # Measure algorithm time
    restart = time.perf_counter_ns()
    m = solution_instance().simplify_debts(data)
    print(f"completed simplify_debts execution in {(time.perf_counter_ns() -  restart)/1000} microseconds, transactions {len(m)}")

    print("simplified debts", m)