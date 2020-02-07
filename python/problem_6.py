from __future__ import print_function

sum_total = 0
square_total = 0
for x in range(101):
    sum_total += x ** 2
    square_total += x

square_total = square_total ** 2

print("Problem 6:", square_total - sum_total)
