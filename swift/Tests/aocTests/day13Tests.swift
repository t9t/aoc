import XCTest
import class Foundation.Bundle

@testable import aoc

final class day13Tests: XCTestCase {
    private let input = """
                        6,10
                        0,14
                        9,10
                        0,3
                        10,4
                        4,11
                        6,0
                        6,12
                        4,1
                        0,13
                        10,12
                        3,4
                        3,0
                        8,4
                        1,10
                        2,14
                        8,10
                        9,0

                        fold along y=7
                        fold along x=5
                        """

    func testPart1() throws {
        let n = Day13(input).part1()
        XCTAssertEqual(n, "17")
    }

    func testPart2() throws {
        let n = Day13(input).part2()
        XCTAssertEqual(n, "0")
    }
}
