def gen_primes():
    current_number = 2
    primes = []
    while True:
        is_prime = True
        for prime in primes:
            if prime*prime > current_number: ## This is saying it's a prime if prime * prime > current_testing_number
                # print("Caught here")
                break # I'd love to see the original proof to this
            if current_number % prime == 0:  # if divisable by a known prime, then it's not a prime
                is_prime = False
                break
        if is_prime:
            primes.append(current_number)
            yield current_number
        current_number += 1

# primes: 2, 3, 5, 7, 11
# not p : 4, 6, 8, 9, 10

x = 2
for p in gen_primes():
    if x > 10001:
        print(p)
        break
    x += 1