import XCTest
import class Foundation.Bundle

@testable import aoc

final class day5Tests: XCTestCase {
    private let input = """
                        0,9 -> 5,9
                        8,0 -> 0,8
                        9,4 -> 3,4
                        2,2 -> 2,1
                        7,0 -> 7,4
                        6,4 -> 2,0
                        0,9 -> 2,9
                        3,4 -> 1,4
                        0,0 -> 8,8
                        5,5 -> 8,2
                        """

    func testPart1() throws {
        let n = Day5(input).part1()
        XCTAssertEqual(n, 5)
    }

    func testPart2() throws {
        let n = Day5(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
