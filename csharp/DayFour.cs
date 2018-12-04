using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using adventofcode;

namespace csharp
{
    public class DayFour : IProblemRunner
    {
        struct GuardTimestamp
        {
            public DateTime timestamp;
            public int? guardId;
            public bool? isArrival, awake;
             
            public GuardTimestamp(DateTime ts, int gid, bool arrival)
            {
                timestamp = ts;
                guardId = gid;
                isArrival = arrival;
                awake = null;
            }

            public GuardTimestamp(DateTime ts, bool a)
            {
                timestamp = ts;
                guardId = null;
                isArrival = null;
                awake = a;
            }
        }

        private Dictionary<int, int[]> guardSleepMap;

        public string InvokeGold()
        {
            var mostPredictableGuard = 0;
            var mostPredictableGuardMinutes = 0;
            foreach (var guardSleepMinutes in guardSleepMap.Select(x => new { GuardId = x.Key, MaxMinutes = x.Value.Max() }))
            {
                if (guardSleepMinutes.MaxMinutes > mostPredictableGuardMinutes)
                {
                    mostPredictableGuard = guardSleepMinutes.GuardId;
                    mostPredictableGuardMinutes = guardSleepMinutes.MaxMinutes;
                }
            }

            var mostPredictableGuardsSleepiestMinute = guardSleepMap[mostPredictableGuard].ToList().IndexOf(mostPredictableGuardMinutes);

            return $"Most Predictable Guard * minutes {mostPredictableGuard} * {mostPredictableGuardsSleepiestMinute} = {mostPredictableGuard * mostPredictableGuardsSleepiestMinute}";
        }

        public string InvokeSilver()
        {
            var lines = File.ReadAllLines("../input/day4.txt");
            var guardShiftRegex = new Regex(@"\[(\d+-\d+-\d+)\ (\d{2}):(\d{2})\]\ Guard #(\d+) begins shift");
            var guardSleepRegex = new Regex(@"\[(\d+-\d+-\d+)\ (\d{2}):(\d{2})\]\ (falls asleep|wakes up)");

            var guardTimestamps = lines.Select(l =>
            {
                var shiftMatch = guardShiftRegex.Matches(l).FirstOrDefault();
                if (shiftMatch != null)
                {
                    var groups = shiftMatch.Groups.ToArray();
                    var date = groups[1].Value.Split('-');
                    Console.WriteLine("shift");
                    return new GuardTimestamp(
                        new DateTime(int.Parse(date[0]), int.Parse(date[1]), int.Parse(date[2]), int.Parse(groups[2].Value), int.Parse(groups[3].Value), 0),
                            int.Parse(groups[4].Value), true);
                }

                var sleepMatch = guardSleepRegex.Matches(l).FirstOrDefault();
                if (sleepMatch != null)
                {
                    var groups = sleepMatch.Groups.ToArray();
                    var date = groups[1].Value.Split('-');
                    Console.WriteLine("sleep");
                    return new GuardTimestamp(
                        new DateTime(int.Parse(date[0]), int.Parse(date[1]), int.Parse(date[2]), int.Parse(groups[2].Value), int.Parse(groups[3].Value), 0),
                            groups[4].Value == "wakes up");
                }

                throw new Exception("All of the lines should match my regex.");
            }).OrderBy(x => x.timestamp);

            guardSleepMap = new Dictionary<int, int[]>();
            var onGuard = 0;
            var fellAsleepAtMinute = 0;
            foreach (var ts in guardTimestamps)
            {
                if (ts.guardId.HasValue)
                {
                    onGuard = ts.guardId.Value;
                    if (!guardSleepMap.ContainsKey(onGuard))
                    {
                        guardSleepMap.Add(onGuard, new int[60]);
                    }
                }
                else
                {
                    if (ts.awake.Value)
                    {
                        // If waking up, mark the sleeping time period as asleep.
                        for (var i = fellAsleepAtMinute; i < ts.timestamp.Minute; i++)
                        {
                            guardSleepMap[onGuard][i]++;
                        }
                    }
                    else
                    {
                        fellAsleepAtMinute = ts.timestamp.Minute;
                    }
                }
            }

            var sleepiestGuard = 0;
            var sleepiestGuardMinutes = 0;
            foreach (var guardSleepMinutes in guardSleepMap.Select(x => new { GuardId = x.Key, TotalMinutes = x.Value.Sum() }))
            {
                if (guardSleepMinutes.TotalMinutes > sleepiestGuardMinutes)
                {
                    sleepiestGuard = guardSleepMinutes.GuardId;
                    sleepiestGuardMinutes = guardSleepMinutes.TotalMinutes;
                }
            }

            var maxMinutesAsleep = guardSleepMap[sleepiestGuard].Max();
            var minuteMostLikelyAsleep = guardSleepMap[sleepiestGuard].ToList().IndexOf(maxMinutesAsleep);

            return $"Sleepiest Guard * minutes {sleepiestGuard} * {minuteMostLikelyAsleep} = {sleepiestGuard * minuteMostLikelyAsleep}";
        }
    }
}