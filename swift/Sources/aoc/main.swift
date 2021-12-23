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

func stringWrapped(_ day: @escaping (String) throws -> Day) -> (String) throws -> StringDay {
    {
        Wrapped(try day($0))
    }
}

internal class Days {
    private static let days: [Int: (String) throws -> StringDay] = [
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
        16: stringWrapped(Day16.init),
        17: stringWrapped(Day17.init),
        18: stringWrapped(Day18.init),
        19: stringWrapped(Day19.init),
        20: stringWrapped(Day20.init),
        21: stringWrapped(Day21.init),
        22: stringWrapped(Day22.init),
        23: stringWrapped(Day23.init),
/*newday*/]

    static func get(num: Int, input: String) throws -> StringDay {
        try days[num]!(input)
    }

    static func getAllDayNumbers() -> Array<Int> {
        days.keys.sorted()
    }
}

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

let input = try getInput(year: 2021, day: dayNum)
let day = try Days.get(num: dayNum, input: input)

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
    let micros = Int(interval.truncatingRemainder(dividingBy: 1) * 1_000_000)
    if micros > 1_000 {
        out.append("\(micros/1_000)ms")
    } else {
        out.append("\(micros)μs")
    }
    return out.joined(separator: " ")
}

private func getInput(year: Int, day: Int) throws -> String {
    let filename = "../input/\(year)/\(day).txt"
    do {
        return try String(contentsOfFile: filename)
    } catch {
        print("Input file not readable, downloading input file")
        let sesh = String(cString: getpass("Session cookie value: "))
        let downloadedInput = downloadInput(year: year, day: day, sessionCookie: sesh)
        try downloadedInput.write(toFile: filename, atomically: true, encoding: String.Encoding.utf8)
        return downloadedInput
    }
}

private func downloadInput(year: Int, day: Int, sessionCookie: String) -> String {
    let semaphore = DispatchSemaphore(value: 0)
    var req = URLRequest(url: URL(string: "https://adventofcode.com/\(year)/day/\(day)/input")!)
    req.setValue("session=\(sessionCookie)", forHTTPHeaderField: "Cookie")

    var input = ""
    URLSession.shared.dataTask(with: req) { (data, response: URLResponse?, error) in
        let status = (response as? HTTPURLResponse)?.statusCode
        if status == 200 {
            input = String(data: data!, encoding: .utf8)!
        } else {
            fatalError("Invalid status \(status.map(String.init) ?? "?") fetching input")
        }
        semaphore.signal()
    }.resume()
    _ = semaphore.wait(timeout: DispatchTime.distantFuture)
    return input
}
