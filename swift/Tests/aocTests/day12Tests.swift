import XCTest
import class Foundation.Bundle

@testable import aoc

final class day12Tests: XCTestCase {
    private let inputSmall = """
                             start-A
                             start-b
                             A-c
                             A-b
                             b-d
                             A-end
                             b-end
                             """

    private let inputSlightlyLarger = """
                                      dc-end
                                      HN-start
                                      start-kj
                                      dc-start
                                      dc-HN
                                      LN-dc
                                      HN-end
                                      kj-sa
                                      kj-HN
                                      kj-dc
                                      """

    private let inputEvenLarger = """
                                  fs-end
                                  he-DX
                                  fs-he
                                  start-DX
                                  pj-DX
                                  end-zg
                                  zg-sl
                                  zg-pj
                                  pj-he
                                  RW-he
                                  fs-DX
                                  pj-RW
                                  zg-RW
                                  start-pj
                                  he-WI
                                  zg-he
                                  pj-fs
                                  start-RW
                                  """

    func testPart1_Small() throws {
        let n = Day12(inputSmall).part1()
        XCTAssertEqual(n, 10)
    }

    func testPart1_SlightlyLarger() throws {
        let n = Day12(inputSlightlyLarger).part1()
        XCTAssertEqual(n, 19)
    }

    func testPart1_EvenLarger() throws {
        let n = Day12(inputEvenLarger).part1()
        XCTAssertEqual(n, 226)
    }

    func testPart2() throws {
        let n = Day12(input).part2()
        XCTAssertEqual(n, 1337)
    }
}
