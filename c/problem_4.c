#include <stdio.h>
#include <string.h>

int palindrome(int num); // determines if number is a palindrome
char *string_reverse(char *dst, const char *src); // reverses a string

int main() 
{
    unsigned int tmp;
    unsigned int answer = 0;
    for(int x = 900; x < 999; x++) {
        for(int y = 990; y < 999; y++) {
            tmp = x * y;
            if (palindrome(tmp) == 0) {
                if (tmp > answer) {
                    answer = tmp;
                }
            }
        }
    }
    printf("Problem 4: %i\n", answer);
    return 0;
}

int palindrome(int num){
    // 14 in the below was arbatray length greater than the length of the string.
    char src[14];
    char dst[14];
    snprintf(src, 14, "%d", num);  // converts number into string rep of number
    string_reverse(dst, src); // reverses string
    return strcmp(src, dst); // compares strings
}


char *string_reverse(char *dst, const char *src)
{
    if (src == NULL) return NULL;
 
    const char *src_start = src;
    char *dst_end = dst + strlen(src);
    *dst_end = '\0';
 
    while ((*--dst_end = *src_start++)) { ; }
 
    return dst;
}