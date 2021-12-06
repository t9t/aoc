import Foundation

class Day6: Day {
    let inputFishes: [Int: Int]

    init(_ input: String) {
        var map: [Int: Int] = [:]

        for i in 0...9 {
            map[i] = 0
        }

        input
                .trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: ",")
                .map { Int($0.trimmingCharacters(in: .whitespaces))! }
                .forEach { map[$0] = map[$0]! + 1 }

        inputFishes = map
    }

    func part1() -> Int {
        evolution(80)
    }

    func part2() -> Int {
        evolution(256)
    }

    private func evolution(_ days: Int) -> Int {
        var fishes = inputFishes
        for _ in 1...days {
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

    private func fishMap(_ inputFishes: [Int]) -> [Int: Int] {
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
