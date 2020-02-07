from __future__ import print_function
ARR_SIZE = 2000000

def removal(arr, num):
    # num should always be either zero, or a prime

    # skip zero'd values, as they've been eliminated
    if num != 0:
        # iterate over array, incrementing by num
        # which gives the multiples of said prime number
        for x in range(num+num, ARR_SIZE, num):
            arr[x] = 0
    return arr

arr = [x for x in range(ARR_SIZE)]
arr[1] = 0
total = 0
for x in range(2,ARR_SIZE):
    arr = removal(arr,arr[x])
    total += arr[x]

print("Problem 10:", total)
