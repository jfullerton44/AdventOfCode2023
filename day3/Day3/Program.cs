namespace Day3
{
    internal class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello, World!");
            string filePath = "C:\\src\\AdventOfCode2023\\day3\\Day3\\in.txt";
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
        }

        static int Part1(List<string> lines)
        {
            var total = 0;
            var data = new List<List<char>>();
            foreach (string line in lines)
            {
                var lineData = new List<char>();
                foreach (char character in line)
                {
                    lineData.Add(character);
                }
                data.Add(lineData);
            }

            var nums = new List<NumValue>();
            for (var i = 0; i < data.Count; i++)
            {
                var line = data[i];
                for (var j = 0; j < line.Count; j++)
                {
                    var character = line[j];
                    int innerTotal = 0;
                    var startj = j;
                    if (char.GetNumericValue(character) != -1.0)
                    {
                       innerTotal = (int)char.GetNumericValue(character);
                        j++;
                       while(j<line.Count && char.GetNumericValue(line[j]) != -1.0)
                       {
                           innerTotal = innerTotal * 10 + (int)char.GetNumericValue(line[j]);
                            j++;
                       }
                       nums.Add(new NumValue()
                       {
                           value = innerTotal,
                           startIndex = startj,
                           endIndex = j,
                           row = i
                       });
                    }
                }
            }

            foreach (var num in nums)
            {
                if (num.isValid(data))
                {
                    total += num.value;
                }
            }

            return total;
        }


        static int Part2(List<string> lines)
        {
            var total = 0;
            var data = new List<List<char>>();
            foreach (string line in lines)
            {
                var lineData = new List<char>();
                foreach (char character in line)
                {
                    lineData.Add(character);
                }
                data.Add(lineData);
            }

            var nums = new List<NumValue>();
            var stars = new List<Star>();
            for (var i = 0; i < data.Count; i++)
            {
                var line = data[i];
                for (var j = 0; j < line.Count; j++)
                {
                    var character = line[j];
                    int innerTotal = 0;
                    var startj = j;
                    if (char.GetNumericValue(character) != -1.0)
                    {
                        innerTotal = (int)char.GetNumericValue(character);
                        j++;
                        while (j < line.Count && char.GetNumericValue(line[j]) != -1.0)
                        {
                            innerTotal = innerTotal * 10 + (int)char.GetNumericValue(line[j]);
                            j++;
                        }
                        nums.Add(new NumValue()
                        {
                            value = innerTotal,
                            startIndex = startj,
                            endIndex = j,
                            row = i
                        });
                        j--;
                    }
                    if (character == '*')
                    {
                        stars.Add(new Star()
                        {
                            col = j,
                            row = i
                        });
                    }

                }
            }

            foreach (var star in stars)
            {
               foreach (var num in nums)
                {
                    if (star.isTouching(num))
                    {
                       star.numValues.Add(num);
                    }
                }
            }

            stars = stars.Where(x => x.numValues.Count == 2).ToList();

            foreach (var star in stars)
            {
                total += star.numValues[0].value * star.numValues[1].value;
            }

            return total;
        }
        
        internal class NumValue
        {
            internal int value;
            internal int startIndex;
            internal int endIndex;
            internal int row;

            internal bool isValid(List<List<char>> data)
            {
                for (var col = startIndex-1; col <= endIndex; col++)
                {
                    for (var row = this.row - 1; row <= this.row + 1; row++)
                    {
                        if (isSymbol(col, row, data))
                        {
                            return true;
                        }
                    }
                }
                return false;
            }
        }

        internal class Star
        {
            internal int col;
            internal int row;
            internal List<NumValue> numValues = new List<NumValue>();

            internal bool isTouching(NumValue num)
            {
               if (Math.Abs(num.row - this.row) <= 1)
                {
                    if (Math.Abs(num.startIndex - this.col) <= 1 || Math.Abs(num.endIndex-1 - this.col) <= 1)
                    {
                        return true;
                    }
                }
                return false;
            }
        }

        static bool isSymbol(int col, int row, List<List<char>> data)
        {
            if (col < 0 || row >= data.Count || row < 0 || col >= data[row].Count)
            {
                return false;
            }
            var character = data[row][col];
            return (char.GetNumericValue(character) == -1.0 && character != '.');
        }
    }

}
