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
        XCTAssertEqual(n, 3993)
    }

    func testAddition() throws {
        let testCases: Array<(String, String, String)> = [
            ("[1,2]", "[3,4]", "[[1,2],[3,4]]"),
            ("[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"),
        ]
        for (l, r, expected) in testCases {
            XCTAssertEqual(toString(Day18.add(Day18.tokenize(l), Day18.tokenize(r))), expected)
        }
    }

    func testAddAndReduce1() throws {
        let left = "[[[[4,3],4],4],[7,[[8,4],9]]]"
        let right = "[1,1]"
        let expected = "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
        XCTAssertEqual(toString(Day18.addAndReduce(Day18.tokenize(left), Day18.tokenize(right))), expected)
    }

    func testAddAndReduce2() throws {
        let left = "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"
        let right = "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"
        let expected = "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"
        XCTAssertEqual(toString(Day18.addAndReduce(Day18.tokenize(left), Day18.tokenize(right))), expected)
    }

    func testAddAndReduce3() throws {
        let left = "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]"
        let right = "[[[[4,2],2],6],[8,7]]"
        let expected = "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
        XCTAssertEqual(toString(Day18.addAndReduce(Day18.tokenize(left), Day18.tokenize(right))), expected)
    }

    func testSum1() throws {
        let input = ["[1,1]","[2,2]","[3,3]","[4,4]","[5,5]","[6,6]"]
        let expected = "[[[[5,0],[7,4]],[5,5]],[6,6]]"
        XCTAssertEqual(toString(Day18.sum(input.map(Day18.tokenize))), expected)
    }

    func testSum2() throws {
        let input = ["[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
                     "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
                     "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
                     "[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
                     "[7,[5,[[3,8],[1,4]]]]",
                     "[[2,[2,2]],[8,[8,1]]]",
                     "[2,9]",
                     "[1,[[[9,3],9],[[9,0],[0,7]]]]",
                     "[[[5,[7,4]],7],1]",
                     "[[[[4,2],2],6],[8,7]]"]
        let expected = "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
        XCTAssertEqual(toString(Day18.sum(input.map(Day18.tokenize))), expected)
    }

    func testExplodeOnceIfNecessary() throws {
        let testCases: Array<(String, String)> = [
            ("[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"),
            ("[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"),
            ("[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"),
            ("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
            ("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"),
            ("[4,2]", "[4,2]"), // Exploding not necessary
        ]
        for (input, expected) in testCases {
            let result = Day18.explodeOnceIfNecessary(Day18.tokenize(input))
            XCTAssertEqual(result.1, input != expected)
            XCTAssertEqual(toString(result.0), expected)
        }
    }

    func testSplitOnceIfNecessary() throws {
        let testCases: Array<(String, String)> = [
            ("[10,1]", "[[5,5],1]"),
            ("[11,1]", "[[5,6],1]"),
            ("[12,1]", "[[6,6],1]"),
            ("[1,11]", "[1,[5,6]]"),
            ("[1,[2,[3,11]]]", "[1,[2,[3,[5,6]]]]"),
            ("[1,[2,[[11,[3,4]],5]]]", "[1,[2,[[[5,6],[3,4]],5]]]"),
            ("[4,2]", "[4,2]"), // Splitting unnecessary
            ("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,9],[[11,9],[11,0]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,9],[[[5,6],9],[11,0]]]]"),
        ]
        for (input, expected) in testCases {
            let result = Day18.splitOnceIfNecessary(Day18.tokenize(input))
            XCTAssertEqual(result.1, input != expected)
            XCTAssertEqual(toString(result.0), expected)
        }
    }

    func testReduce() throws {
        let input = "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
        let expected = "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
        XCTAssertEqual(toString(Day18.reduce(Day18.tokenize(input))), expected)
    }

    func testDetermineMagnitude() throws {
        let testCases: Array<(String, Int)> = [
            ("[9,1]", 29),
            ("[1,9]", 21),
            ("[[9,1],[1,9]]", 129),
            ("[[1,2],[[3,4],5]]", 143),
            ("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384),
            ("[[[[1,1],[2,2]],[3,3]],[4,4]]", 445),
            ("[[[[3,0],[5,3]],[4,4]],[5,5]]", 791),
            ("[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137),
            ("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488),
        ]
        for (input, expected) in testCases {
            XCTAssertEqual(Day18.determineMagnitude(Day18.tokenize(input)), expected)
        }
    }

    private func toString(_ num: Array<Day18.Token>) -> String {
        num.map({ t in
            switch t {
            case let .Char(c):
                return String(c)
            case let .Number(n):
                return String(n)
            }
        }).joined()
    }
}
