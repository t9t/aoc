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

    func part2() -> Int {
        var onCuboids = Array<Cuboid2>()

        for step in rebootSteps {
            let cuboid = Cuboid2(x: step.cuboid.x, y: step.cuboid.y, z: step.cuboid.z, excluded: [])

            for (i, _) in onCuboids.enumerated() {
                onCuboids[i].exclude(cuboid)
            }

            if step.state {
                onCuboids.append(cuboid)
            }
        }

        return onCuboids.map({ $0.volume() }).reduce(0, +)
    }

    private struct RebootStep {
        let state: Bool, cuboid: Cuboid
    }

    private struct Position: Hashable, Equatable {
        let x: Int, y: Int, z: Int
    }

    private struct Cuboid: Hashable, Equatable {
        let x: Range, y: Range, z: Range
    }

    private struct Cuboid2: Hashable, Equatable {
        let x: Range, y: Range, z: Range
        var excluded: Array<Cuboid2>

        func volume() -> Int {
            let total = (x.upperBound - x.lowerBound + 1) * (y.upperBound - y.lowerBound + 1) * (z.upperBound - z.lowerBound + 1)
            let excludedVolume = excluded.map({ $0.volume() }).reduce(0, +)
            return total - excludedVolume
        }

        mutating func exclude(_ other: Cuboid2) {
            let lowerX = max(x.lowerBound, other.x.lowerBound)
            let upperX = min(x.upperBound, other.x.upperBound)

            let lowerY = max(y.lowerBound, other.y.lowerBound)
            let upperY = min(y.upperBound, other.y.upperBound)

            let lowerZ = max(z.lowerBound, other.z.lowerBound)
            let upperZ = min(z.upperBound, other.z.upperBound)
            if lowerX <= upperX && lowerY <= upperY && lowerZ <= upperZ {
                let intersection = Cuboid2(x: lowerX...upperX, y: lowerY...upperY, z: lowerZ...upperZ, excluded: [])
                for (i, _) in excluded.enumerated() {
                    excluded[i].exclude(intersection)
                }
                excluded.append(intersection)
            }
        }
    }
}
