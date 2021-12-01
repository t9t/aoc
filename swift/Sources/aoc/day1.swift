import Foundation

func day1part1(_ input: String) -> Int {
    let lines = input.split(separator: "\n")
    var items = Array<Int>()
    for line in lines {
        items.append(Int(line)!)
    }

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