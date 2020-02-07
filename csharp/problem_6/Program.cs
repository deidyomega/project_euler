using System;

namespace problem_6
{
    class Program
    {
        static void Main(string[] args)
        {
            int sum_total = 0;
            int square_total = 0;

            for (int i = 0; i <= 100; i++)
            {
                sum_total += (int)Math.Pow(i, 2);
                square_total += i;
            }
            square_total = (int)Math.Pow(square_total, 2);

            Console.WriteLine("Problem 6: " + (square_total - sum_total));
        }
    }
}
