using System;
using System.Collections.Generic;
using System.IO;
using csharp;

namespace adventofcode
{
    class Program
    {
        static void Main(string[] args)
        {
            IProblemRunner runner = new DayFive();
            Console.WriteLine(runner.InvokeSilver());
            Console.WriteLine(runner.InvokeGold());
        }
    }
}
