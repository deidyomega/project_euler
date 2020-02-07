answer = 0
for x in 900 .. 1000 do
    for y in 990 .. 1000 do
        tmp = x * y
        stmp = String(tmp)
        if stmp == stmp.reverse
            answer = tmp
        end
    end
end

puts("Answer 4: " + String(answer))