import XCTest
import class Foundation.Bundle

@testable import aoc

final class day1Tests: XCTestCase {
    func testPart1() throws {
        let n = day1part1()
        XCTAssertEqual(n, 7)
    }
}
