#include <stdio.h>
#include <math.h>

int main() {
    int stop = 0;
    int a, b, c;
    for(c = 100; c < 1000; c++) {
        for(b = 100; b < c; b++) {
            for(a = 100; a < b; a++) {
                if (a + b + c == 1000) {
                    if(pow(a,2) + pow(b,2) == pow(c,2)) {
                        printf("A: %i, B: %i, C: %i\n", a,b,c);
                        printf("Problem 9: %i\n", a*b*c);
                        stop = 1;
                    }
                }
                if (stop) {break;}
            }
            if (stop) {break;}
        }
        if (stop) {break;}
    }
    return 0;
}