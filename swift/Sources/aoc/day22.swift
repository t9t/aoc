import Foundation

class Day22: Day {
    private typealias Range = ClosedRange<Int>
    private let rebootSteps: Array<RebootStep>

    init(_ input: String) {
        func parseRange(_ s: Substring) -> Range {
            let rangeParts = s.split(separator: "=")[1].components(separatedBy: "..")
            return Int(rangeParts[0])!...Int(rangeParts[1])!
        }

        rebootSteps = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n")
                .map({ $0.split(separator: " ") })
                .map({ stateAndRanges in
                    let ranges = stateAndRanges[1].split(separator: ",")
                    return RebootStep(state: stateAndRanges[0] == "on", x: parseRange(ranges[0]), y: parseRange(ranges[1]), z: parseRange(ranges[2]))
                })
    }

    func part1() -> Int {
        func isEntirelyOutOfRange(_ r: Range) -> Bool {
            let l = r.lowerBound
            let u = r.upperBound
            return l < -50 && u < -50 || l > 50 && u > 50
        }

        func isEntirelyOutOfRange(_ step: RebootStep) -> Bool {
            isEntirelyOutOfRange(step.x) || isEntirelyOutOfRange(step.y) || isEntirelyOutOfRange(step.z)
        }

        var onCubes = Set<Position>()
        for step in rebootSteps {
            if isEntirelyOutOfRange(step) {
                continue
            }
            for x in step.x {
                if x < -50 || x > 50 {
                     continue
                }
                for y in step.y {
                    if y < -50 || y > 50 {
                        continue
                    }
                    for z in step.z {
                        if z < -50 || z > 50 {
                            continue
                        }
                        let p = Position(x: x, y: y, z: z)
                        if step.state {
                            onCubes.insert(p)
                        } else {
                            onCubes.remove(p)
                        }
                    }
                }
            }
        }
        return onCubes.count
    }

    func part2() -> Int {
        return 1337
    }

    private struct RebootStep {
        let state: Bool, x: Range, y: Range, z: Range
    }

    private struct Position: Hashable, Equatable {
        let x: Int, y: Int, z: Int
    }
}
