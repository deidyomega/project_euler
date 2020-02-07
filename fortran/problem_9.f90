program problem_9
    integer a
    integer b
    integer c
    integer total
    total = 0
    do c=100,999
        do b=100,c
            do a=100,b
                if (a + b + c == 1000) then
                    if (a**2 + b**2 == c**2) then
                        print *,"A:", a, "B:", b, "C:", c
                        print *,"Problem 9:", a*b*c
                        stop ! We found the correct solution, no need to keep testing
                    end if
                end if
            end  do
        end  do
    end  do
end program problem_9
