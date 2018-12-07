using System;
using System.Collections.Generic;
using System.Collections.Specialized;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using adventofcode;

namespace csharp
{
    public class DaySeven : IProblemRunner
    {
        struct DirectedEdge
        {
            public string Start, End;

            public DirectedEdge(string start, string end)
            {
                Start = start;
                End = end;
            }
        }

        private readonly IEnumerable<DirectedEdge> _edges;
        private SortedList<string, string> _sortedList;
        private Dictionary<string, ICollection<string>> _graph;
        private Dictionary<string, ICollection<string>> _requirementsGraph;
        private int _numVertices;

        public DaySeven()
        {
            var regex = new Regex(@"Step (.) must be finished before step (.) can begin\.");
            _edges = File.ReadAllLines("../input/day7.txt").Select(l =>
            {
                var matches = regex.Match(l).Groups;
                return new DirectedEdge(matches[1].Value, matches[2].Value);
            });
        }

        public string InvokeSilver()
        {
            Setup();
            var result = "";
            while (_sortedList.Count > 0)
            {
                var vertex = _sortedList.ElementAt(0).Value;
                _sortedList.RemoveAt(0);

                result += vertex;

                if (_graph.ContainsKey(vertex))
                {
                    var children = _graph[vertex];
                    foreach (var child in children)
                    {
                        var requirements = _requirementsGraph[child];
                        requirements.Remove(vertex);

                        if (!requirements.Any())
                        {
                            _sortedList.Add(child, child);
                        }
                    }
                }
            }

            return result;
        }

        public string InvokeGold()
        {
            Setup();
            const int maxWorkers = 5;
            var workers = 0;
            var timer = 0;
            var result = "";
            var doneDict = new Dictionary<int, SortedList<string, string>>();
            while (result.Length < _numVertices)
            {
                if (doneDict.ContainsKey(timer))
                {
                    foreach (var item in doneDict[timer])
                    {
                        var vertex = item.Value;
                        result += vertex;

                        if (_graph.ContainsKey(vertex))
                        {
                            var children = _graph[vertex];
                            foreach (var child in children)
                            {
                                var requirements = _requirementsGraph[child];
                                requirements.Remove(vertex);

                                if (!requirements.Any())
                                {
                                    _sortedList.Add(child, child);
                                }
                            }
                        }

                        workers--;
                    }
                }

                while (_sortedList.Count > 0 && workers < maxWorkers)
                {
                    workers++;

                    var vertex = _sortedList.ElementAt(0).Value;
                    _sortedList.RemoveAt(0);

                    var doneTime = timer + (int)vertex[0] - 4;
                    if (!doneDict.ContainsKey(doneTime))
                    {
                        doneDict.Add(doneTime, new SortedList<string, string>());
                    }
                    doneDict[doneTime].Add(vertex, vertex);
                }
                timer++;
            }
            // Note: have to subtract one second, because the timer increments an extra time after everything is done.
            return $"Took {timer - 1} seconds.";
        }

        private void Setup()
        {
            var seenSet = new HashSet<string>();
            var beginningSet = new HashSet<string>();
            _graph = new Dictionary<string, ICollection<string>>();
            _requirementsGraph = new Dictionary<string, ICollection<string>>();
            foreach (var edge in _edges)
            {
                if (!_graph.ContainsKey(edge.Start))
                {
                    _graph.Add(edge.Start, new HashSet<string>());
                }
                _graph[edge.Start].Add(edge.End);


                if (seenSet.Add(edge.Start))
                {
                    beginningSet.Add(edge.Start);
                }

                if (!_requirementsGraph.ContainsKey(edge.End))
                {
                    _requirementsGraph.Add(edge.End, new HashSet<string>());
                }
                _requirementsGraph[edge.End].Add(edge.Start);

                seenSet.Add(edge.End);
                beginningSet.Remove(edge.End);
            }

            _numVertices = seenSet.Count();

            _sortedList = new SortedList<string, string>();
            beginningSet.ToList().ForEach(x => _sortedList.Add(x, x));
        }
    }
}