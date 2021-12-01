import Foundation

func day1part1() -> Int {
    let items = [
        199,
        200,
        208,
        210,
        200,
        207,
        240,
        269,
        260,
        263]

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