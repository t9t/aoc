import XCTest
import class Foundation.Bundle

@testable import aoc

final class day2Tests: XCTestCase {
    private let input = """
                        forward 5
                        down 5
                        forward 8
                        up 3
                        down 8
                        forward 2
                        """

    func testPart1() throws {
        let n = Day2(input).part1()
        XCTAssertEqual(n, 150)
    }

    func testPart2() throws {
        let n = Day2(input).part2()
        XCTAssertEqual(n, 5)
    }
}
