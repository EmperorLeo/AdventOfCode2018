using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using adventofcode;

namespace csharp
{
    public class DaySix : IProblemRunner
    {
        struct Coordinate
        {
            public int Id, X, Y;

            public Coordinate(int id, int x, int y)
            {
                Id = id;
                X = x;
                Y = y;
            }

            public int ManhattanDistance(int x, int y)
            {
                return Math.Abs(X - x) + Math.Abs(Y - y);
            }
        }

        private IEnumerable<Coordinate> points;
        private int maxX;
        private int maxY;

        public string InvokeGold()
        {
            var areaSize = 0;
            for (var x = 0; x <= maxX; x++)
            {
                for (var y = 0; y <= maxY; y++)
                {
                    if (points.Sum(p => p.ManhattanDistance(x, y)) < 10000)
                    {
                        areaSize++;
                    }
                }
            }

            return $"Area Size = {areaSize}";
        }

        public string InvokeSilver()
        {
            points = File.ReadAllLines("../input/day6.txt").Select((x, i) =>
            {
                var split = x.Split(", ");
                return new Coordinate(i, int.Parse(split[0]), int.Parse(split[1]));
            });
            maxY = points.Max(p => p.Y);
            maxX = points.Max(p => p.X);

            var map = new int[maxY + 1][];
            for (var i = 0; i < map.Length; i ++)
            {
                map[i] = new int[maxX + 1];
            }

            var ties = 0;

            for (var x = 0; x <= maxX; x++)
            {
                for (var y = 0; y <= maxY; y++)
                {
                    var isTie = false;
                    var winnerId = -1;
                    var closestDist = int.MaxValue;
                    foreach (var point in points)
                    {
                        var dist = point.ManhattanDistance(x, y);
                        if (closestDist == dist)
                        {
                            isTie = true;
                        }
                        else if (dist < closestDist)
                        {
                            isTie = false;
                            closestDist = dist;
                            winnerId = point.Id;
                        }
                    }

                    if (isTie)
                    {
                        ties++;
                    }

                    map[y][x] = isTie ? -1 : winnerId;
                }
            }

            var totalAreaMap = new Dictionary<int, int>();
            var disqualifiedSet = new HashSet<int>();
            for (var x = 0; x <= maxX; x++)
            {
                for (var y = 0; y <= maxY; y++)
                {
                    var id = map[y][x];
                    if (id == -1)
                    {
                        continue;
                    }

                    if (x == 0 || y == 0 || x == maxX || y == maxY)
                    {
                        disqualifiedSet.Add(id);
                    }

                    if (!totalAreaMap.ContainsKey(id))
                    {
                        totalAreaMap.Add(id, 0);
                    }

                    totalAreaMap[id]++;
                }
            }

            var total = totalAreaMap.Values.Sum();

            disqualifiedSet.ToList().ForEach(x => totalAreaMap.Remove(x));
            return $"Largest non-infinite area is {totalAreaMap.Values.Max()}";
        }
    }
}