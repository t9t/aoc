import Foundation

class Day3 {
    let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.split(separator: "\n")
    }

    func part1() -> Int {
        let len = inputLines[0].count
        var gamma = 0, epsilon = 0, n = 1
        for i in (0...len - 1).reversed() {
            var zeroes = 0, ones = 0
            for line in inputLines {
                if line[line.index(line.startIndex, offsetBy: i)] == "0" {
                    zeroes += 1
                } else {
                    ones += 1
                }
            }
            if zeroes < ones {
                gamma += n
            } else {
                epsilon += n
            }
            n = n << 1
        }
        return gamma * epsilon
    }

    func part2() -> Int {
        return 1337
    }
}
