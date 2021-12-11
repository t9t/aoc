import Foundation

class Day11: Day {
    private let octopuses: Array<Array<Octopus>>

    init(_ input: String) {
        octopuses = input
                .trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: "\n")
                .map({ line in line.map({ Octopus(energyLevel: Int(String($0))!, hasFlashed: false) }) })
    }

    func part1() -> Int {
        let maxY = octopuses.count - 1, maxX = octopuses[0].count - 1
        var totalFlashes = 0

        for _ in 1...100 {
            for row in octopuses {
                for octopus in row {
                    octopus.energyLevel += 1
                    octopus.hasFlashed = false
                }
            }

            while true {
                var anyFlashed = false
                outer: for (y, row) in octopuses.enumerated() {
                    for (x, octopus) in row.enumerated() {
                        if octopus.energyLevel > 9 && !octopus.hasFlashed {
                            totalFlashes += 1
                            anyFlashed = true
                            for dy in max(0, y - 1)...min(y + 1, maxY) {
                                for dx in max(0, x - 1)...min(x + 1, maxX) {
                                    if dx != x || dy != y {
                                        octopuses[dy][dx].energyLevel += 1
                                    }
                                }
                            }
                            octopus.hasFlashed = true
                        }
                    }
                }
                if !anyFlashed {
                    break
                }
            }

            for row in octopuses {
                for octopus in row {
                    if octopus.energyLevel > 9 {
                        octopus.energyLevel = 0
                    }
                }
            }
        }

        return totalFlashes
    }

    func part2() throws -> Int {
        let maxY = octopuses.count - 1, maxX = octopuses[0].count - 1

        for step in 1...10_000 {
            for row in octopuses {
                for octopus in row {
                    octopus.energyLevel += 1
                    octopus.hasFlashed = false
                }
            }

            while true {
                var anyFlashed = false
                outer: for (y, row) in octopuses.enumerated() {
                    for (x, octopus) in row.enumerated() {
                        if octopus.energyLevel > 9 && !octopus.hasFlashed {
                            anyFlashed = true
                            for dy in max(0, y - 1)...min(y + 1, maxY) {
                                for dx in max(0, x - 1)...min(x + 1, maxX) {
                                    if dx != x || dy != y {
                                        octopuses[dy][dx].energyLevel += 1
                                    }
                                }
                            }
                            octopus.hasFlashed = true
                        }
                    }
                }
                if !anyFlashed {
                    break
                }
            }

            var allFlashed = true
            for row in octopuses {
                for octopus in row {
                    if octopus.energyLevel > 9 {
                        octopus.energyLevel = 0
                    }
                    if !octopus.hasFlashed {
                        allFlashed = false
                    }
                }
            }
            if allFlashed {
                return step
            }
        }
        throw NoAnswerFound()
    }

    private class Octopus: CustomStringConvertible {
        var energyLevel: Int
        var hasFlashed: Bool

        init(energyLevel: Int, hasFlashed: Bool) {
            self.energyLevel = energyLevel
            self.hasFlashed = hasFlashed
        }

        var description: String {
            "\(energyLevel)"
        }
    }

    private struct NoAnswerFound: Error {
    }
}
