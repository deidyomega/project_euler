using System;

namespace problem_1
{
    class Program
    {
        static void Main(string[] args)
        {
            int total = 0;
            for (int i = 0; i < 1000; i++)
            {
                if (i % 3 == 0 || i % 5 == 0) {
                    total += i;
                }
            }
            Console.WriteLine("Problem 1: " + total);
            
        }
    }
}
