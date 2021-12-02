import Foundation

class Day2 {
    let input: String

    init(_ input: String) {
        self.input = input
    }

    func part1() -> Int {
        let lines = input.split(separator: "\n")

        var x = 0, depth = 0
        for line in lines {
            let parts = line.split(separator: " ")
            let direction = parts[0]
            let amount = Int(parts[1])!

            switch direction {
            case "forward":
                x += amount
            case "down":
                depth += amount
            case "up":
                depth -= amount
            default:
                break
            }
        }

        return x * depth
    }

    func part2() -> Int {
        return 0
    }
}
