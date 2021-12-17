import Foundation

class Day17: Day {
    private let targetArea: ((Int, Int), (Int, Int))

    init(_ input: String) {
        func parseRange(_ s: String) -> (Int, Int) {
            let parts = s.components(separatedBy: "..").map({ Int($0)! })
            return (parts[0], parts[1])
        }

        let xAndY = input.trimmingCharacters(in: .whitespacesAndNewlines).components(separatedBy: ": x=")[1].components(separatedBy: ", y=").map(parseRange)
        targetArea = (xAndY[0], xAndY[1])
    }

    func part1() -> Int {
        simulateVectors().0
    }

    func part2() -> Int {
        simulateVectors().1
    }

    private func simulateVectors() -> (Int, Int) {
        var finalPoints = Array<(Int, Int)>()
        var finalMaxHeight = 0
        var totalHits = 0

        for vx in 1...targetArea.0.1 {
            for vy in (targetArea.1.0 - 1)...1000 {
                var vector = (vx, vy)
                var position = (0, 0)
                var points = Array<(Int, Int)>()

                while true {
                    let newX = position.0 + vector.0
                    let newY = position.1 + vector.1
                    position = (newX, newY)
                    points.append(position)

                    if isInTargetArea(newX, newY) {
                        totalHits += 1
                        let maxHeight = points.map({ $0.1 }).max()!
                        if maxHeight > finalMaxHeight {
                            finalMaxHeight = maxHeight
                            finalPoints = points
                        }
                        break
                    }

                    let newDx = max(0, vector.0 - 1)
                    let newDy = vector.1 - 1

                    vector = (newDx, newDy)

                    if newX > targetArea.0.1 || (vector.0 == 0 && newX < targetArea.0.0) || newY < targetArea.1.0 {
                        break
                    }
                }
            }
        }

        #if false
        printGrid(points: finalPoints)
        #endif
        return (finalMaxHeight, totalHits)
    }

    private func isInTargetArea(_ x: Int, _ y: Int) -> Bool {
        x >= targetArea.0.0 && x <= targetArea.0.1 && y >= targetArea.1.0 && y <= targetArea.1.1
    }

    private func printGrid(points: Array<(Int, Int)>) {
        let minX = min(0, min(points.map({ $0.0 }).min()!, targetArea.0.0)), maxX = max(0, max(points.map({ $0.0 }).max()!, targetArea.0.1))
        let minY = min(0, min(points.map({ $0.1 }).min()!, targetArea.1.0)), maxY = max(0, max(points.map({ $0.1 }).max()!, targetArea.1.1))

        for y in (minY...maxY).reversed() {
            for x in minX...maxX {
                let isPoint = points.contains(where: { $0.0 == x && $0.1 == y })
                let char = x == 0 && y == 0 ? "S" : (isPoint ? "#" : (isInTargetArea(x, y) ? "T" : "."))
                print(char, terminator: "")
            }
            print("")
        }
    }
}
