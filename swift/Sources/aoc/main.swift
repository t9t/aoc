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
    8: Day8.init,
    9: Day9.init,
/*newday*/]

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
