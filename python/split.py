from collections import defaultdict
from array import *
import pandas as pd
"""
Template code to be benchmarked
"""
import csv
import datetime

"""
Solution

Class containing minimum transaction algorithm
"""
class Solution: 

    def __init__(self, input_csv):      
        self.input_csv = input_csv

    def decode_transactions(self):
        input_data = pd.read_csv(self.input_csv) 
        return input_data


    def min_transfers(self, transactions) -> int:
        score = defaultdict(int)

        for f, t, a in transactions:
            score[f] -= a
            score[t] += a
        
        positives = [v for v in score.values() if v > 0]
        negatives = [v for v in score.values() if v < 0]

        def recurse(positives, negatives):
            negative = negatives[0]

            count = inf
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

def main():
    data = Solution("input.csv").decode_transactions()
    m = Solution("input.csv").min_transfers(data)
    print ("minimum transfers were", m)


if __name__ == "__main__":
    main()