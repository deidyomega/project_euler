from __future__ import print_function

def removal(arr, num):
        # num should always be either zero, or a prime
        # skip zero'd values, as they've been eliminated
        if num != 0:
            # iterate over array, incrementing by num
            # which gives the multiples of said prime number
            for x in range(num+num, len(arr), num):
                arr[x] = 0
        return arr
        
def prime_list(ARR_SIZE):
    arr = [x for x in range(ARR_SIZE)]
    arr[1] = 0

    for x in range(2,ARR_SIZE):
        arr = removal(arr,arr[x])

    ## remove zeros from list
    while x < len(arr):
        if arr[x] == 0:
            del arr[x]
            x -= 4
        x += 1
    return arr

def prime_factorization(num, primes, p_list):
    for prime in primes:
        if num % prime == 0:
            p_list.append(prime)
            p_list = prime_factorization(num / prime, primes, p_list)
            break
    return p_list

p_list = prime_factorization(600851475143, prime_list(10000), [])
print(p_list)
print ("Problem 3:", max(p_list))
