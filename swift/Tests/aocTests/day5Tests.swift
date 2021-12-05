import XCTest
import class Foundation.Bundle

@testable import aoc

final class day5Tests: XCTestCase {
    private let input = """
                        0,9 -> 5,9
                        8,0 -> 0,8
                        9,4 -> 3,4
                        2,2 -> 2,1
                        7,0 -> 7,4
                        6,4 -> 2,0
                        0,9 -> 2,9
                        3,4 -> 1,4
                        0,0 -> 8,8
                        5,5 -> 8,2
                        """

    func testPart1() throws {
        let n = Day5(input).part1()
        XCTAssertEqual(n, 5)
    }

    func testPart2() throws {
        let n = Day5(input).part2()
        XCTAssertEqual(n, 12)
    }

    func test_LineSegment_points_horizontal_leftToRight() throws {
        testLineSegmentPoints(from: Day5.Point(x: 1, y: 2), to: Day5.Point(x: 4, y: 2), expected: [
            Day5.Point(x: 1, y: 2),
            Day5.Point(x: 2, y: 2),
            Day5.Point(x: 3, y: 2),
            Day5.Point(x: 4, y: 2),
        ])
    }

    func test_LineSegment_points_horizontal_rightToLeft() throws {
        testLineSegmentPoints(from: Day5.Point(x: 4, y: 2), to: Day5.Point(x: 1, y: 2), expected: [
            Day5.Point(x: 1, y: 2),
            Day5.Point(x: 2, y: 2),
            Day5.Point(x: 3, y: 2),
            Day5.Point(x: 4, y: 2),
        ])
    }

    func test_LineSegment_points_vertical_topToBottom() throws {
        testLineSegmentPoints(from: Day5.Point(x: 2, y: 1), to: Day5.Point(x: 2, y: 4), expected: [
            Day5.Point(x: 2, y: 1),
            Day5.Point(x: 2, y: 2),
            Day5.Point(x: 2, y: 3),
            Day5.Point(x: 2, y: 4),
        ])
    }

    func test_LineSegment_points_vertical_bottomToTop() throws {
        testLineSegmentPoints(from: Day5.Point(x: 2, y: 4), to: Day5.Point(x: 2, y: 1), expected: [
            Day5.Point(x: 2, y: 1),
            Day5.Point(x: 2, y: 2),
            Day5.Point(x: 2, y: 3),
            Day5.Point(x: 2, y: 4),
        ])
    }

    func test_LineSegment_points_diagonal_topLeftToBottomRight() throws {
        testLineSegmentPoints(from: Day5.Point(x: 1, y: 2), to: Day5.Point(x: 4, y: 5), expected: [
            Day5.Point(x: 1, y: 2),
            Day5.Point(x: 2, y: 3),
            Day5.Point(x: 3, y: 4),
            Day5.Point(x: 4, y: 5),
        ])
    }

    func test_LineSegment_points_diagonal_topRightToBottomLeft() throws {
        testLineSegmentPoints(from: Day5.Point(x: 4, y: 2), to: Day5.Point(x: 1, y: 5), expected: [
            Day5.Point(x: 4, y: 2),
            Day5.Point(x: 3, y: 3),
            Day5.Point(x: 2, y: 4),
            Day5.Point(x: 1, y: 5),
        ])
    }

    func test_LineSegment_points_diagonal_bottomLeftToTopRight() throws {
        testLineSegmentPoints(from: Day5.Point(x: 1, y: 5), to: Day5.Point(x: 4, y: 2), expected: [
            Day5.Point(x: 1, y: 5),
            Day5.Point(x: 2, y: 4),
            Day5.Point(x: 3, y: 3),
            Day5.Point(x: 4, y: 2),
        ])
    }

    func test_LineSegment_points_diagonal_bottomRightToTopLeft() throws {
        testLineSegmentPoints(from: Day5.Point(x: 4, y: 5), to: Day5.Point(x: 1, y: 2), expected: [
            Day5.Point(x: 4, y: 5),
            Day5.Point(x: 3, y: 4),
            Day5.Point(x: 2, y: 3),
            Day5.Point(x: 1, y: 2),
        ])
    }

    func testLineSegmentPoints(from: Day5.Point, to: Day5.Point, expected: [Day5.Point]) {
        XCTAssertEqual(Day5.LineSegment(from: from, to: to).points(), expected)
    }
}
