def test_num(test_number)
    test_set = [11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
        for num in test_set
            if test_number % num != 0
                return false
            end
        end
    return true
end

for x in (2520..10000000000).step(2520) do
    if test_num(x)
        puts("Problem 5: " + String(x))
        break
    end
end