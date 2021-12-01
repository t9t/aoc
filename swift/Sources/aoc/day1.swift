import Foundation

func day1part1(_ input: String) -> Int {
    let items = parseInput(input)

    var n = 0
    var prev = Int.max
    for item in items {
        if item > prev {
            n += 1
        }
        prev = item
    }
    return n
}

func day1part2(_ input: String) -> Int {
    let items = parseInput(input)

    var n = 0
    var prev = Int.max
    for i in 0...items.count - 3 {
        let window = items[i] + items[i + 1] + items[i + 2]
        if window > prev {
            n += 1
        }
        prev = window
    }
    return n
}

func parseInput(_ input: String) -> Array<Int> {
    let lines = input.split(separator: "\n")
    var items = Array<Int>()
    for line in lines {
        items.append(Int(line)!)
    }
    return items
}
