import XCTest
import class Foundation.Bundle

@testable import aoc

final class day14Tests: XCTestCase {
    private let input = """
                        NNCB

                        CH -> B
                        HH -> N
                        CB -> H
                        NH -> C
                        HB -> C
                        HC -> B
                        HN -> C
                        NN -> C
                        BH -> H
                        NC -> B
                        NB -> B
                        BN -> B
                        BB -> N
                        BC -> B
                        CC -> N
                        CN -> C
                        """

    func testPart1() throws {
        let n = Day14(input).part1()
        XCTAssertEqual(n, 1588)
    }

    func testPart2() throws {
        let n = Day14(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
