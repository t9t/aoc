import Foundation

class Day10: Day {
    private let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.split(separator: "\n")
    }

    func part1() -> Int {
        inputLines.map({ Day10.determineSyntaxErrorScore($0) }).reduce(0, +)
    }

    func part2() -> Int {
        let scores = inputLines
                .filter({ Day10.determineSyntaxErrorScore($0) == 0 })
                .map({ Day10.findCompletionString($0) })
                .map({ Day10.calculateScore($0) })
                .sorted()

        return scores[scores.startIndex.advanced(by: scores.count / 2)]
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

    internal static func findCompletionString<S: StringProtocol>(_ line: S) -> String {
        var expectedClosingBraces = Array<Character>()
        var completionString = ""
        for c in line {
            if let closing = matchesMap[c] {
                expectedClosingBraces.append(closing)
            } else {
                while true {
                    let expected = expectedClosingBraces.popLast()
                    if expected == nil || c == expected {
                        break
                    }
                    completionString += String(expected!)
                }
            }
        }
        for c in expectedClosingBraces.reversed() {
            completionString += String(c)
        }
        return completionString
    }

    private static let completionStringScoreTable: [Character: Int] = [")": 1, "]": 2, "}": 3, ">": 4]

    internal static func calculateScore(_ completionString: String) -> Int {
        var score = 0
        for c in completionString {
            score *= 5
            score += completionStringScoreTable[c]!
        }
        return score
    }
}
