total = 0
for x in 0..1000
    if x % 3 == 0 or x % 5 == 0
        total += x
    end
end
puts("Problem 1: " + String(total))