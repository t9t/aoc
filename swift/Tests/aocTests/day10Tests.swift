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
        XCTAssertEqual(n, 1337)
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
}
