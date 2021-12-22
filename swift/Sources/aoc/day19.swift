import Foundation

class Day19: Day {
    private typealias Orientation = (Position) -> Position

    private let inputPositions: Array<Set<Position>>
    private let orientations: Array<Orientation>

    init(_ input: String) {
        var inputPositions = Array<Set<Position>>()
        var currentScanner = Set<Position>()
        for line in input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n") {
            if line.hasPrefix("---") {
                if !currentScanner.isEmpty {
                    inputPositions.append(currentScanner)
                    currentScanner = Set<Position>()
                }
                continue
            }
            let parts = line.split(separator: ",")
            currentScanner.insert(Position(x: Int(parts[0])!, y: Int(parts[1])!, z: Int(parts[2])!))
        }
        inputPositions.append(currentScanner)
        self.inputPositions = inputPositions

        orientations = [
            ({ pos in Day19.newPos(+pos.x, +pos.y, +pos.z) }),
            ({ pos in Day19.newPos(+pos.y, +pos.z, +pos.x) }),
            ({ pos in Day19.newPos(+pos.z, +pos.x, +pos.y) }),
            ({ pos in Day19.newPos(+pos.z, +pos.y, -pos.x) }),
            ({ pos in Day19.newPos(+pos.y, +pos.x, -pos.z) }),
            ({ pos in Day19.newPos(+pos.x, +pos.z, -pos.y) }),

            ({ pos in Day19.newPos(+pos.x, -pos.y, -pos.z) }),
            ({ pos in Day19.newPos(+pos.y, -pos.z, -pos.x) }),
            ({ pos in Day19.newPos(+pos.z, -pos.x, -pos.y) }),
            ({ pos in Day19.newPos(+pos.z, -pos.y, +pos.x) }),
            ({ pos in Day19.newPos(+pos.y, -pos.x, +pos.z) }),
            ({ pos in Day19.newPos(+pos.x, -pos.z, +pos.y) }),

            ({ pos in Day19.newPos(-pos.x, +pos.y, -pos.z) }),
            ({ pos in Day19.newPos(-pos.y, +pos.z, -pos.x) }),
            ({ pos in Day19.newPos(-pos.z, +pos.x, -pos.y) }),
            ({ pos in Day19.newPos(-pos.z, +pos.y, +pos.x) }),
            ({ pos in Day19.newPos(-pos.y, +pos.x, +pos.z) }),
            ({ pos in Day19.newPos(-pos.x, +pos.z, +pos.y) }),

            ({ pos in Day19.newPos(-pos.x, -pos.y, +pos.z) }),
            ({ pos in Day19.newPos(-pos.y, -pos.z, +pos.x) }),
            ({ pos in Day19.newPos(-pos.z, -pos.x, +pos.y) }),
            ({ pos in Day19.newPos(-pos.z, -pos.y, -pos.x) }),
            ({ pos in Day19.newPos(-pos.y, -pos.x, -pos.z) }),
            ({ pos in Day19.newPos(-pos.x, -pos.z, -pos.y) }),
        ]
    }

    private static func newPos(_ x: Int, _ y: Int, _ z: Int) -> Position {
        Position(x: x, y: y, z: z)
    }

    func part1() -> Int {
        doTheThing().0
    }

    func part2() -> Int {
        // TODO: speed up
        doTheThing().1
    }

    func doTheThing() -> (Int, Int) {
        var beacons: Set<Position> = Set(inputPositions[0])
        var remaining = Array<Set<Position>>(inputPositions[1...])
        var locations: Set<Position> = Set([Position(x: 0, y: 0, z: 0)])

        while !remaining.isEmpty {
            for (i, d) in remaining.enumerated() {
                let result = doScannersMatch(beacons, d)
                if result != nil {
                    let (offset, orient) = result!
                    locations.insert(offset)

                    var others = Set<Position>()
                    for b in d {
                        others.insert(orient(b) + offset)
                    }
                    beacons = beacons.union(others)
                    remaining.remove(at: i)
                    break
                }
            }
        }

        let maxDistance = locations.flatMap({ l in locations.map({ $0 - l }) }).map({ abs($0.x) + abs($0.y) + abs($0.z) }).max()!

        return (beacons.count, maxDistance)
    }

    private func doScannersMatch(_ scanner1: Set<Position>, _ scanner2: Set<Position>) -> (Position, Orientation)? {
        let offsets1 = makeOffsetTable(scanner1)
        for (orientation, beacons) in allPossibleOrientations(scanner2) {
            let offsets2 = makeOffsetTable(beacons)

            for (x, xof) in offsets1 {
                for (y, yof) in offsets2 {
                    var count = 0
                    for r in yof {
                        if xof.contains(r) {
                            count += 1
                            if count >= 12 {
                                return (x - y, orientation)
                            }
                        }
                    }
                }
            }
        }
        return nil
    }

    private func makeOffsetTable(_ positions: Set<Position>) -> Dictionary<Position, Set<Position>> {
        var ret = Dictionary<Position, Set<Position>>()
        for p in positions {
            var arr = Set<Position>()
            for q in positions {
                arr.insert(p - q)
            }
            ret[p] = arr
        }
        return ret
    }

    private func allPossibleOrientations(_ positions: Set<Position>) -> Array<(Orientation, Set<Position>)> {
        var ret = Array<(Orientation, Set<Position>)>()
        for orient in orientations {
            var arr = Set<Position>()
            for pos in positions {
                arr.insert(orient(pos))
            }
            ret.append((orient, arr))
        }
        return ret
    }

    private struct Position: Equatable, Hashable, CustomStringConvertible {
        let x: Int, y: Int, z: Int

        func tuple() -> (Int, Int, Int) {
            (x, y, z)
        }

        var description: String {
            "(\(x),\(y),\(z))"
        }

        static func +(lhs: Position, rhs: Position) -> Position {
            Position(x: lhs.x + rhs.x, y: lhs.y + rhs.y, z: lhs.z + rhs.z)
        }

        static func -(lhs: Position, rhs: Position) -> Position {
            Position(x: lhs.x - rhs.x, y: lhs.y - rhs.y, z: lhs.z - rhs.z)
        }

        static func fromTuple(_ tup: (Int, Int, Int)) -> Position {
            Position(x: tup.0, y: tup.1, z: tup.2)
        }
    }
}
