import Foundation

class Day13: StringDay {
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

    func part1() -> String {
        let folded = fold(grid: inputGrid, fold: folds[0])
        #if false
        printGrid(folded, colorized: false)
        #endif

        return String(countOn(folded))
    }

    func part2() -> String {
        var grid = inputGrid
        for fold in folds {
            grid = self.fold(grid: grid, fold: fold)
        }
        #if false
        printGrid(grid, colorized: true)
        #endif

        return splitLetterGrids(grid: grid).map(toString)
                .map({ letters[$0] })
                .map({ $0 == nil ? "?" : $0! })
                .map(String.init)
                .joined()
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

    private func splitLetterGrids(grid: Grid) -> Array<Grid> {
        let letterCount = grid[0].count / 5
        var out = Array<Grid>()
        for i in 0...letterCount - 1 {
            let startX = i * 5
            var letterGrid = Grid()
            for row in grid {
                letterGrid.append(Array<Bool>(row[startX...startX + 4]))
            }
            out.append(letterGrid)
        }
        return out
    }

    private func toString(grid: Grid) -> String {
        grid.map({ row in row.map({ $0 ? "#" : "." }).joined() }).joined(separator: "\n")
    }

    private let letters: [String: Character] = [
        // TODO: need more inputs for more letters
        "###..\n#..#.\n###..\n#..#.\n#..#.\n###..": "B",
        ".##..\n#..#.\n#....\n#....\n#..#.\n.##..": "C",
        "####.\n#....\n###..\n#....\n#....\n#....": "F",
        ".##..\n#..#.\n#....\n#.##.\n#..#.\n.###.": "G",
        "#..#.\n#.#..\n##...\n#.#..\n#.#..\n#..#.": "K",
        "####.\n...#.\n..#..\n.#...\n#....\n####.": "Z",
        "n#####\n#...#\n#...#\n#...#\n#####\n.....\n.....": "0", // Test case
    ]

    private func printGrid(_ grid: Grid, colorized: Bool) {
        print("\u{001b}[0m")
        for row in grid {
            for on in row {
                print(colorized ? (on ? "\u{001b}[0m\u{001b}[7m#" : "\u{001b}[0m\u{001b}[30m.") : (on ? "#" : "."), terminator: "")
            }
            print("\u{001b}[0m")
        }
    }

    private enum Fold {
        case AlongX(x: Int)
        case AlongY(y: Int)
    }
}
