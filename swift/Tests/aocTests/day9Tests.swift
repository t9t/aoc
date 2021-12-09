import XCTest
import class Foundation.Bundle

@testable import aoc

final class day9Tests: XCTestCase {
    private let input = """
                        2199943210
                        3987894921
                        9856789892
                        8767896789
                        9899965678
                        """

    func testPart1() throws {
        let n = Day9(input).part1()
        XCTAssertEqual(n, 15)
    }

    func testPart2() throws {
        let n = Day9(input).part2()
        XCTAssertEqual(n, 1134)
    }
}
