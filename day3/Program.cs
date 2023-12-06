// See https://aka.ms/new-console-template for more information
using System.Runtime.CompilerServices;

Console.WriteLine("Hello, World!");
string filePath = "in.txt";
List<string> lines = new List<string>();

try
{
    lines = File.ReadAllLines(filePath).ToList();
}
catch (Exception ex)
{
    Console.WriteLine($"Error reading file: {ex.Message}");
}
Console.WriteLine($"Part 1: {Part1(lines)}");
Console.WriteLine($"Part 2: {Part2(lines)}");

static int Part1(List<string> lines)
{
    var total = 0;
    var data = new List<List<string>();
    foreach (string line in lines)
    {
       var line = new List<string>();
       foreach (string letter in line)
       {
              line.(letter);
       }
    }
    return total;
}


static int Part2(List<string> lines)
{
    var total = 0;
    var id = 1;
    foreach (string line in lines)
    {
        var minRed = 0;
        var minBlue = 0;
        var minGreen = 0;
        var split = line.Split(":");
        var games = split[1].Split(";");
        foreach (string game in games){
            var red = 0;
            var blue = 0;
            var green = 0;

            var picks = game.Split(",");
            foreach (string pick in picks){
                var pickParts = pick.Split(" ");
                var pickColor = pickParts[2];
                if (pickColor == "red"){
                    red += int.Parse(pickParts[1]);
                }
                else if (pickColor == "blue"){
                    blue += int.Parse(pickParts[1]);
                }
                else if (pickColor == "green"){
                    green += int.Parse(pickParts[1]);
                }
            }
            minGreen = Math.Max(minGreen, green);
            minBlue = Math.Max(minBlue, blue);
            minRed = Math.Max(minRed, red);
        }
        total += minRed * minBlue * minGreen;
        id++;
    }
    return total;
}