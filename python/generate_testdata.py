# generate random integer values
import csv
from itertools import count
import numpy  
from numpy.random import seed
from numpy.random import randint
from numpy import *



# seed random number generator
seed(2500)

counter = 0
with open('input.csv', 'w', encoding='UTF8') as f:
    writer = csv.writer(f)
    while counter < 1000: 
        # generate some integers
        values = randint(1, 12, 2)
        amt = randint(1, 1000, 1)
        data = numpy.append(values, amt)
        # write the data
        writer.writerow(data)
        counter += 1