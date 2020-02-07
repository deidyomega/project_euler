program  problem_1
    integer i
    integer total
    total = 0
    do i=0,999
        if (MODULO(i, 3) == 0 .OR. MODULO(i, 5) == 0) then
            total = total + i
        end if
    end  do
    print *,"Problem 1:", total
end program problem_1