import Foundation

class Day14: Day {
    private let inputLines: Array<Substring>

    private let polymerTemplate: String
    private let pairInsertionRules: [Pair: Character]

    init(_ input: String) {
        inputLines = input.split(separator: "\n")

        let inputParts = input.components(separatedBy: "\n\n")
        polymerTemplate = String(inputParts[0].trimmingCharacters(in: .whitespacesAndNewlines))

        var pairInsertionRules: [Pair: Character] = [:]
        for line in inputParts[1].trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n") {
            let lineParts = line.components(separatedBy: " -> ")
            let pair = Pair(left: charAt(lineParts[0], 0), right: charAt(lineParts[0], 1))
            pairInsertionRules[pair] = Character(lineParts[1])
        }
        self.pairInsertionRules = pairInsertionRules
    }

    func part1() -> Int {
        var polymer = polymerTemplate
        for _ in 1...10 {
            var newPolymer = ""

            for (i, c) in polymer.enumerated() {
                newPolymer.append(c)
                if i >= polymer.count - 1 {
                    break
                }
                let next = charAt(polymer, i + 1)
                let pair = Pair(left: c, right: next)
                let insertion = pairInsertionRules[pair]!
                newPolymer.append(insertion)
            }

            polymer = newPolymer
        }

        var elementCounts = [Character: Int]()
        for c in polymer {
            if let count = elementCounts[c] {
                elementCounts[c] = count + 1
            } else {
                elementCounts[c] = 1
            }
        }
        var minQuantity = Int.max, maxQuantity = Int.min
        for (_,count) in elementCounts {
            minQuantity = min(minQuantity, count)
            maxQuantity = max(maxQuantity, count)
        }

        return maxQuantity - minQuantity
    }

    func part2() -> Int {
        return 1337
    }

    private struct Pair: Hashable {
        let left: Character
        let right: Character
    }
}

private func charAt(_ s: String, _ i: Int) -> Character {
    s[s.index(s.startIndex, offsetBy: i)]
}
