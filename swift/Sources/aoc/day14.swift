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
        runInsertionProcess(steps: 10)
    }

    func part2() -> Int {
        runInsertionProcess(steps: 40)
    }

    private func runInsertionProcess(steps: Int) -> Int {
        var elementCounts = [Character: Int]()
        // Assumption: for every element there is a rule with that element appearing "left"
        pairInsertionRules.forEach({ elementCounts[$0.key.left] = 0 })
        polymerTemplate.forEach({ elementCounts[$0] = elementCounts[$0]! + 1 })

        var pairs = buildPairMap()
        for _ in 1...steps {
            var newPairs = pairs
            for (pair, count) in pairs {
                let insertion = pairInsertionRules[pair]!
                let newPairLeft = Pair(left: pair.left, right: insertion)
                let newPairRight = Pair(left: insertion, right: pair.right)
                newPairs[pair] = newPairs[pair]! - count
                newPairs[newPairLeft] = newPairs[newPairLeft]! + count
                newPairs[newPairRight] = newPairs[newPairRight]! + count
                elementCounts[insertion] = elementCounts[insertion]! + count
            }
            pairs = newPairs
        }

        return elementCounts.values.max()! - elementCounts.values.min()!
    }

    private func buildPairMap() -> [Pair: Int] {
        var pairs = [Pair: Int]()
        // Assumption: there exists a rule for every conceivable pair
        pairInsertionRules.forEach({ pairs[$0.key] = 0 })
        for i in 0...polymerTemplate.count - 2 {
            let pair = Pair(left: charAt(polymerTemplate, i), right: charAt(polymerTemplate, i + 1))
            pairs[pair] = pairs[pair]! + 1
        }
        return pairs
    }

    private struct Pair: Hashable, CustomStringConvertible {
        let left: Character
        let right: Character
        var description: String {
            "\(left)\(right)"
        }
    }
}

private func charAt(_ s: String, _ i: Int) -> Character {
    s[s.index(s.startIndex, offsetBy: i)]
}
