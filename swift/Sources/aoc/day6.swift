import Foundation

class Day6: Day {
    let inputFishes: Array<Int>

    init(_ input: String) {
        inputFishes = input
                .trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: ",")
                .map {
                    Int($0.trimmingCharacters(in: .whitespaces))!
                }
    }

    func part1() -> Int {
        var fishes = inputFishes
        for _ in 1...80 {
            var newFishes = Array<Int>()
            for fish in fishes {
                if fish == 0 {
                    newFishes.append(6) // Existing fish reset
                    newFishes.append(8) // New fish created
                    continue
                }
                newFishes.append(fish - 1)
            }
            fishes = newFishes
        }
        return fishes.count
    }

    func part2() -> Int {
        var fishes = fishMap()
        for _ in 1...256 {
            let currentZeroes = fishes[0]!
            let newSixes = currentZeroes
            let newEights = currentZeroes
            fishes[0] = 0

            for fish in 1...9 {
                fishes[fish - 1] = fishes[fish]!
                fishes[fish] = 0
            }

            fishes[6] = fishes[6]! + newSixes
            fishes[8] = fishes[8]! + newEights
        }

        var total = 0
        for (_, n) in fishes {
            total += n
        }

        return total
    }

    private func fishMap() -> [Int: Int] {
        var map: [Int: Int] = [:]

        for i in 0...9 {
            map[i] = 0
        }

        for fish in inputFishes {
            if let n = map[fish] {
                map[fish] = n + 1
            }
        }
        return map
    }
}
