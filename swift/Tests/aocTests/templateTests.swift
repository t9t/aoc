import XCTest
import class Foundation.Bundle

@testable import aoc

final class templateTests: XCTestCase {
    private let input = """
                        x
                        """

    func testPart1() throws {
        let n = TemplateDay(input).part1()
        XCTAssertEqual(n, 42)
    }

    func testPart2() throws {
        let n = TemplateDay(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
