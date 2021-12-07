import Foundation

class Day5: Day {
    private let lineSegments: Array<LineSegment>
    private var grid: [Point: Int] = [:]

    init(_ input: String) {
        lineSegments = input.split(separator: "\n")
                .map { line in
                    let fromTo = line.components(separatedBy: " -> ")
                    let from = fromTo[0].split(separator: ",")
                    let to = fromTo[1].split(separator: ",")

                    return LineSegment(
                            from: Point(x: Int(from[0])!, y: Int(from[1])!),
                            to: Point(x: Int(to[0])!, y: Int(to[1])!))
                }
    }

    func part1() -> Int {
        drawLinesReturningCount(includeDiagonals: false)
    }

    func part2() -> Int {
        drawLinesReturningCount(includeDiagonals: true)
    }

    private func drawLinesReturningCount(includeDiagonals: Bool) -> Int {
        for segment in lineSegments {
            let diagonal = segment.isDiagonal()
            if !diagonal || (diagonal && includeDiagonals) {
                incrementAll(segment.points())
            }
        }
        #if false
        printGrid()
        #endif
        return countAtLeastTwoOverlaps()
    }

    private func countAtLeastTwoOverlaps() -> Int {
        var count = 0
        for (_, val) in grid {
            if val >= 2 {
                count += 1
            }
        }
        return count
    }

    private func incrementAll(_ points: [Point]) {
        for point in points {
            increment(point)
        }
    }

    private func increment(_ point: Point) {
        if let curr = grid[point] {
            grid[point] = curr + 1
        } else {
            grid[point] = 1
        }
    }


    private func printGrid() {
        var minX = Int.max, maxX = Int.min, minY = Int.max, maxY = Int.min
        for (key, _) in grid {
            minX = min(key.x, minX)
            maxX = max(key.x, maxX)
            minY = min(key.y, minY)
            maxY = max(key.y, maxY)
        }
        for y in minY...maxY {
            for x in minX...maxX {
                var char = "."
                if let n = grid[Point(x: x, y: y)] {
                    char = "\(n)"
                }
                print(char, terminator: "")
            }
            print("")
        }
    }

    internal struct Point: Hashable, Equatable {
        let x: Int, y: Int
    }

    internal struct LineSegment {
        let from: Point
        let to: Point

        internal func isDiagonal() -> Bool {
            from.x != to.x && from.y != to.y
        }

        internal func points() -> Array<Point> {
            let totalDx = to.x - from.x
            let totalDy = to.y - from.y

            let dx = totalDx == 0 ? 0 : (totalDx < 0 ? -1 : 1)
            let dy = totalDy == 0 ? 0 : (totalDy < 0 ? -1 : 1)

            var points = Array<Point>()
            var x = from.x - dx, y = from.y - dy
            repeat {
                x += dx
                y += dy
                points.append(Point(x: x, y: y))
            } while x != to.x || y != to.y
            return points
        }
    }
}
