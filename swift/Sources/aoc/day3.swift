import Foundation

class Day3: Day {
    private let inputLines: Array<Substring>

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
        let oxygen = findRating(mostCommon: true)
        let co2 = findRating(mostCommon: false)
        return oxygen * co2
    }

    private func findRating(mostCommon: Bool) -> Int {
        let len = inputLines[0].count
        var linesLeft = inputLines
        for i in 0...len - 1 {
            var withOnes = Array<Substring>(), withZeroes = Array<Substring>()
            for line in linesLeft {
                let c: Character = line[line.index(line.startIndex, offsetBy: i)]
                if c == "0" {
                    withZeroes.append(line)
                } else {
                    withOnes.append(line)
                }
            }
            linesLeft = ((withOnes.count >= withZeroes.count) == mostCommon) ? withOnes : withZeroes
            if linesLeft.count == 1 {
                break
            }
        }
        return Int(linesLeft[0], radix: 2)!
    }
}
