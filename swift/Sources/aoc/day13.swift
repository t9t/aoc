import Foundation

class Day13: Day {
    private typealias Grid = Array<Array<Bool>>

    private let inputGrid: Grid
    private let folds: Array<Fold>

    init(_ input: String) {
        let coordsAndFolds = input.components(separatedBy: "\n\n")
        let inputCoords = coordsAndFolds[0].split(separator: "\n")
                .map({ $0.split(separator: ",") })
                .map({ (Int($0[0])!, Int($0[1])!) })
        let (maxX, maxY) = inputCoords.reduce((Int.min, Int.min), { (max($0.0, $1.0), max($0.1, $1.1)) })

        var grid = Grid(repeating: Array<Bool>(repeating: false, count: maxX + 1), count: maxY + 1)
        for coord in inputCoords {
            grid[coord.1][coord.0] = true
        }

        inputGrid = grid
        folds = coordsAndFolds[1].split(separator: "\n")
                .map({ $0.split(separator: "=") })
                .map({ ($0[0].hasSuffix("x"), Int($0[1])!) })
                .map({ $0.0 ? Fold.AlongX(x: $0.1) : Fold.AlongY(y: $0.1) })
    }

    func part1() -> Int {
        let folded = fold(grid: inputGrid, fold: folds[0])
        #if false
        printGrid(folded)
        #endif

        return countOn(folded)
    }

    func part2() -> Int {
        return 1337
    }

    private func fold(grid: Grid, fold: Fold) -> Grid {
        switch fold {
        case .AlongX(let x):
            return foldAlongX(grid: grid, x: x)
        case .AlongY(let y):
            return foldAlongY(grid: grid, y: y)
        }
    }

    private func foldAlongY(grid: Grid, y: Int) -> Grid {
        var top = grid[0...y - 1]
        let bottom = Grid(grid[(y + 1)...])
        for (by, row) in bottom.enumerated() {
            let ty = y - by - 1
            for (bx, on) in row.enumerated() {
                if on {
                    top[ty][bx] = true
                }
            }
        }
        return Grid(top)
    }

    private func foldAlongX(grid: Grid, x: Int) -> Grid {
        var out = Grid()
        for row in grid {
            var left = row[0...x - 1]
            let right = row[(x + 1)...]
            for (rx, on) in right.enumerated() {
                let lx = x - rx - 1
                if on {
                    left[lx] = true
                }
            }
            out.append(Array<Bool>(left))
        }

        return out
    }

    private func countOn(_ grid: Grid) -> Int {
        grid.map({ $0.filter({ $0 }).count }).reduce(0, +)
    }

    private func printGrid(_ grid: Grid) {
        for row in grid {
            for dot in row {
                print(dot ? "#" : ".", terminator: "")
            }
            print("")
        }
    }

    private enum Fold {
        case AlongX(x: Int)
        case AlongY(y: Int)
    }
}
