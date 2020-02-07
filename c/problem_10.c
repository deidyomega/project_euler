#include <stdio.h>
#include <stdlib.h>

// Size of array / number of elements
static int ARR_SIZE = 2000000;
// Removes from arr any number that is a multiple of num
int *removal(int *arr, unsigned long num);

int main()
{
    // Create an array of ARR_SIZE, in the heap
    int *primes = malloc(ARR_SIZE*sizeof(long));
    unsigned long total = 0;

    // Fill the array with values 0 .. ARR_SIZE
    for(unsigned long x = 2; x < ARR_SIZE; x++) {
        primes[x] = x;
    }
    
    // Using: Sieve of Eratosthenes
    for(unsigned long x=2; x < ARR_SIZE; x++) { // iterate over the list
        // remove any number that is a multiple of the number
        // encountered, all that remain are primes
        removal(primes, primes[x]);
        total += primes[x];
    }
    // Get the total of the array, now that non primes are gone
    printf("Answer: %lu\n", total);

    // Free the array from memory
    free(primes);
    return 0;
}

int *removal(int *arr, unsigned long num) {
    // num should always be either zero, or a prime

    // skip zero'd values, as they've been eliminated
    if(num != 0){
        // iterate over array, incrementing by num
        // which gives the multiples of said prime number
        for(unsigned long x = num+num; x < ARR_SIZE; x += num){
            // Set said multiple to zero, effectively eliminating them from set
            arr[x] = 0;
        }
    }
    return 0;
}
