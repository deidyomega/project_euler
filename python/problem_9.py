from __future__ import print_function
stop = False
for c in range(100,1000):
    for b in range(100,c):
        for a in range(100,b):
            if a + b + c == 1000:
                if a**2 + b**2 == c**2:
                    print("A:", a, "B:", b, "C:", c)
                    print("Problem 9:", a*b*c)
                    stop = True
            if stop: break
        if stop: break
    if stop: break