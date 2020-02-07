SUBROUTINE removal (num)
    IMPLICIT  NONE
    integer(kind=8) num
    integer i
    integer, parameter :: ARR_SIZE = 2000000
    COMMON primes
    integer(kind=8), DIMENSION(ARR_SIZE) :: primes

    if (num /= 0) then
        do i=num+num, ARR_SIZE, num
            primes(i) = 0
        end do
    end if
END SUBROUTINE removal

program  problem_10
    integer(kind=8) i
    integer, parameter :: ARR_SIZE = 2000000
    integer(kind=8), DIMENSION(ARR_SIZE) :: primes
    integer(kind=8) total
    integer(kind=8) rez
    integer(kind=8) lg_sum
    COMMON primes
    total = 0

    ! Set primes to 1, 2, .. ARR_SIZE
    primes = (/ (i, i = 1, ARR_SIZE) /)
    primes(1) = 0
    ! Using: Sieve of Eratosthenes
    do i=1, ARR_SIZE
        CALL removal(primes(i))
        total = total + primes(i)
    end do

    print *,"Problem 10:", total


end program problem_10

