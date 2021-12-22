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
        // TODO: fix
        let sp1 = startingPositions[1]!, sp2 = startingPositions[2]!
        var wins1 = 0
        var wins2 = 0
        for turns in 2...21 {
            wins1 += waysToWin(turns: turns, start: sp1) * waysToLose(turns: turns - 1, start: sp2)
            wins2 += waysToWin(turns: turns, start: sp2) * waysToLose(turns: turns, start: sp1)
        }
        print("wins1", wins1)
        print("wins2", wins2)
        return max(wins1, wins2)
    }

    private let rolls = [6: 7, 5: 6, 7: 6, 4: 3, 8: 3, 3: 1, 9: 1]

    private func backward(position: Int, roll: Int) -> Int {
        ((position - roll + 9) % 10) + 1
    }

    private func waysToLose(turns: Int, start: Int) -> Int {
        var n = 0
        for score in 0...20 {
            n += allWaysToReach(score: score, turns: turns, start: start)
        }
        return n
    }

    private func waysToWin(turns: Int, start: Int) -> Int {
        var ways = 0
        for finalScore in 21...30 {
            for prevScore in finalScore - 10...min(finalScore, 21)-1 {
                let finalPos = finalScore - prevScore
                for (roll, prob) in rolls {
                    let pos = backward(position: finalPos, roll: roll)
                    ways += waysToReach(score: prevScore, position: pos, turns: turns - 1, start: start) * prob
                }
            }
        }
        return ways
    }

    private func waysToReach(score: Int, position: Int, turns: Int, start: Int) -> Int {
        if score == 0 && turns == 0 && position == start {
            return 1
        }
        if score < 1 || turns == 0 {
            return 0
        }
        var n = 0
        for (roll, prob) in rolls {
            let pos = backward(position: position, roll: roll)
            n += waysToReach(score: score - position, position: pos, turns: turns - 1, start: start) * prob
        }
        return n
    }

    private func allWaysToReach(score: Int, turns: Int, start: Int) -> Int {
        var n = 0
        for position in 1...10 {
            n += waysToReach(score: score, position: position, turns: turns, start: start)
        }
        return n
    }
}
