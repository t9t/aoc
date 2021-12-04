import Foundation

class Day4 {
    typealias Card = Array<Array<(Int, Bool)>>

    let numbers: Array<Int>
    let inputCards: Array<Card>

    init(_ input: String) {
        let lines = input.split(separator: "\n", omittingEmptySubsequences: false)
        numbers = lines[0].split(separator: ",").map {
            Int($0)!
        }
        var cards = Array<Card>()
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
        inputCards = cards
    }

    func part1() throws -> Int {
        var cards = inputCards
        for num in numbers {
            cards = mark(cards, num)
            if let card = cards.first(where: { isWinning($0) }) {
                return num * sumUnmarked(card)
            }
        }
        throw NoWinnerFound()
    }

    func part2() throws -> Int {
        var cards = inputCards
        for num in numbers {
            cards = mark(cards, num)
            let notWinners = cards.filter {
                !isWinning($0)
            }
            if notWinners.isEmpty {
                return num * sumUnmarked(cards[0])
            }
            cards = notWinners
        }
        throw NoWinnerFound()
    }

    func sumUnmarked(_ card: Card) -> Int {
        card.reduce(0, { sum, row in sum + (row.filter({ !$0.1 }).reduce(0, { $0 + $1.0 })) })
    }

    func mark(_ inCards: Array<Card>, _ num: Int) -> Array<Card> {
        var outCards = inCards
        for (c, card) in outCards.enumerated() {
            for (r, row) in card.enumerated() {
                for (i, (n, _)) in row.enumerated() {
                    if (n == num) {
                        outCards[c][r][i] = (n, true)
                    }
                }
            }
        }
        return outCards
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

    class NoWinnerFound: Error {}
}