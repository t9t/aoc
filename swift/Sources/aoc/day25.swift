import Foundation

class Day25: Day {
    private typealias Grid = Array<Array<Character>>

    private let inputGrid: Grid

    init(_ input: String) {
        inputGrid = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map({ Array<Character>($0) })
    }

    func part1() -> Int {
        let rowLength = inputGrid[0].count
        let maxX = rowLength - 1, maxY = inputGrid.count - 1
        var grid = inputGrid

        func step(left: Bool) -> Bool {
            let herd: Character = left ? ">" : "v"
            var nextGrid = grid, moved = false
            for (y, row) in grid.enumerated() {
                for (x, c) in row.enumerated() {
                    if c == herd {
                        let nx = left ? (x == maxX ? 0 : x + 1) : x
                        let ny = left ? y : (y == maxY ? 0 : y + 1)
                        if grid[ny][nx] == "." {
                            nextGrid[ny][nx] = herd
                            nextGrid[y][x] = "."
                            moved = true
                        }
                    }
                }
            }
            grid = nextGrid
            return moved
        }

        var steps = 0
        while true {
            steps += 1
            let movedLeft = step(left: true)
            if !step(left: false) && !movedLeft {
                return steps
            }
        }
    }

    func part2() -> Int {
        return 1337
    }
}
