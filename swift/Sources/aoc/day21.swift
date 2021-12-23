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
        startingPositions.keys.forEach({ scores[$0] = 0 })

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
                let totalMoves = roll1 + roll2 + roll3

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
        let initialState = GameState(player1: Player(position: startingPositions[1]!, score: 0), player2: Player(position: startingPositions[2]!, score: 0), player1Turn: true)
        var cache = [GameState: (Int, Int)]()
        let (wins1, wins2) = countWins(initialState, cache: &cache)
        return max(wins1, wins2)
    }

    private func countWins(_ state: GameState, cache: inout [GameState: (Int, Int)]) -> (Int, Int) {
        if state.player1.score >= 21 {
            return (1, 0)
        } else if state.player2.score >= 21 {
            return (0, 1)
        } else if let outcome = cache[state] {
            return outcome
        }

        var totalWins1 = 0, totalWins2 = 0
        for die1 in 1...3 {
            for die2 in 1...3 {
                for die3 in 1...3 {
                    let totalRoll = die1 + die2 + die3
                    var player1 = state.player1, player2 = state.player2
                    let currentPlayer = state.player1Turn ? player1 : player2
                    var newPosition = (currentPlayer.position + totalRoll) % 10
                    if newPosition == 0 {
                        newPosition = 10
                    }
                    let newScore = currentPlayer.score + newPosition
                    let updatedPlayer = Player(position: newPosition, score: newScore)
                    if state.player1Turn {
                        player1 = updatedPlayer
                    } else {
                        player2 = updatedPlayer
                    }

                    let newState = GameState(player1: player1, player2: player2, player1Turn: !state.player1Turn)
                    let (wins1, wins2) = countWins(newState, cache: &cache)
                    totalWins1 += wins1
                    totalWins2 += wins2
                }
            }
        }

        let ret = (totalWins1, totalWins2)
        cache[state] = ret
        return ret
    }

    private struct Player: Hashable, Equatable {
        let position: Int, score: Int
    }

    private struct GameState: Hashable, Equatable {
        let player1: Player, player2: Player, player1Turn: Bool
    }

}
