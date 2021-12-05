import Foundation

class Day5: Day {
    let lineSegments: Array<LineSegment>

    init(_ input: String) {
        let inputLines = input.split(separator: "\n")
        var segments = Array<LineSegment>()
        for line in inputLines {
            let fromTo = line.components(separatedBy: " -> ")
            let from = fromTo[0].split(separator: ",")
            let to = fromTo[1].split(separator: ",")

            let segment = LineSegment(
                    from: Point(x: Int(from[0])!, y: Int(from[1])!),
                    to: Point(x: Int(to[0])!, y: Int(to[1])!))
            if segment.from.x == segment.to.x || segment.from.y == segment.to.y {
                segments.append(segment)
            }
        }
        lineSegments = segments
    }

    func part1() -> Int {
        var grid: [Point: Int] = [:]

        func increment(x: Int, y: Int) {
            let point = Point(x: x, y: y)
            if let curr = grid[point] {
                grid[point] = curr + 1
            } else {
                grid[point] = 1
            }
        }

        for segment in lineSegments {
            if segment.from.x == segment.to.x { // Horizontal
                let fromY = min(segment.from.y, segment.to.y)
                let toY = max(segment.from.y, segment.to.y)
                for y in fromY...toY {
                    increment(x: segment.from.x, y: y)
                }
            }
            if segment.from.y == segment.to.y { // Vertical
                let fromX = min(segment.from.x, segment.to.x)
                let toX = max(segment.from.x, segment.to.x)
                for x in fromX...toX {
                    increment(x: x, y: segment.from.y)
                }
            }
        }
        var count = 0
        for (_, val) in grid {
            if val >= 2 {
                count += 1
            }
        }
        return count
    }

    func part2() -> Int {
        return 1337
    }

    struct Point: Hashable, Equatable {
        let x: Int, y: Int
    }

    struct LineSegment {
        let from: Point
        let to: Point
    }
}
