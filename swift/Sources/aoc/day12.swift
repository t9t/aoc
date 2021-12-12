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
        return 1337
    }

    private func findPathsToEnd(start: String, pathSoFar: Array<String>) -> Array<Array<String>> {
        if start == "end" {
            var finalPath = pathSoFar
            finalPath.append("end")
            return [finalPath]
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