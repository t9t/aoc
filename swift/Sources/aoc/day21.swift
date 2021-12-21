import Foundation

class Day21: Day {
    private let startingPositions: [Int: Int]

    init(_ input: String) {
        var startingPositions = [Int: Int]()
        input.trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: "\n")
                .map({ $0.components(separatedBy: " starting position: ") })
                .map({ (Int($0[0].components(separatedBy: "Player ")[1])!, Int($0[1])!) })
                .forEach({ startingPositions[$0.0] = $0.1 })
        self.startingPositions = startingPositions
    }

    func part1() -> Int {
        let maxPlayerNum = startingPositions.keys.max()!
        var positions = startingPositions
        var scores = [Int: Int]()
        var nextDieValue = 1, totalDieRolls = 0
        startingPositions.keys.forEach({scores[$0] = 0})

        func rollDie() -> Int {
            let roll = nextDieValue
            totalDieRolls += 1
            nextDieValue += 1
            if nextDieValue > 100 {
                nextDieValue = 1
            }
            return roll
        }

        while true {
            for player in 1...maxPlayerNum {
                let position = positions[player]!
                let roll1 = rollDie(), roll2 = rollDie(), roll3 = rollDie()
                let totalMoves = roll1+roll2+roll3

                var newPosition = position + totalMoves
                while newPosition > 10 {
                    newPosition -= 10
                }
                positions[player] = newPosition
                let currentScore = scores[player]!
                let newScore = currentScore + newPosition
                scores[player] = newScore

                #if false
                print("Player \(player) at \(position) (score: \(currentScore)) rolls \(roll1)+\(roll2)+\(roll3)=\(totalMoves) and moves to space \(newPosition) for a total score of \(newScore).")
                #endif

                if newScore >= 1000 {
                    return totalDieRolls * scores.values.min()!
                }
            }
        }
    }

    func part2() -> Int {
        return 1337
    }
}
