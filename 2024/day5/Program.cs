var input = await File.ReadAllTextAsync("../../../input");

Console.WriteLine("Day 5 result 1: " + Part1(input));
Console.WriteLine("Day 5 result 2:" + Part2(input));
return;


int Part1(string input)
{
    (IEnumerable<(int left, int right)> rules, IEnumerable<int[]> updates) = ParseInput(input);
    var result = 0;
    foreach (var update in updates)
    {
        List<(int left, int right)> rulesForUpdate = rules.Where(m => update.Contains(m.left) && update.Contains(m.right)).ToList();
        List<int> ruleBasedList = GenerateRuleBasedList(rulesForUpdate);
        if (ruleBasedList.SequenceEqual(update)) result += update[(update.Length - 1) / 2];
    }

    return result;
}

int Part2(string input)
{
    (IEnumerable<(int left, int right)> rules, IEnumerable<int[]> updates) = ParseInput(input);

    return (from update in updates let rulesForUpdate = rules.Where(m => update.Contains(m.left) && update.Contains(m.right)).ToList() let ruleBasedList = GenerateRuleBasedList(rulesForUpdate) where !ruleBasedList.SequenceEqual(update) select ruleBasedList[(ruleBasedList.Count - 1) / 2]).Sum();
}

(IEnumerable<(int left, int right)> rules, IEnumerable<int[]> updates) ParseInput(string input)
{
    var rulesStr = input.Split("\n\n")[0];
    var updatesStr = input.Split("\n\n")[1];
    IEnumerable<(int left, int right)> rules = rulesStr
        .Split("\n")
        .Select(m => (Convert.ToInt32(m.Split("|")[0]), Convert.ToInt32(m.Split("|")[1])));

    IEnumerable<int[]> updates = updatesStr.Split("\n").Select(m => m.Split(",").Select(n => Convert.ToInt32(n)).ToArray());

    return (rules, updates);
}


List<int> GenerateRuleBasedList(IEnumerable<(int left, int right)> rules)
{
    List<int> ruleList = [];
    rules = rules.OrderBy(m => m.left);
    foreach (var rule in rules)
    {
        var rightNumberExist = ruleList.Contains(rule.right);
        var leftNumberExist = ruleList.Contains(rule.left);

        var rightNumberIndex = ruleList.IndexOf(rule.right);
        var leftNumberIndex =ruleList.IndexOf(rule.left);

        if (rightNumberExist)
        {
            if (leftNumberExist)
            {
                if (leftNumberIndex < rightNumberIndex) continue;
                ruleList.RemoveAt(leftNumberIndex);
            }
            ruleList.Insert(rightNumberIndex, rule.left);
            continue;
        }

        if (leftNumberExist)
        {
            if (rightNumberExist)
            {
                if (rightNumberIndex < leftNumberIndex) continue;
                ruleList.RemoveAt(rightNumberIndex);
            }

            var indexToPut = leftNumberIndex + 1;
            if (indexToPut >= ruleList.Count)
                ruleList.Add(rule.right);
            else
                ruleList.Insert(indexToPut, rule.right);
            continue;
        }

        ruleList.Add(rule.left);
        ruleList.Add(rule.right);
    }

    return ruleList;
}