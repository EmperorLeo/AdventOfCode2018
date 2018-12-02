using System;
using System.IO;
using System.Linq;

namespace adventofcode
{
    public class DayTwo : IProblemRunner
    {
        public string InvokeSilver()
        {
            var lines = File.ReadAllLines("./day2.txt");
            var twoMatch = 0;
            var threeMatch = 0;
            foreach (var box in lines)
            {
                var grouping = box.ToCharArray().GroupBy(x => x);
                if (grouping.Any(g => g.Count() == 3))
                {
                    threeMatch++;
                }
                if (grouping.Any(g => g.Count() == 2))
                {
                    twoMatch++;
                }
            }
            return $"Checksum = {twoMatch * threeMatch} ({twoMatch} * {threeMatch}).";
        }

        public string InvokeGold()
        {
            var lines = File.ReadAllLines("./day2.txt");
            foreach (var box1 in lines)
            {
                foreach (var box2 in lines)
                {
                    if (box1 != box2)
                    {
                        var zipped = box1.Zip(box2, (char1, char2) => char1 == char2).ToList();
                        if (zipped.Count(x => !x) == 1)
                        {
                            return box1.Remove(zipped.IndexOf(false), 1);
                        }
                    }
                }
            }

            return "Box not found.";
        }
    }
}