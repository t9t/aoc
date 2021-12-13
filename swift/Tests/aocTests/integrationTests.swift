import XCTest
import class Foundation.Bundle

@testable import aoc

final class integrationTests: XCTestCase {
    func test2021() throws {
        let results = try loadResults()
        for result in results {
            if result.year != 2021 {
                XCTFail("Year \(result.year) is not 2021")
                continue
            }
            if result.part != 1 && result.part != 2 {
                XCTFail("Part \(result.part) is not 1 or 2")
                continue
            }

            let input = try String(contentsOfFile: "../input/2021/\(result.day).txt")
            let day = Days.get(num: result.day, input: input)

            let output = try (result.part == 1 ? day.part1 : day.part2)()
            XCTAssertEqual(result.result, output, "\(result)")
        }
    }

    private func loadResults() throws -> Array<Result> {
        try String(contentsOfFile: "../input/2021/results.txt")
                .trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: "\n")
                .map({ $0.components(separatedBy: ": ") })
                .map({ ($0[0].split(separator: "-"), $0[1]) })
                .map({ Result(year: Int($0.0[0])!, day: Int($0.0[1])!, part: Int($0.0[2])!, result: $0.1) })
    }

    private struct Result {
        let year: Int, day: Int, part: Int, result: String
    }
}
