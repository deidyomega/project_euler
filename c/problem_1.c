#include <stdio.h>

int main() 
{
    int total;
    for(int x = 0; x < 1000; x++) {
        if (x % 3 == 0 || x % 5 == 0) {
            total += x;
        }
    }
    printf("Problem 1: %i\n", total);
    return 0;
}