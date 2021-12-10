import XCTest
import class Foundation.Bundle

@testable import aoc

final class day10Tests: XCTestCase {
    private let input = """
                        [({(<(())[]>[[{[]{<()<>>
                        [(()[<>])]({[<{<<[]>>(
                        {([(<{}[<>[]}>{[]{[(<()>
                        (((({<>}<{<{<>}{[]{[]{}
                        [[<[([]))<([[{}[[()]]]
                        [{[{({}]{}}([{[{{{}}([]
                        {<[[]]>}<{[{[{[]{()[[[]
                        [<(<(<(<{}))><([]([]()
                        <{([([[(<>()){}]>(<<{{
                        <{([{{}}[<[[[<>{}]]]>[]]
                        """

    func testPart1() throws {
        let n = Day10(input).part1()
        XCTAssertEqual(n, 26397)
    }

    func testPart2() throws {
        let n = Day10(input).part2()
        XCTAssertEqual(n, 288957)
    }

    func testDetermineSyntaxErrorScore() throws {
        let cases = [
            "()": 0,
            "[]": 0,
            "([])": 0,
            "{()()()}": 0,
            "<([{}])>": 0,
            "[<>({}){}[([])<>]]": 0,
            "(((((((((())))))))))": 0,
            "(]": 57,
            "{()()()>": 25137,
            "(((()))}": 1197,
        ]
        for (testCase, expected) in cases {
            let actual = Day10.determineSyntaxErrorScore(testCase)
            XCTAssertEqual(actual, expected)
        }
    }

    func testFindCompletionStringX() throws {
        let cases = [
            "[]": Day10.MatchingResult.Complete,
            "[()(){}(({<<>>}))]": Day10.MatchingResult.Complete,
            "[": Day10.MatchingResult.Incomplete(completionString: "]"),
            "[()": Day10.MatchingResult.Incomplete(completionString: "]"),
            "[(){]": Day10.MatchingResult.Corrupted(firstInvalidCharacter: "]"),
            "[](){<>": Day10.MatchingResult.Incomplete(completionString: "}"),
            "[({(<(())[]>[[{[]{<()<>>": Day10.MatchingResult.Incomplete(completionString: "}}]])})]"),
            "[(()[<>])]({[<{<<[]>>(": Day10.MatchingResult.Incomplete(completionString: ")}>]})"),
            "(((({<>}<{<{<>}{[]{[]{}": Day10.MatchingResult.Incomplete(completionString: "}}>}>))))"),
            "{<[[]]>}<{[{[{[]{()[[[]": Day10.MatchingResult.Incomplete(completionString: "]]}}]}]}>"),
            "<{([{{}}[<[[[<>{}]]]>[]]": Day10.MatchingResult.Incomplete(completionString: "])}>"),
        ]
        for (testCase, expected) in cases {
            let actual = Day10.matchBraces(testCase)
            XCTAssertEqual(actual, expected)
        }
    }

    func testCalculateScore() throws {
        let cases = [
            "}}]])})]": 288957,
            ")}>]})": 5566,
            "}}>}>))))": 1480781,
            "]]}}]}]}>": 995444,
            "])}>": 294,
        ]
        for (testCase, expected) in cases {
            let actual = Day10.calculateScore(testCase)
            XCTAssertEqual(actual, expected)
        }
    }
}
