import Foundation

class Day15: Day {
    private typealias Grid = Array<Array<Int>>
    private let inputGrid: Grid

    init(_ input: String) {
        inputGrid = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n")
                .map({ $0.map({ Int(String($0))! }) })
    }

    func part1() throws -> Int {
        let path = try aStar()
        #if false
        printGrid(inputGrid, highlights: Set(path))
        #endif
        return path.filter({ $0.x != 0 || $0.y != 0 }).map({ inputGrid[$0.y][$0.x] }).reduce(0, +)
    }

    func part2() -> Int {
        return 1337
    }

    /// Borrowed from https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode
    private func aStar() throws -> Array<Point> {
        func h(_ point: Point) -> Int {
            inputGrid[point.y][point.x]
        }

        func d(_ current: Point, _ neighbor: Point) -> Int {
            h(neighbor)
        }

        let start = Point(x: 0, y: 0)
        let goal = Point(x: inputGrid[0].count - 1, y: inputGrid.count - 1)

        var openSet = Set([start])
        var cameFrom = [Point: Point]()

        var gScore = [start: 0]
        var fScore = [start : h(start)]

        func reconstructPath(_ current: Point) -> Array<Point> {
            var current = current
            var totalPath = [current]
            while let next = cameFrom[current] {
                current = next
                totalPath.append(current)
            }
            return totalPath.reversed()
        }

        func findPointWithLowestFScoreValue() -> Point {
            var lowestScore = Int.max
            var lowest: Point? = nil
            for point in openSet {
                let level = fScore[point] ?? Int.max
                if level <= lowestScore {
                    lowestScore = level
                    lowest = point
                }
            }
            return lowest!
        }

        while !openSet.isEmpty {
            let current = findPointWithLowestFScoreValue()
            if current == goal {
                return reconstructPath(current)
            }

            openSet.remove(current)
            for neighbor in neighbors(current) {
                let tentativeGScore = gScore[current]! + d(current, neighbor)
                if tentativeGScore < (gScore[neighbor] ?? Int.max) {
                    cameFrom[neighbor] = current
                    gScore[neighbor] = tentativeGScore
                    fScore[neighbor] = tentativeGScore + h(neighbor)
                    if !openSet.contains(neighbor) {
                        openSet.insert(neighbor)
                    }
                }
            }
        }

        throw NoPathFound()
    }

    private func neighbors(_ point: Point) -> Array<Point> {
        var out = Array<Point>()
        if point.x > 0 {
            out.append(Point(x: point.x - 1, y: point.y))
        }
        if point.x < inputGrid[0].count - 1 {
            out.append(Point(x: point.x + 1, y: point.y))
        }
        if point.y > 0 {
            out.append(Point(x: point.x, y: point.y - 1))
        }
        if point.y < inputGrid.count - 1 {
            out.append(Point(x: point.x, y: point.y + 1))
        }
        return out
    }

    private func printGrid(_ grid: Grid, highlights: Set<Point> = Set()) {
        for (y, row) in grid.enumerated() {
            for (x, n) in row.enumerated() {
                var color = 0
                if highlights.contains(Point(x: x, y: y)) {
                    color = 7
                }

                // \u001b[7m

                print("\u{001b}[\(color)m\(n)", terminator: "")
            }
            print("\u{001b}[0m")
        }
    }

    private struct Point: Hashable, Equatable, CustomStringConvertible {
        let x: Int, y: Int
        var description: String {
            "\(x)x\(y)"
        }
    }

    private class NoPathFound: Error {
    }
}
