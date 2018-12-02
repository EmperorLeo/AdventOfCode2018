using System;
using System.Collections.Generic;
using System.IO;

namespace adventofcode
{
    class Program
    {
        static void Main(string[] args)
        {
            IProblemRunner runner = new DayTwo();
            Console.WriteLine(runner.InvokeSilver());
            Console.WriteLine(runner.InvokeGold());
        }
    }
}
