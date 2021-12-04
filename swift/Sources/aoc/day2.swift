import Foundation

class Day2: Day {
    let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.split(separator: "\n")
    }

    func part1() -> Int {
        var x = 0, depth = 0
        for line in inputLines {
            let parts = line.split(separator: " ")
            let direction = parts[0], amount = Int(parts[1])!

            if direction == "forward" {
                x += amount
            } else if direction == "down" {
                depth += amount
            } else if direction == "up" {
                depth -= amount
            }
        }

        return x * depth
    }

    func part2() -> Int {
        var x = 0, depth = 0, aim = 0
        for line in inputLines {
            let parts = line.split(separator: " ")
            let direction = parts[0], amount = Int(parts[1])!

            if direction == "forward" {
                x += amount
                depth += (aim * amount)
            } else if direction == "down" {
                aim += amount
            } else if direction == "up" {
                aim -= amount
            }
        }

        return x * depth
    }
}
