import XCTest
import class Foundation.Bundle

@testable import aoc

final class day25Tests: XCTestCase {
    private let input = """
                        v...>>.vv>
                        .vv>>.vv..
                        >>.>v>...v
                        >>v>>.>.v.
                        v>v.vv.v..
                        >.>>..v...
                        .vv..>.>v.
                        v.v..>>v.v
                        ....v..v.>
                        """

    func testPart1() throws {
        let n = Day25(input).part1()
        XCTAssertEqual(n, 58)
    }

    func testPart2() throws {
        let n = Day25(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
