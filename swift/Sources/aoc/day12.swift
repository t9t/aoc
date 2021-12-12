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
        foundPaths = 0
        findPathsToEnd(start: "start", pathSoFar: ["start"], allowSmallTwice: false, hasSmallTwice: false)
        return foundPaths
    }

    private var foundPaths = 0

    func part2() -> Int {
        foundPaths = 0
        findPathsToEnd(start: "start", pathSoFar: ["start"], allowSmallTwice: true, hasSmallTwice: false)
        return foundPaths
    }

    private func findPathsToEnd(start: String, pathSoFar: Array<String>, allowSmallTwice: Bool, hasSmallTwice: Bool) {
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
                if allowSmallTwice {
                    if hasSmallTwice {
                        continue
                    } else {
                        hasSmallTwice = true
                    }
                } else {
                    continue
                }
            }

            var nextPath = pathSoFar
            nextPath.append(to)

            findPathsToEnd(start: to, pathSoFar: nextPath, allowSmallTwice: allowSmallTwice, hasSmallTwice: hasSmallTwice)
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