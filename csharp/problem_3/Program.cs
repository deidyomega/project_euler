using System;
using System.Collections.Generic;

namespace problem_3
{
    class Program
    {
        static void Main(string[] args)
        {
            List<int> p_list = prime_factorization(600851475143, new List<int>());
            Console.WriteLine("Problem 3: " + p_list[ p_list.Count-1 ]);
        }

        public static List<int> prime_factorization(long num, List<int> p_list) {
            IEnumerable<int> primes = gen_primes(num + 1); 
                        
            foreach (int prime in primes)
            {
                if (num % prime == 0) {
                    p_list.Add(prime);
                    p_list = prime_factorization(num / prime, p_list);
                    break;
                }
            }
            return p_list;
        }

        public static IEnumerable<int> gen_primes(long limit)
        {
            int current_number = 2;
            List<int> primes = new List<int>();
            bool is_prime;

            while (current_number < limit)
            {
                is_prime = true;
                foreach (int prime in primes)
                {
                    if (prime * prime > current_number) {
                        break;
                    }
                    if (current_number % prime == 0) {
                        is_prime = false;
                        break;
                    }
                }
                if (is_prime) {
                    primes.Add(current_number);
                    yield return current_number;
                }
                current_number += 1;
            }
        }
    }
}
