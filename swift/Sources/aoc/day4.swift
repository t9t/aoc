import Foundation

class Day4 {
    let numbers: Array<Int>
    var cards: Array<Array<Array<(Int, Bool)>>>

    init(_ input: String) {
        let lines = input.split(separator: "\n", omittingEmptySubsequences: false)
        numbers = lines[0].split(separator: ",").map {
            Int($0)!
        }
        cards = Array<Array<Array<(Int, Bool)>>>()
        var card = Array<Array<(Int, Bool)>>()
        for line in lines[2...] {
            if line.isEmpty {
                cards.append(card)
                card = Array<Array<(Int, Bool)>>()
                continue
            }

            var row = Array<(Int, Bool)>()
            let nums = line.split(separator: " ")
            for num in nums {
                row.append((Int(num.trimmingCharacters(in: .whitespaces))!, false))
            }
            card.append(row)
        }
        if !card.isEmpty {
            cards.append(card)
        }
    }

    func part1() throws -> Int {
        for num in numbers {
            mark(num)
            for card in cards {
                if wins(card) {
                    return unmarkedSum(card) * num
                }
            }
        }
        throw NoWinnerFound()
    }

    func part2() -> Int {
        return 1337
    }

    func mark(_ num: Int) {
        for c in 0...cards.count - 1 {
            let card: Array<Array<(Int, Bool)>> = cards[c]
            for r in 0...card.count - 1 {
                let row: Array<(Int, Bool)> = card[r]
                for i in 0...row.count - 1 {
                    let n = row[i]
                    if n.0 == num {
                        cards[c][r][i] = (n.0, true)
                    }
                }
            }
        }
    }

    func wins(_ card: Array<Array<(Int, Bool)>>) -> Bool {
        for row in card {
            var allHit = true
            for num in row {
                if !num.1 {
                    allHit = false
                    break
                }
            }
            if allHit {
                return true
            }
        }

        for i in 0...card[0].count - 1 {
            var allHit = true
            for row in card {
                if !row[i].1 {
                    allHit = false
                    break
                }
            }
            if allHit {
                return true
            }
        }
        return false
    }

    func unmarkedSum(_ card: Array<Array<(Int, Bool)>>) -> Int {
        var sum = 0
        for row in card {
            for (n, m) in row {
                if !m {
                    sum += n
                }
            }
        }
        return sum
    }

    class NoWinnerFound: Error {
    }
}
