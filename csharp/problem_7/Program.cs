using System;
using System.Collections.Generic;

namespace problem_7
{
    class Program
    {
        static void Main(string[] args)
        {
            
            IEnumerable<int> elements = gen_primes(); 
            int x = 2;
            foreach (int prime in elements)
            {
                if (x > 10001) {
                    Console.WriteLine("PRoblem 7: " + prime);
                    break;
                }
                x++;
            }

        }

        public static IEnumerable<int> gen_primes()
        {
            int current_number = 2;
            List<int> primes = new List<int>();
            bool is_prime;

            while (true)
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
