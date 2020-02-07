using System;

namespace problem_5
{
    class Program
    {
        static void Main(string[] args)
        {
            uint i;
            for (i = 2520; true; i += 2520)
            {
                if (test_num(i)) {
                    Console.WriteLine("Problem 5: " + i);
                    break;
                }
                
            }
        }

        static bool test_num(uint test) {
            int[] test_set = new int[10] {11, 12, 13, 14, 15, 16, 17, 18, 19, 20};
            foreach (int num in test_set)
            {
                if (test % num != 0) {
                    return false;
                }
            }
            return true;
        }
    }
}
