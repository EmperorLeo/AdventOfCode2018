using System.Collections.Generic;
using System.IO;

namespace adventofcode
{
    public class DayOne : IProblemRunner
    {
        public string InvokeSilver()
        {
            var lines = File.ReadAllLines("../input/day1.txt");
            var frequency = 0;
            foreach (var line in lines)
            {
                var firstChar = line[0];
                var rest = int.Parse(line.Substring(1));
                if (firstChar == '+')
                {
                    frequency += rest;
                }
                else
                {
                    frequency -= rest;
                }
            }
            return $"Frequency is {frequency}.";
        }
        
        public string InvokeGold()
        {
            var lines = File.ReadAllLines("../input/day1.txt");
            var frequency = 0;
            var frequencySet = new HashSet<int> { 0 };
            while (true) {
                foreach (var line in lines)
                {
                    var firstChar = line[0];
                    var rest = int.Parse(line.Substring(1));
                    if (firstChar == '+')
                    {
                        frequency += rest;
                    }
                    else
                    {
                        frequency -= rest;
                    }
                    
                    if (!frequencySet.Add(frequency))
                    {
                        return frequency.ToString();
                    }
                }
            }
        }
    }
}