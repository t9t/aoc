import Foundation

protocol Day {
    func part1() throws -> Int
    func part2() throws -> Int
}

let days: [Int: (String) -> Day] = [
    1: Day1.init,
    2: Day2.init,
    3: Day3.init,
    4: Day4.init,
    5: Day5.init,
    6: Day6.init,
    7: Day7.init,
]

print(CommandLine.arguments)

if CommandLine.arguments.count != 3 {
    print("invalid arguments, provide day & part")
    exit(1)
}

let dayNum = Int(CommandLine.arguments[1])!
let part = Int(CommandLine.arguments[2])!

if part != 1 && part != 2 {
    print("invalid part \(part), has to be 1 or 2")
    exit(1)
}

print("Running day \(dayNum) part \(part)")

let input = try String(contentsOfFile: "../input/2021/\(dayNum).txt")
let day = days[dayNum]!(input)

print(try (part == 1 ? day.part1 : day.part2)())
