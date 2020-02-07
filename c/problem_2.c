#include <stdio.h>

int main() 
{
    unsigned long int total; // It's a really large number...
    int a = 1;
    int b = 1;
    int c = 0;
    while(c < 4000000) {
        c = a + b;
        a = b;
        b = c;
        if (c % 2 == 0) {
            total += c;
        }
    }
    printf("Problem 2: %lu\n", total);
    return 0;
}