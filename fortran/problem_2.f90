program  problem_2
    integer a
    integer b
    integer c
    integer total
    a = 1
    b = 1
    c = 0
    total = 0
    do while (c < 4000000)
        c = a + b
        a = b
        b = c

        if (MODULO(c, 2) == 0) then
            total = total + c
        end if
    end do
    print *,"Problem 2:", total
end program problem_2
