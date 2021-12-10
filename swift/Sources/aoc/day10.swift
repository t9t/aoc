import Foundation

class Day10: Day {
    private let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.split(separator: "\n")
    }

    func part1() -> Int {
        return inputLines.map({ Day10.determineSyntaxErrorScore($0) }).reduce(0, +)
    }

    func part2() -> Int {
        return 1337
    }

    private static let matchesMap: [Character: Character] = ["(": ")", "[": "]", "{": "}", "<": ">"]
    private static let scoreTable: [Character: Int] = [")": 3, "]": 57, "}": 1197, ">": 25137]

    internal static func determineSyntaxErrorScore<S: StringProtocol>(_ line: S) -> Int {
        var expectedClosingBraces = Array<Character>()
        for c in line {
            if let closing = matchesMap[c] {
                expectedClosingBraces.append(closing)
            } else {
                if c != expectedClosingBraces.popLast() {
                    return scoreTable[c]!
                }
            }
        }
        return 0
    }
}
