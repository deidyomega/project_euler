def test_num(test):
    test_set = [11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
    for num in test_set:
        if test % num != 0:
            return False
    return True

for x in range(2520, 10000000000, 2520):
    if test_num(x):
        print("Answer 5:", x)
        break
