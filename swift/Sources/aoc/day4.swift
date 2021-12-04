import Foundation

class Day4 {
    typealias Card = Array<Array<(Int, Bool)>>

    let numbers: Array<Int>
    var cards: Array<Card>

    init(_ input: String) {
        let lines = input.split(separator: "\n", omittingEmptySubsequences: false)
        numbers = lines[0].split(separator: ",").map {
            Int($0)!
        }
        cards = Array<Card>()
        var card = Card()
        for line in lines.dropFirst(2) {
            if line.isEmpty {
                cards.append(card)
                card = Card()
                continue
            }

            card.append(line.split(separator: " ").map {
                (Int($0.trimmingCharacters(in: .whitespaces))!, false)
            })
        }
        if !card.isEmpty {
            cards.append(card)
        }
    }

    func part1() throws -> Int {
        for num in numbers {
            mark(num)
            for card in cards.filter { isWinning($0) } {
                return num * (card.reduce(0, { sum, row in sum + (row.filter({ !$0.1 }).reduce(0, { $0 + $1.0 })) }))
            }
        }
        throw NoWinnerFound()
    }

    func part2() -> Int {
        return 1337
    }

    func mark(_ num: Int) {
        for (c, card) in cards.enumerated() {
            for (r, row) in card.enumerated() {
                for (i, (n, _)) in row.enumerated() {
                    if (n == num) {
                        cards[c][r][i] = (n, true)
                    }
                }
            }
        }
    }

    func isWinning(_ card: Array<Array<(Int, Bool)>>) -> Bool {
        if let _ = card.firstIndex(where: { row in row.map { $0.1 }.reduce(true, { $0 && $1 }) }) {
            return true
        }

        if let _ = (0...card[0].count - 1).firstIndex(where: { i in card.map { $0[i].1 }.reduce(true, { $0 && $1 }) }) {
            return true
        }

        return false
    }

    class NoWinnerFound: Error {
    }
}