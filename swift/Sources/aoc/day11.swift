import Foundation

class Day11: Day {
    private let octopuses: Array<Array<Octopus>>

    init(_ input: String) {
        octopuses = input
                .trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: "\n")
                .map({ line in line.map({ Octopus(energyLevel: Int(String($0))!) }) })
    }

    func part1() throws -> Int {
        try findSolution(findTotalFlashes: true, maxIterations: 100)
    }

    func part2() throws -> Int {
        try findSolution(findTotalFlashes: false, maxIterations: 10_000)
    }

    private func findSolution(findTotalFlashes: Bool, maxIterations: Int) throws -> Int {
        let maxY = octopuses.count - 1, maxX = octopuses[0].count - 1
        var totalFlashes = 0

        for step in 1...maxIterations {
            var allFlashed = true
            for row in octopuses {
                for octopus in row {
                    if octopus.energyLevel > 0 {
                        allFlashed = false
                        octopus.energyLevel += 1
                    } else if octopus.energyLevel <= 0 {
                        octopus.energyLevel = 1
                    }
                }
            }
            if !findTotalFlashes && allFlashed {
                return step - 1
            }

            while true {
                var anyFlashed = false
                outer: for (y, row) in octopuses.enumerated() {
                    for (x, octopus) in row.enumerated() {
                        if octopus.energyLevel > 9 {
                            totalFlashes += 1
                            anyFlashed = true
                            for dy in max(0, y - 1)...min(y + 1, maxY) {
                                for dx in max(0, x - 1)...min(x + 1, maxX) {
                                    let other = octopuses[dy][dx]
                                    if (dx != x || dy != y) && other.energyLevel > 0 {
                                        other.energyLevel += 1
                                    }
                                }
                            }
                            octopus.energyLevel = 0
                        }
                    }
                }
                if !anyFlashed {
                    break
                }
            }
        }
        if !findTotalFlashes {
            throw NoAnswerFound()
        }
        return totalFlashes
    }

    private class Octopus: CustomStringConvertible {
        var energyLevel: Int

        init(energyLevel: Int) {
            self.energyLevel = energyLevel
        }

        var description: String {
            "\(energyLevel)"
        }
    }

    private struct NoAnswerFound: Error {
    }
}
