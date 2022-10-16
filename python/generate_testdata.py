# generate random integer values
import csv
from itertools import count
import numpy  
from numpy.random import seed
from numpy.random import randint
from numpy import *



# seed random number generator
seed(2500)

payments = 1000
players = 21
max_payable = 100

counter = 0
with open('input.csv', 'w', encoding='UTF8') as f:
    writer = csv.writer(f)
    while counter < payments: 
        # generate some integers
        transactors = randint(1, players+1, 2)
        if transactors[0] == transactors[1]:
            continue
        amt = randint(1, max_payable, 1)
        transaction_data = numpy.append(transactors, amt)
        # write the data
        writer.writerow(transaction_data)
        counter += 1