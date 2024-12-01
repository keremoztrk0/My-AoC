using System.Text.RegularExpressions;

string[] file = await File.ReadAllLinesAsync("input_day2.txt");
// int part1_result = day1_part1(file);
int part2_result = day1_part2(file);
// Console.WriteLine(part1_result);
Console.WriteLine(part2_result);
int day1_part1(string[] input)
{
    const int MAX_RED = 12;
    const int MAX_GREEN = 13;
    const int MAX_BLUE = 14;
    int total_sum = 0;
    foreach (string line in input)
    {
        Console.WriteLine(line);
        Match gameIdMatch = Regex.Match(line, "Game\\s(\\d+):");
        int gameId = int.Parse(gameIdMatch.Groups[1].Value);
        MatchCollection colorNumbers = Regex.Matches(line, "(?:(\\d+)\\s(red|green|blue))");
        bool allRedsValid = colorNumbers.Where(m => m.Groups[2].Value == "red").Select(m => int.Parse(m.Groups[1].Value)).All(m => m <= MAX_RED);
        bool allGreensValid = colorNumbers.Where(m => m.Groups[2].Value == "green").Select(m => int.Parse(m.Groups[1].Value)).All(m => m <= MAX_GREEN);
        bool allBluesValid = colorNumbers.Where(m => m.Groups[2].Value == "blue").Select(m => int.Parse(m.Groups[1].Value)).All(m => m <= MAX_BLUE);
        if (allRedsValid && allBluesValid && allGreensValid)
        {
            total_sum += gameId;
            Console.WriteLine($"Game {gameId} is valid");
        }

    }
    return total_sum;
}
int day1_part2(string[] input)
{
    int total_sum = 0;
    foreach (string line in input)
    {
        Console.WriteLine(line);
        MatchCollection colorNumbers = Regex.Matches(line, "(?:(\\d+)\\s(red|green|blue))");
        int minRed = colorNumbers.Where(m => m.Groups[2].Value == "red").Select(m => int.Parse(m.Groups[1].Value)).Max();
        int minGreen = colorNumbers.Where(m => m.Groups[2].Value == "green").Select(m => int.Parse(m.Groups[1].Value)).Max();
        int minBlue = colorNumbers.Where(m => m.Groups[2].Value == "blue").Select(m => int.Parse(m.Groups[1].Value)).Max();
        Console.WriteLine(minRed);
        Console.WriteLine(minBlue);
        Console.WriteLine(minGreen);
        int power = minRed * minBlue * minGreen;
        Console.WriteLine(power);

        total_sum+=power;
    }
    return total_sum;
}