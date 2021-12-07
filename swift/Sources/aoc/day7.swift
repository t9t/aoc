import Foundation

class Day7: Day {
    let inputPositions: Array<Int>
    let minPos: Int, maxPos: Int

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
        var minFuel = Int.max
        for pos in minPos...maxPos {
            let fuel = calculateFuelForPos(pos)
            print("pos", pos, "fuel", fuel)
            minFuel = min(minFuel, fuel)
        }
        return minFuel
    }

    func part2() -> Int {
        return 1337
    }

    private func calculateFuelForPos(_ target: Int) -> Int {
        var fuel = 0
        for pos in inputPositions {
            fuel += abs(target - pos)
        }
        return fuel
    }
}
