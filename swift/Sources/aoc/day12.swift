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

    private var foundPaths = 0
    func part2() -> Int {
        foundPaths = 0
        findPathsToEnd2(start: "start", pathSoFar: ["start"], hasSmallTwice: false)
        return foundPaths
    }


    private func findPathsToEnd(start: String, pathSoFar: Array<String>) -> Array<Array<String>> {
        if start == "end" {
            return [pathSoFar]
        }

        var out = Array<Array<String>>()
        for (from, to) in connections {
            if from != start || (pathSoFar.contains(to) && isSmall(to)) {
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

    private func findPathsToEnd2(start: String, pathSoFar: Array<String>, hasSmallTwice: Bool) {
        if start == "end" {
            foundPaths += 1
            return
        }

        for (from, to) in connections {
            if to == "start" || from != start {
                continue
            }
            var hasSmallTwice = hasSmallTwice
            if isSmall(to) && pathSoFar.contains(to) {
                if hasSmallTwice {
                    continue
                } else {
                    hasSmallTwice = true
                }
            }

            var nextPath = pathSoFar
            nextPath.append(to)

            findPathsToEnd2(start: to, pathSoFar: nextPath, hasSmallTwice: hasSmallTwice)
        }
    }

    private func isSmall(_ s: String) -> Bool {
        s.lowercased() == s
    }
}

extension Array {
    mutating func appendAll(_ other: Array) {
        for e in other {
            self.append(e)
        }
    }
}