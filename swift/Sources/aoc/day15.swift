import Foundation

class Day15: Day {
    private typealias Grid = Array<Array<Int>>
    private let inputGrid: Grid

    init(_ input: String) {
        inputGrid = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n")
                .map({ $0.map({ Int(String($0))! }) })
    }

    func part1() throws -> Int {
        let path = try aStar(inputGrid)
        #if false
        printGrid(inputGrid, highlights: Set(path))
        #endif
        return totalRiskLevel(inputGrid, path)
    }

    func part2() throws -> Int {
        let grid = enlargeGrid(inputGrid)
        let path = try aStar(grid)

        #if false
        printGrid(grid, highlights: Set(path))
        #endif

        return totalRiskLevel(grid, path)
    }

    private func totalRiskLevel(_ grid: Grid, _ path: Array<Point>) -> Int {
        path.filter({ $0.x != 0 || $0.y != 0 }).map({ grid[$0.y][$0.x] }).reduce(0, +)
    }

    /// Borrowed from https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode
    private func aStar(_ grid: Grid) throws -> Array<Point> {
        func h(_ point: Point) -> Int {
            grid[point.y][point.x]
        }

        func d(_ current: Point, _ neighbor: Point) -> Int {
            h(neighbor)
        }

        let start = Point(x: 0, y: 0)
        let goal = Point(x: grid[0].count - 1, y: grid.count - 1)

        var openSet = [h(start): Set([start])]
        var cameFrom = [Point: Point]()

        var gScore = [start: 0]
        var fScore = [start: h(start)]

        func reconstructPath(_ current: Point) -> Array<Point> {
            var current = current
            var totalPath = [current]
            while let next = cameFrom[current] {
                current = next
                totalPath.append(current)
            }
            return totalPath.reversed()
        }

        func findPointWithLowestFScoreValue() -> (Int, Point) {
            var lowestScore = Int.max
            var lowestPoint : Point?
            for (score, set) in openSet {
                if score < lowestScore && !set.isEmpty {
                    lowestScore = score
                    lowestPoint = set.first
                }
            }
            return (lowestScore, lowestPoint!)
        }

        func neighbors(_ point: Point) -> Array<Point> {
            var out = Array<Point>()
            if point.x > 0 {
                out.append(Point(x: point.x - 1, y: point.y))
            }
            if point.x < goal.x {
                out.append(Point(x: point.x + 1, y: point.y))
            }
            if point.y > 0 {
                out.append(Point(x: point.x, y: point.y - 1))
            }
            if point.y < goal.y {
                out.append(Point(x: point.x, y: point.y + 1))
            }
            return out
        }

        func removeFromOpenSet(_ fScore: Int, _ current: Point) {
            openSet[fScore]?.remove(current)
        }

        func ensureInOpenSet(previousFScore: Int?, newFScore: Int, point: Point) {
            if previousFScore != nil {
                openSet[previousFScore!]?.remove(point)
            }
            if openSet[newFScore] != nil {
                openSet[newFScore]?.insert(point)
            } else {
                openSet[newFScore] = Set([point])
            }
        }

        func isOpenSetNotEmpty() -> Bool {
            for (_, set) in openSet {
                if !set.isEmpty {
                    return true
                }
            }
            return false
        }

        while isOpenSetNotEmpty() {
            let (lowestFScore, current) = findPointWithLowestFScoreValue()
            if current == goal {
                return reconstructPath(current)
            }

            removeFromOpenSet(lowestFScore, current)
            for neighbor in neighbors(current) {
                let tentativeGScore = gScore[current]! + d(current, neighbor)
                if tentativeGScore < (gScore[neighbor] ?? Int.max) {
                    cameFrom[neighbor] = current
                    gScore[neighbor] = tentativeGScore
                    let currentNeighborFScore = fScore[neighbor]
                    let newNeighborFScore = tentativeGScore + h(neighbor)
                    fScore[neighbor] = newNeighborFScore
                    ensureInOpenSet(previousFScore: currentNeighborFScore, newFScore: newNeighborFScore, point: neighbor)
                }
            }
        }

        throw NoPathFound()
    }

    private func enlargeGrid(_ inputGrid: Grid) -> Grid {
        var out = Grid()

        for gridY in 0...4 {
            for row in inputGrid {
                var newRow = Array<Int>()
                for gridX in 0...4 {
                    for n in row {
                        var newN = n + gridX + gridY
                        if newN > 9 {
                            newN -= 9
                        }
                        newRow.append(newN)
                    }
                }
                out.append(newRow)
            }
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
