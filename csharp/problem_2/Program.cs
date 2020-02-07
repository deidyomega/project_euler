using System;

namespace problem_2
{
    class Program
    {
        static void Main(string[] args)
        {
            int a = 1;
            int b = 1;
            int c = 0;
            int total = 0;

            while (c < 4000000)
            {
                c = a + b;
                a = b;
                b = c;
                if (c % 2 == 0) {
                    total += c;
                } 
            }

            Console.WriteLine("Problem 2: " + total);
        }
    }
}