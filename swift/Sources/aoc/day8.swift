import Foundation

class Day8: Day {
    private let notes: Array<(Array<String>, Array<String>)>

    init(_ input: String) {
        notes = input.split(separator: "\n").map { Day8.parseLine($0) }
    }

    func part1() -> Int {
        notes.map { $0.1 }
                .joined()
                .filter { value in
                    value.count == 2 /*1*/ || value.count == 4 /*4*/ || value.count == 3 /*7*/ || value.count == 7 /*8*/
                }
                .count
    }

    func part2() -> Int {
        notes.map { Day8.decode($0.0, $0.1) }.reduce(0, +)
    }

    internal static func parseLine<S: StringProtocol>(_ line: S) -> (Array<String>, Array<String>) {
        let parts = line.components(separatedBy: " | ")
        return (Day8.splitAndSort(parts[0]), Day8.splitAndSort(parts[1]))
    }

    private static func splitAndSort<S: StringProtocol>(_ s: S) -> Array<String> {
        s.split(separator: " ").map { String($0.sorted()) }
    }

    internal static func decode(_ signalPatterns: Array<String>, _ outputValues: Array<String>) -> Int {
        var numberToValue: [Int: String] = [:]

        for value in signalPatterns {
            if value.count == 2 {
                numberToValue[1] = value
            } else if value.count == 4 {
                numberToValue[4] = value
            } else if value.count == 3 {
                numberToValue[7] = value
            } else if value.count == 7 {
                numberToValue[8] = value
            }
        }

        let one = numberToValue[1]!

        func containsOne(_ value: String) -> Bool {
            value.contains(one[0]) && value.contains(one[1])
        }

        for value in signalPatterns {
            if value.count == 5 && containsOne(value) {
                numberToValue[3] = value
                break
            }
        }

        let three = numberToValue[3]!

        for value in signalPatterns {
            if value.count == 6 && three.allSatisfy({ value.contains($0) }) {
                numberToValue[9] = value
                break
            }
        }

        let nine = numberToValue[9]!

        for value in signalPatterns {
            if value.count == 6 && value != nine {
                if containsOne(value) {
                    numberToValue[0] = value
                } else {
                    numberToValue[6] = value
                }
            }
        }

        let six = numberToValue[6]!

        for value in signalPatterns {
            if value.count == 5 && value != three {
                let sixMatches = value.filter({ six.contains($0) }).count
                if sixMatches == 4 {
                    numberToValue[2] = value
                } else if sixMatches == 5 {
                    numberToValue[5] = value
                }
            }
        }

        var valueToNumber: [String: Int] = [:]
        for i in 0...9 {
            let v = numberToValue[i]!
            valueToNumber[v] = i
        }

        return Int(outputValues.map({ String(valueToNumber[$0]!) }).reduce("", +))!
    }
}

extension String {
    func at(_ i: Int) -> Character {
        self[self.index(startIndex, offsetBy: i)]
    }

    subscript(_ i: Int) -> Character {
        at(i)
    }
}