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
        var maxestY = 0
        var totalHits = 0

        for startVx in 1...targetArea.0.1 {
            // TODO: for my answer, the max vy was 147, but how to determine it up front?
            for startVy in (targetArea.1.0 - 1)...200 {
                var vx = startVx, vy = startVy
                var x = 0, y = 0, maxY = 0
                var points = Array<(Int, Int)>()

                while true {
                    x += vx
                    y += vy
                    maxY = max(maxY, y)
                    points.append((x, y))

                    if isInTargetArea(x, y) {
                        totalHits += 1
                        if maxY > maxestY {
                            maxestY = maxY
                            finalPoints = points
                        }
                        break
                    }

                    vx = max(0, vx - 1)
                    vy -= 1

                    if x > targetArea.0.1 || (vx == 0 && x < targetArea.0.0) || y < targetArea.1.0 {
                        break
                    }
                }
            }
        }

        #if false
        printGrid(points: finalPoints)
        #endif
        return (maxestY, totalHits)
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
