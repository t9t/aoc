import Foundation

class Day9: Day {
    private let rows: Array<Array<Int>>

    init(_ input: String) {
        rows = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n")
                .map({ $0.map({ Int(String($0))! }) })
    }

    func part1() -> Int {
        var riskLevelsSum = 0
        let maxY = rows.count - 1, maxX = rows[0].count - 1
        for (y, row) in rows.enumerated() {
            for (x, n) in row.enumerated() {
                let left = x > 0 ? row[x - 1] : Int.max
                let right = x < maxX ? row[x + 1] : Int.max
                let above = y > 0 ? rows[y - 1][x] : Int.max
                let below = y < maxY ? rows[y + 1][x] : Int.max

                if left > n && right > n && above > n && below > n {
                    riskLevelsSum += 1 + n
                }
            }
        }
        return riskLevelsSum
    }

    func part2() -> Int {
        var nextBasinNumber = 1
        var basinMap: [Point: Int] = [:]

        for (y, row) in rows.enumerated() {
            for (x, n) in row.enumerated() {
                if n == 9 {
                    // 9 is never part of any basin
                    continue
                }

                let point = Point(x: x, y: y)
                if basinMap[point] != nil {
                    // Already filled in earlier, ignore
                    continue
                }

                // Must be a new basin here
                basinMap = fillIn(point: point, basinNumber: nextBasinNumber, inputBasinMap: basinMap)
                nextBasinNumber += 1
            }
        }

        #if false
        prettyPrint(basinMap: basinMap)
        #endif

        var basinSizes = Array<Int>()
        for basinNumber in 1...nextBasinNumber - 1 {
            basinSizes.append(basinMap.filter({ $1 == basinNumber }).count)
        }
        basinSizes.sort()
        basinSizes.reverse()

        return basinSizes[0] * basinSizes[1] * basinSizes[2]
    }

    private func fillIn(point: Point, basinNumber: Int, inputBasinMap: [Point: Int]) -> [Point: Int] {
        var basinMap = inputBasinMap
        var next = Array<Point>()
        next.append(point)
        let maxY = rows.count - 1, maxX = rows[0].count - 1

        func maybeAppend(x: Int, y: Int) {
            if x >= 0 && x <= maxX && y >= 0 && y <= maxY && rows[y][x] != 9 {
                let point = Point(x: x, y: y)
                if basinMap[point] == nil {
                    next.append(point)
                }
            }
        }

        while !next.isEmpty {
            let current = next.popLast()!
            basinMap[current] = basinNumber
            let x = current.x, y = current.y
            maybeAppend(x: x - 1, y: y) // Left
            maybeAppend(x: x + 1, y: y) // Right
            maybeAppend(x: x, y: y - 1) // Above
            maybeAppend(x: x, y: y + 1) // Below
        }
        return basinMap
    }

    private func prettyPrint(basinMap: [Point: Int]) {
        print("")
        for (y, row) in rows.enumerated() {
            print("    ", terminator: "")
            for (x, n) in row.enumerated() {
                if n == 9 {
                    print("\u{001b}[0m\(n)", terminator: "")
                } else {
                    print("\u{001b}[7m\u{001b}[38;5;\(basinMap[Point(x: x, y: y)]! % 256)m\(n)", terminator: "")
                }
            }
            print("\u{001b}[0m")
        }
        print("")
    }

    private struct Point: Hashable, Equatable, CustomStringConvertible {
        let x: Int, y: Int

        var description: String {
            "\(x)x\(y)"
        }
    }
}
