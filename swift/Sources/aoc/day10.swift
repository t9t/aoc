import Foundation

class Day10: Day {
    private static let matchesMap: [Character: Character] = ["(": ")", "[": "]", "{": "}", "<": ">"]
    private static let corruptionScoreTable: [Character: Int] = [")": 3, "]": 57, "}": 1197, ">": 25137]
    private static let completionStringScoreTable: [Character: Int] = [")": 1, "]": 2, "}": 3, ">": 4]
    private let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.split(separator: "\n")
    }

    func part1() -> Int {
        inputLines.map(Day10.determineSyntaxErrorScore).reduce(0, +)
    }

    func part2() -> Int {
        var scores = Array<Int>()
        for line in inputLines {
            if case .Incomplete(let completionString) = Day10.matchBraces(line) {
                scores.append(Day10.calculateScore(completionString))
            }
        }
        scores.sort()

        return scores[scores.startIndex.advanced(by: scores.count / 2)]
    }

    internal static func determineSyntaxErrorScore<S: StringProtocol>(_ line: S) -> Int {
        if case .Corrupted(let firstInvalidCharacter) = matchBraces(line) {
            return corruptionScoreTable[firstInvalidCharacter]!
        }
        return 0
    }

    internal static func matchBraces<S: StringProtocol>(_ line: S) -> MatchingResult {
        var expectedClosingBraces = Array<Character>()
        for c in line {
            if let closing = matchesMap[c] {
                expectedClosingBraces.append(closing)
                continue
            }
            let expected = expectedClosingBraces.popLast()
            if expected != nil && c != expected {
                return MatchingResult.Corrupted(firstInvalidCharacter: c)
            }
        }

        if expectedClosingBraces.isEmpty {
            return MatchingResult.Complete
        }

        return MatchingResult.Incomplete(completionString: String(expectedClosingBraces.reversed()))
    }

    internal static func calculateScore(_ completionString: String) -> Int {
        var score = 0
        for c in completionString {
            score *= 5
            score += completionStringScoreTable[c]!
        }
        return score
    }

    internal enum MatchingResult: Equatable {
        case Complete
        case Incomplete(completionString: String)
        case Corrupted(firstInvalidCharacter: Character)
    }
}
