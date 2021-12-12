import Foundation

class Day12: Day {
    private let connections: Array<(String, String)>

    init(_ input: String) {
        let inputConnections = input.split(separator: "\n")
                .map({ $0.split(separator: "-") })
                .map({ (String($0[0]), String($0[1])) })
        var allConnections = inputConnections
        for (from, to) in inputConnections {
            allConnections.append((to, from))
        }
        connections = allConnections
    }

    func part1() -> Int {
        let pathsToEnd = findPathsToEnd(start: "start", pathSoFar: ["start"])
        return pathsToEnd.count
    }

    func part2() -> Int {
        let visitCounts = ["start": 1]
        let pathsToEnd = findPathsToEnd2(start: "start", pathSoFar: ["start"], visitCounts: visitCounts)
        var asdf = Array<Array<String>>()
        for path in pathsToEnd {
            var bla: [String: Int] = [:]
            for n in path {
                if isBig(n) {
                    continue
                }
                if let x = bla[n] {
                    bla[n] = x + 1
                } else {
                    bla[n] = 1
                }
            }
            var z = 0
            for (_, v) in bla {
                if v >= 2 {
                    z += 1
                }
            }
            if z <= 1 {
                asdf.append(path)
            }
        }
        return asdf.count
    }

    private func findPathsToEnd(start: String, pathSoFar: Array<String>) -> Array<Array<String>> {
        if start == "end" {
            return [pathSoFar]
        }

        var out = Array<Array<String>>()
        for (from, to) in connections {
            if from != start || (pathSoFar.contains(to) && !isBig(to)) {
                continue
            }

            var nextPath = pathSoFar
            nextPath.append(to)

            let further = findPathsToEnd(start: to, pathSoFar: nextPath)
            if !further.isEmpty {
                out.appendAll(further)
            }
        }
        return out
    }

    private func findPathsToEnd2(start: String, pathSoFar: Array<String>, visitCounts: [String: Int]) -> Array<Array<String>> {
        if start == "end" {
            return [pathSoFar]
        }

        var out = Array<Array<String>>()
        outer: for (from, to) in connections {
            if to == "start" || from != start {
                continue
            }
            if !isBig(to) {
                if let visitCount = visitCounts[to] {
                    if visitCount >= 2 {
                        continue
                    }
                }
            }

            var nextPath = pathSoFar
            nextPath.append(to)
            var visitCounts2 = visitCounts
            if visitCounts2[to] == nil {
                visitCounts2[to] = 1
            } else {
                visitCounts2[to] = visitCounts2[to]! + 1
            }

            let further = findPathsToEnd2(start: to, pathSoFar: nextPath, visitCounts: visitCounts2)
            if !further.isEmpty {
                out.appendAll(further)
            }
        }
        return out
    }

    private func isBig(_ s: String) -> Bool {
        s.uppercased() == s
    }
}

extension Array {
    mutating func appendAll(_ other: Array) {
        for e in other {
            self.append(e)
        }
    }
}