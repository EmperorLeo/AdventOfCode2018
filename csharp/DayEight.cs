using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using adventofcode;

namespace csharp
{
    public class DayEight : IProblemRunner
    {
        private readonly int[] _data;
        public DayEight()
        {
            _data = File.ReadAllText("../input/day8.txt")
                .Split(" ")
                .Select(int.Parse)
                .ToArray();
        }

        public string InvokeSilver()
        {
            var tree = ParseNodes(_data, null);
            return $"Sum of all metadata entries = {tree.Sum(t => t.MetadataSum)}";
        }

        public string InvokeGold()
        {
            var tree = ParseNodes(_data, null);
            return $"Root value = {tree[0].Value}";
        }

        private Node[] ParseNodes(int[] stuff, int? numNodes)
        {
            var incrementor = 0;
            var currentNode = 0;
            var returnNodes = new List<Node>();
            while (incrementor < stuff.Length && (!numNodes.HasValue || currentNode < numNodes))
            {
                var numChildren = stuff[incrementor];
                var numMetadata = stuff[incrementor + 1];
                var header = new Header(numChildren, numMetadata);
                var nodeSize = 2 + numMetadata;

                // BASE CASE
                if (numChildren == 0)
                {
                    var metadata = stuff.Skip(incrementor + 2).Take(numMetadata);
                    var value = metadata.Sum();

                    returnNodes.Add(new Node
                    {
                        Header = header,
                        Nodes = new Node[0],
                        Size = nodeSize,
                        Metadata = metadata,
                        MetadataSum = value,
                        Value = value
                    });
                }
                else
                {
                    // Recursion!! Yay....
                    var nodes = ParseNodes(stuff.Skip(incrementor + 2).ToArray(), numChildren);
                    var childrenSize = nodes.Sum(n => n.Size);
                    nodeSize += childrenSize;
                    var metadata = stuff.Skip(incrementor + childrenSize + 2).Take(numMetadata);

                    var value = 0;
                    foreach (var metadatum in metadata)
                    {
                        if (metadatum > 0 && metadatum <= nodes.Length)
                        {
                            value += nodes[metadatum - 1].Value;
                        }
                    }

                    returnNodes.Add(new Node
                    {
                        Header = header,
                        Nodes = nodes,
                        Size = nodeSize,
                        Metadata = metadata,
                        MetadataSum = nodes.Sum(n => n.MetadataSum) + metadata.Sum(),
                        Value = value
                    });
                }

                incrementor += nodeSize;
                currentNode++;
            }

            return returnNodes.ToArray();
        }

        class Node
        {
            // Useless variable so far
            public Header Header { get; set; }
            public Node[] Nodes { get; set; }
            public int Size { get; set; }
            public IEnumerable<int> Metadata { get; set; }
            public int MetadataSum { get; set; }
            public int Value { get; set; }
        }

        struct Header
        {
            public int NumChildren, NumMetadata;

            public Header(int numChildren, int numMetadata)
            {
                NumChildren = numChildren;
                NumMetadata = numMetadata;
            }
        }
    }
}