import Foundation

class Day5: Day {
    let lineSegments: Array<LineSegment>
    var grid: [Point: Int] = [:]

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

    func drawLinesReturningCount(includeDiagonals: Bool) -> Int {
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

    func countAtLeastTwoOverlaps() -> Int {
        var count = 0
        for (_, val) in grid {
            if val >= 2 {
                count += 1
            }
        }
        return count
    }

    func incrementAll(_ points: [Point]) {
        for point in points {
            increment(point)
        }
    }

    func increment(_ point: Point) {
        if let curr = grid[point] {
            grid[point] = curr + 1
        } else {
            grid[point] = 1
        }
    }


    func printGrid() {
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

    struct Point: Hashable, Equatable {
        let x: Int, y: Int
    }

    struct LineSegment {
        let from: Point
        let to: Point

        func isDiagonal() -> Bool {
            from.x != to.x && from.y != to.y
        }

        func points() -> Array<Point> {
            var points = Array<Point>()
            if (from.x == to.x) { // Horizontal
                let minY = min(from.y, to.y), maxY = max(from.y, to.y)
                for y in minY...maxY {
                    points.append(Point(x: from.x, y: y))
                }
                return points
            }
            if (from.y == to.y) { // Vertical
                let minX = min(from.x, to.x), maxX = max(from.x, to.x)
                for x in minX...maxX {
                    points.append(Point(x: x, y: from.y))
                }
                return points
            }
            // Diagonal
            var dx = to.x - from.x
            var dy = to.y - from.y
            let steps = abs(dx)
            dx = dx > 0 ? 1 : -1
            dy = dy > 0 ? 1 : -1

            for step in 0...steps {
                let x = from.x + (dx * step)
                let y = from.y + (dy * step)
                points.append(Point(x: x, y: y))
            }
            return points
        }
    }
}
