import XCTest
import class Foundation.Bundle

@testable import aoc

final class integrationTests: XCTestCase {
    func test2021() throws {
        let results = try loadResults()
        for dayNum in Days.getAllDayNumbers() {
            for part in 1...2 {
                let id = ResultId(year: 2021, day: dayNum, part: part)
                if let result = results[id] {
                    let input = try String(contentsOfFile: "../input/2021/\(dayNum).txt")
                    let day = try Days.get(num: dayNum, input: input)

                    let output = try (part == 1 ? day.part1 : day.part2)()
                    XCTAssertEqual(result, output, "\(id)")
                } else {
                    XCTFail("No result for day \(dayNum) part \(part)")
                }
            }
        }
    }

    private func loadResults() throws -> [ResultId: String] {
        var results: [ResultId: String] = [:]
        for line in try String(contentsOfFile: "../input/2021/results.txt")
                .trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: "\n") {
            let idAndResult = line.components(separatedBy: ": ")
            let idParts = idAndResult[0].split(separator: "-")
            results[ResultId(year: Int(idParts[0])!, day: Int(idParts[1])!, part: Int(idParts[2])!)] = idAndResult[1]
        }
        return results
    }

    private struct ResultId: Hashable {
        let year: Int, day: Int, part: Int
    }
}
