using System.IO;
using System.Linq;

namespace adventofcode
{
    public class DayThree : IProblemRunner
    {
        private Claim[] claims;
        private int[][] fabricAreaClaims;

        public string InvokeGold()
        {
            // Assumes silver has been run
            foreach (var claim in claims)
            {
                var conflicted = false;
                for (var x = claim.x; x < claim.x + claim.width; x++)
                {
                    for (var y = claim.y; y < claim.y + claim.height; y++)
                    {
                        if (fabricAreaClaims[y][x] > 1)
                        {
                            conflicted = true;
                            break;
                        }
                    }

                    if (conflicted)
                    {
                        break;
                    }
                }

                if (!conflicted)
                {
                    return $"Here is the cloth ID: {claim.number}.";
                }
            }

            return "Advent of Code Santa lied! All cloth pieces conflict.";
        }

        public string InvokeSilver()
        {
            var claimStrs = File.ReadAllLines("../input/day3.txt");
            claims = new Claim[claimStrs.Length];
            var maxClothWidth = 0;
            var maxClothHeight = 0;
            for (var i = 0; i < claimStrs.Length; i++)
            {
                var claimStr = claimStrs[i];
                var split = claimStr.Split(' ');
                var posString = split[2].Split(',');
                var dimString = split[3].Split('x');
                var claim = new Claim(
                    int.Parse(split[0].Substring(1)),
                    int.Parse(posString[0]),
                    int.Parse(posString[1].Substring(0, posString[1].Length - 1)),
                    int.Parse(dimString[0]),
                    int.Parse(dimString[1]));
                claims[i] = claim;
                
                if (claim.x + claim.width > maxClothWidth)
                {
                    maxClothWidth = claim.x + claim.width;
                }

                if (claim.y + claim.height > maxClothHeight)
                {
                    maxClothHeight = claim.y + claim.height;
                }
            }

            fabricAreaClaims = new int[maxClothHeight][];
            for (var i = 0; i < fabricAreaClaims.Length; i ++)
            {
                fabricAreaClaims[i] = new int[maxClothWidth];
            }

            foreach (var claim in claims)
            {
                for (var x = claim.x; x < claim.x + claim.width; x++)
                {
                    for (var y = claim.y; y < claim.y + claim.height; y++)
                    {
                        fabricAreaClaims[y][x]++;
                    }
                }
            }

            var totalConflictSquareInches = 0;
            foreach (var row in fabricAreaClaims)
            {
                foreach (var squareInch in row)
                {
                    if (squareInch > 1)
                    {
                        totalConflictSquareInches++;
                    }
                }
            }

            return $"Total conflicted square inches: {totalConflictSquareInches}";
        }

        struct Claim
        {
            public int number, x, y, width, height;

            public Claim(int num, int p1, int p2, int w, int h)
            {
                number = num;
                x = p1;
                y = p2;
                width = w;
                height = h;
            }
        }
    }
}