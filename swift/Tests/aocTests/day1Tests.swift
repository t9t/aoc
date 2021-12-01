import XCTest
import class Foundation.Bundle

@testable import aoc

final class day1Tests: XCTestCase {
    private let input = "199\n200\n208\n210\n200\n207\n240\n269\n260\n263\n"

    func testPart1() throws {
        let n = Day1(input).part1()
        XCTAssertEqual(n, 7)
    }

    func testPart2() throws {
        let n = Day1(input).part2()
        XCTAssertEqual(n, 5)
    }
}
