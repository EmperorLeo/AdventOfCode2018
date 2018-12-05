using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using adventofcode;

namespace csharp
{
    public class DayFive : IProblemRunner
    {
        private readonly string _input;

        public DayFive()
        {
            _input = File.ReadAllText("../input/day5.txt");
        }

        public string InvokeGold()
        {

            var polymerLengthsPerIgnored = "abcdefghijklmnopqrstuvwxyz".Select(c =>
            {
                return GetResultingPolymerLength(c);
            });
            return $"Shortest possible: {polymerLengthsPerIgnored.Min()}";
        }

        public string InvokeSilver()
        {
            return $"Length: {GetResultingPolymerLength(null)}";
        }

        private int GetResultingPolymerLength(char? ignore)
        {
            var stack = new Stack<char>();
            foreach (var polymer in _input)
            {
                if (ignore.HasValue && Char.ToLower(ignore.Value) == Char.ToLower(polymer))
                {
                    continue;
                }
                if (stack.Count == 0)
                {
                    stack.Push(polymer);
                    continue;
                }
                var last = stack.Peek();
                if ((Char.IsLower(last) && Char.IsUpper(polymer) || Char.IsUpper(last) && Char.IsLower(polymer)) && Char.ToLower(last) == Char.ToLower(polymer))
                {
                    stack.Pop();
                }
                else
                {
                    stack.Push(polymer);
                }
            }
            return stack.Count;
        }
    }
}