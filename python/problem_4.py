from __future__ import print_function

answer = 0
for x in range(900, 1000):
    for y in range(990, 1000):
        tmp = x * y
        stmp = str(tmp)
        if stmp == stmp[::-1]:
            answer = tmp

print("Answer 4:", answer)

