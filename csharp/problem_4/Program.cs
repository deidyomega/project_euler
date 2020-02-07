using System;

namespace problem_4
{
    class Program
    {
        static void Main(string[] args)
        {
            int largest = 0;
            string stmp;
            int tmp;
            int x;
            int y;
            
            for (x = 900; x < 1000; x++)
            {
                for (y = 900; y < 1000; y++)
                {
                    tmp = x * y;
                    stmp = tmp.ToString();
                    if (stmp == Reverse(stmp)) {
                        if (tmp > largest) {
                            largest = tmp;
                        }
                        
                    }

                }                
            }
            Console.WriteLine("Problem 4: " + largest);
        }
        public static string Reverse( string s )
        {
            char[] charArray = s.ToCharArray();
            Array.Reverse( charArray );
            return new string( charArray );
        }
    }
}
