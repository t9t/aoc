import Foundation

class TemplateDay: Day {
    private let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.split(separator: "\n")
    }

    func part1() -> Int {
        return 42
    }

    func part2() -> Int {
        return 1337
    }
}
