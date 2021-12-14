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
            }
            pairs = newPairs
        }

        var elementCounts = [Character: Int]()
        // Since we're counting only left characters, the last character of the template isn't counted; need to add it explicitly
        elementCounts[charAt(polymerTemplate, polymerTemplate.count - 1)] = 1
        for (pair, pairCount) in pairs {
            let c = pair.left
            if let elementCount = elementCounts[c] {
                elementCounts[c] = elementCount + pairCount
            } else {
                elementCounts[c] = pairCount
            }
        }

        return elementCounts.values.max()! - elementCounts.values.min()!
    }

    private func buildPairMap() -> [Pair: Int] {
        var pairs = [Pair: Int]()
        for (pair, _) in pairInsertionRules {
            // Assumption: there exists a rule for every conceivable pair
            pairs[pair] = 0
        }
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
