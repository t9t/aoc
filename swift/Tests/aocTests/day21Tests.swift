import XCTest
import class Foundation.Bundle

@testable import aoc

final class day21Tests: XCTestCase {
    private let input = """
                        Player 1 starting position: 4
                        Player 2 starting position: 8
                        """

    func testPart1() throws {
        let n = Day21(input).part1()
        XCTAssertEqual(n, 739785)
    }

    func testPart2() throws {
        let n = Day21(input).part2()
        XCTAssertEqual(n, 444356092776315)
    }
}
