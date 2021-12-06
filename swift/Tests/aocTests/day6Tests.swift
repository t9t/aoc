import XCTest
import class Foundation.Bundle

@testable import aoc

final class day6Tests: XCTestCase {
    private let input = "3,4,3,1,2"

    func testPart1() throws {
        let n = Day6(input).part1()
        XCTAssertEqual(n, 5934)
    }

    func testPart2() throws {
        let n = Day6(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
