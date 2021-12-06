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
        return 1337
    }
}
