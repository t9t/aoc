import Foundation

class Day22: Day {
    internal typealias Range = ClosedRange<Int>
    private let rebootSteps: Array<RebootStep>
    private let inputLines: Array<String>

    init(_ input: String) {
        func parseRange(_ s: Substring) -> Range {
            let rangeParts = s.split(separator: "=")[1].components(separatedBy: "..")
            return Int(rangeParts[0])! ... Int(rangeParts[1])!
        }

        inputLines = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map(String.init)
        rebootSteps = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n")
                .map({ $0.split(separator: " ") })
                .map({ stateAndRanges in
                    let ranges = stateAndRanges[1].split(separator: ",")
                    return RebootStep(state: stateAndRanges[0] == "on", cuboid: Cuboid(x: parseRange(ranges[0]), y: parseRange(ranges[1]), z: parseRange(ranges[2])))
                })
    }

    func part1() -> Int {
        func isOutOfRange(_ n: Int) -> Bool {
            n < -50 || n > 50
        }

        func isEntirelyOutOfRange(_ r: Range) -> Bool {
            let l = r.lowerBound
            let u = r.upperBound
            return l < -50 && u < -50 || l > 50 && u > 50
        }

        func isEntirelyOutOfRange(_ cuboid: Cuboid) -> Bool {
            isEntirelyOutOfRange(cuboid.x) || isEntirelyOutOfRange(cuboid.y) || isEntirelyOutOfRange(cuboid.z)
        }

        var onCubes = Set<Position>()
        for step in rebootSteps {
            let cuboid = step.cuboid
            if isEntirelyOutOfRange(cuboid) {
                continue
            }
            for x in cuboid.x {
                if isOutOfRange(x) {
                    continue
                }
                for y in cuboid.y {
                    if isOutOfRange(y) {
                        continue
                    }
                    for z in cuboid.z {
                        if isOutOfRange(z) {
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

    private struct Bla: Hashable, Equatable {
        let x0: Int, x1: Int, y0: Int, y1: Int, z0: Int, z1: Int
    }

    func part2() -> Int {
        var counts = [Cuboid: Int]()

        for step in rebootSteps {
            let cuboid = step.cuboid
            var newCounts = counts

            for (other, sign) in counts {
                let lowerX = max(cuboid.x.lowerBound, other.x.lowerBound)
                let upperX = min(cuboid.x.upperBound, other.x.upperBound)

                let lowerY = max(cuboid.y.lowerBound, other.y.lowerBound)
                let upperY = min(cuboid.y.upperBound, other.y.upperBound)

                let lowerZ = max(cuboid.z.lowerBound, other.z.lowerBound)
                let upperZ = min(cuboid.z.upperBound, other.z.upperBound)
                if lowerX <= upperX && lowerY <= upperY && lowerZ <= upperZ {
                    let intersection = Cuboid(x: lowerX...upperX, y: lowerY...upperY, z: lowerZ...upperZ)
                    newCounts[intersection] = (newCounts[intersection] ?? 0) - sign
                }
            }

            if step.state {
                newCounts[cuboid] = (newCounts[cuboid] ?? 0) + 1
            }

            counts = newCounts
        }

        return counts.map({ (n, sgn) in (n.x.upperBound - n.x.lowerBound + 1) * (n.y.upperBound - n.y.lowerBound + 1) * (n.z.upperBound - n.z.lowerBound + 1) * sgn }).reduce(0, +)
    }

    private struct RebootStep {
        let state: Bool, cuboid: Cuboid
    }

    private struct Position: Hashable, Equatable {
        let x: Int, y: Int, z: Int
    }

    private struct Cuboid: Hashable, Equatable {
        let x: Range, y: Range, z: Range

        func fullyInside(_ other: Cuboid) -> Bool {
            x.lowerBound >= other.x.lowerBound && x.upperBound <= other.x.upperBound
                    && y.lowerBound >= other.y.lowerBound && y.upperBound <= other.y.upperBound
                    && z.lowerBound >= other.z.lowerBound && z.upperBound <= other.z.upperBound
        }
    }
}
