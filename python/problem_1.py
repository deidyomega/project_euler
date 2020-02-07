from __future__ import print_function

total = 0
for x in range(1000):
    if x % 3 == 0 or x % 5 == 0:
        total += x

print("Problem 1: " + str(total))