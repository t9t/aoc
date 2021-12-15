import Foundation

protocol Day {
    func part1() throws -> Int
    func part2() throws -> Int
}

protocol StringDay {
    func part1() throws -> String
    func part2() throws -> String
}

class Wrapped: StringDay {
    private let intDay: Day

    init(_ intDay: Day) {
        self.intDay = intDay
    }

    func part1() throws -> String {
        String(try intDay.part1())
    }

    func part2() throws -> String {
        String(try intDay.part2())
    }
}

func stringWrapped(_ day: @escaping (String) -> Day) -> (String) -> StringDay {
    {
        Wrapped(day($0))
    }
}

internal class Days {
    private static let days: [Int: (String) -> StringDay] = [
        1: stringWrapped(Day1.init),
        2: stringWrapped(Day2.init),
        3: stringWrapped(Day3.init),
        4: stringWrapped(Day4.init),
        5: stringWrapped(Day5.init),
        6: stringWrapped(Day6.init),
        7: stringWrapped(Day7.init),
        8: stringWrapped(Day8.init),
        9: stringWrapped(Day9.init),
        10: stringWrapped(Day10.init),
        11: stringWrapped(Day11.init),
        12: stringWrapped(Day12.init),
        13: Day13.init,
        14: stringWrapped(Day14.init),
        15: stringWrapped(Day15.init),
/*newday*/]

    static func get(num: Int, input: String) -> StringDay {
        days[num]!(input)
    }

    static func getAllDayNumbers() -> Array<Int> {
        days.keys.sorted()
    }
}

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
let day = Days.get(num: dayNum, input: input)

let start = Date()
let output = try (part == 1 ? day.part1 : day.part2)()
let took = start.timeIntervalSinceNow

print(output)
print("Took \(formatInterval(took))")

private func formatInterval(_ intervalInput: TimeInterval) -> String {
    let interval = abs(intervalInput)
    let totalSeconds = Int(interval)

    var out = Array<String>()
    let minutes = totalSeconds / 60
    if minutes > 0 {
        out.append("\(minutes)m")
    }
    let seconds = totalSeconds % 60
    if seconds > 0 {
        out.append("\(seconds)s")
    }
    let ms = Int(interval.truncatingRemainder(dividingBy: 1) * 1000)
    if ms > 0 {
        out.append("\(ms)ms")
    }
    return out.joined(separator: " ")
}
