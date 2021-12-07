import Foundation

class Day7: Day {
    private let inputPositions: Array<Int>
    private let minPos: Int, maxPos: Int

    init(_ input: String) {
        inputPositions = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: ",").map {
            Int($0)!
        }
        var minPos = Int.max, maxPos = Int.min
        for pos in inputPositions {
            minPos = min(minPos, pos)
            maxPos = max(maxPos, pos)
        }
        self.minPos = minPos
        self.maxPos = maxPos
    }

    func part1() -> Int {
        determineMinFuel { $0 }
    }

    func part2() -> Int {
        determineMinFuel { $0 * ($0 + 1) / 2 }
    }

    private func determineMinFuel(_ fuelForPosFunc: (Int) -> Int) -> Int {
        var minFuel = Int.max
        for target in minPos...maxPos {
            var fuel = 0
            for pos in inputPositions {
                fuel += fuelForPosFunc(abs(target - pos))
            }
            minFuel = min(minFuel, fuel)
        }
        return minFuel
    }
}
