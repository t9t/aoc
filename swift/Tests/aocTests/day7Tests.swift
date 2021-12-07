import XCTest
import class Foundation.Bundle

@testable import aoc

final class day7Tests: XCTestCase {
    private let input = "16,1,2,0,4,2,7,1,2,14\n"

    func testPart1() throws {
        let n = Day7(input).part1()
        XCTAssertEqual(n, 37)
    }

    func testPart2() throws {
        let n = Day7(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
