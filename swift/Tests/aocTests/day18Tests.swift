import XCTest
import class Foundation.Bundle

@testable import aoc

final class day18Tests: XCTestCase {
    private let input = """
                        [[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
                        [[[5,[2,8]],4],[5,[[9,9],0]]]
                        [6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
                        [[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
                        [[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
                        [[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
                        [[[[5,4],[7,7]],8],[[8,3],8]]
                        [[9,3],[[9,9],[6,[4,9]]]]
                        [[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
                        [[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
                        """

    func testPart1() throws {
        let n = try Day18(input).part1()
        XCTAssertEqual(n, 4140)
    }

    func testPart2() throws {
        let n = try Day18(input).part2()
        XCTAssertEqual(n, 1337)
    }

    func testAddition() throws {
        let testCases: Array<(Day18.Number, Day18.Number, Day18.Number)> = [
            (num(1), num(2), pair(1,2)),
            (num(1), pair(num(2), num(3)), pair(1, pair(num(2), num(3)))),
            (pair(1, 2), pair(pair(3, 4), 5), pair(pair(1, 2), pair(pair(3, 4), 5)))
        ]
        for (l, r, expected) in testCases {
            XCTAssertEqual(l + r, expected)
        }
    }

    func testParseNumber() throws {
        let testCases: [String: Day18.Number] = [
            "[1,2]": pair(1, 2),
            "[[1,2],3]": pair(pair(1, 2), 3),
            "[9,[8,7]]": pair(num(9), pair(8, 7)),
            "[[1,9],[8,5]]": pair(pair(1, 9), pair(8, 5)),
            "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]": pair(pair(pair(pair(1, 2), pair(3, 4)), pair(pair(5, 6), pair(7, 8))), num(9)),
            "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]": pair(pair(pair(9, pair(3, 8)), pair(pair(0, 9), 6)), pair(pair(pair(3, 7), pair(4, 9)), num(3))),
        ]
        for (input, expected) in testCases {
            XCTAssertEqual(try Day18.parseNumber(input), expected)
        }
    }

    func testExplodeOnceIfNecessary() throws {
        let testCases: Array<(String, String)> = [
            ("[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"),
            ("[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"),
            ("[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"),
            ("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
            ("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"),
        ]
        for (inputString, expected) in testCases {
            let input = (try Day18.parseNumber(inputString) as? Day18.PairNumber)!
            XCTAssertEqual(Day18.explodeOnceIfNecessary(input), try Day18.parseNumber(expected))
        }
    }

    private func num(_ value: Int) -> Day18.RegularNumber {
        Day18.RegularNumber(value: value)
    }

    private func pair(_ left: Day18.Number, _ right: Day18.Number) -> Day18.PairNumber {
        Day18.PairNumber(left: left, right: right)
    }

    private func pair(_ left: Int, _ right: Int) -> Day18.PairNumber {
        pair(num(left), num(right))
    }

    private func pair(_ left: Int, _ right: Day18.PairNumber) -> Day18.PairNumber {
        pair(num(left), right)
    }

    private func pair(_ left: Day18.PairNumber, _ right: Int) -> Day18.PairNumber {
        pair(left, num(right))
    }
}
