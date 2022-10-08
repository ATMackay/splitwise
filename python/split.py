from collections import defaultdict
from array import *
import pandas as pd
import numpy as np
import math
import csv
import datetime
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

    def __init__(self, input_csv):      
        self.input_csv = input_csv
        self.data = np.loadtxt(self.input_csv, dtype=int)



    def min_transfers(self) -> int:
        score = defaultdict(int)

        for f, t, a in self.data:
            score[f] -= a
            score[t] += a
        
        positives = [v for v in score.values() if v > 0]
        negatives = [v for v in score.values() if v < 0]

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

def main():
    m = Solution("input.csv").min_transfers()
    print ("minimum transfers were", m)


if __name__ == "__main__":
    main()