import XCTest
import class Foundation.Bundle

@testable import aoc

final class day3Tests: XCTestCase {
    private let input = """
                        00100
                        11110
                        10110
                        10111
                        10101
                        01111
                        00111
                        11100
                        10000
                        11001
                        00010
                        01010
                        """

    func testPart1() throws {
        let n = Day3(input).part1()
        XCTAssertEqual(n, 198)
    }

    func testPart2() throws {
        let n = Day3(input).part2()
        XCTAssertEqual(n, 230)
    }
}
