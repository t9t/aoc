import XCTest
import class Foundation.Bundle

@testable import aoc

final class day17Tests: XCTestCase {
    private let input = "target area: x=20..30, y=-10..-5"

    func testPart1() throws {
        let n = Day17(input).part1()
        XCTAssertEqual(n, 45)
    }

    func testPart2() throws {
        let n = Day17(input).part2()
        XCTAssertEqual(n, 112)
    }
}
