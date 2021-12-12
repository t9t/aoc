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
        findPathsToEnd(start: "start", pathSoFar: ["start"], allowSmallTwice: false, hasSmallTwice: false)
    }

    func part2() -> Int {
        findPathsToEnd(start: "start", pathSoFar: ["start"], allowSmallTwice: true, hasSmallTwice: false)
    }

    private func findPathsToEnd(start: String, pathSoFar: Array<String>, allowSmallTwice: Bool, hasSmallTwice: Bool) -> Int {
        if start == "end" {
            return 1
        }

        var foundPaths = 0
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

            foundPaths += findPathsToEnd(start: to, pathSoFar: nextPath, allowSmallTwice: allowSmallTwice, hasSmallTwice: hasSmallTwice)
        }
        return foundPaths
    }

    private func isSmall(_ s: String) -> Bool {
        s.lowercased() == s
    }
}