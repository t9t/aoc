import Foundation

class Day18: Day {
    private let inputNumbers: Array<Number>

    init(_ input: String) throws {
        inputNumbers = try input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map({ try Day18.parseNumber($0) })
    }

    func part1() -> Int {
        for n in inputNumbers {
            print(n)
        }
        return 42
    }

    func part2() -> Int {
        return 1337
    }

    internal static func parseNumber<S: StringProtocol>(_ s: S) throws -> Number {
        if s.count == 1 {
            return Day18.RegularNumber(value: Int(s)!)
        }

        let within = s[s.index(after: s.startIndex)...s.index(s.endIndex, offsetBy: -2)]
        var bc = 0
        for (i, c) in within.enumerated() {
            if c == "," && bc == 0 {
                let left = within[...within.index(within.startIndex, offsetBy: i - 1)]
                let right = within[within.index(within.startIndex, offsetBy: i + 1)...]

                return Day18.PairNumber(left: try Day18.parseNumber(left), right: try Day18.parseNumber(right))
            } else if c == "[" {
                bc += 1
            } else if c == "]" {
                bc -= 1
            }
        }

        throw InvalidNumberString(number: String(s))
    }

    internal class Number: Equatable {
        func equalTo(rhs: Day18.Number) -> Bool {
            false
        }

        static func ==(lhs: Day18.Number, rhs: Day18.Number) -> Bool {
            lhs.equalTo(rhs: rhs)
        }
    }

    internal class RegularNumber: Number, CustomStringConvertible {
        let value: Int

        init(value: Int) {
            self.value = value
            super.init()
        }

        var description: String {
            "\(value)"
        }

        override func equalTo(rhs: Number) -> Bool {
            if let reg = rhs as? RegularNumber {
                return reg == self
            }
            return false
        }

        static func ==(lhs: Day18.RegularNumber, rhs: Day18.RegularNumber) -> Bool {
            lhs.value == rhs.value
        }
    }

    internal class PairNumber: Number, CustomStringConvertible {
        let left: Number
        let right: Number

        init(left: Number, right: Number) {
            self.left = left
            self.right = right
            super.init()
        }

        var description: String {
            "[\(left),\(right)]"
        }

        override func equalTo(rhs: Number) -> Bool {
            if let reg = rhs as? PairNumber {
                return reg == self
            }
            return false
        }

        static func ==(lhs: Day18.PairNumber, rhs: Day18.PairNumber) -> Bool {
            lhs.left.equalTo(rhs: rhs.left) && lhs.right.equalTo(rhs: rhs.right)
        }
    }

    private struct InvalidNumberString: Error {
        let number: String
    }
}
