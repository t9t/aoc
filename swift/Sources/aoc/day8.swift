import Foundation

class Day8: Day {
    let notes: Array<(Array<Substring>, Array<Substring>)>

    init(_ input: String) {
        notes = input.split(separator: "\n")
                .map { $0.components(separatedBy: " | ") }
                .map { ($0[0].split(separator: " "), $0[1].split(separator: " ")) }
    }

    func part1() -> Int {
        notes.map { $0.1 }
                .joined()
                .filter { value in
                    value.count == 2 /*1*/ || value.count == 4 /*4*/ || value.count == 3 /*7*/ || value.count == 7 /*8*/
                }
                .count
    }

    func part2() -> Int {
        return 1337
    }
}
