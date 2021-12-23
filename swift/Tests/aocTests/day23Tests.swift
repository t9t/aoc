import XCTest
import class Foundation.Bundle

@testable import aoc

final class day23Tests: XCTestCase {
    private let input = """
                        #############
                        #...........#
                        ###B#C#B#D###
                          #A#D#C#A#
                          #########
                        """

    func testPart1() throws {
        let n = Day23(input).part1()
        XCTAssertEqual(n, 12521)
    }

    func testPart2() throws {
        let n = Day23(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
