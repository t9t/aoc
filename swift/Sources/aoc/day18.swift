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
            return RegularNumber(value: Int(s)!)
        }

        let within = s[s.index(after: s.startIndex)...s.index(s.endIndex, offsetBy: -2)]
        var bc = 0
        for (i, c) in within.enumerated() {
            if c == "," && bc == 0 {
                let left = within[...within.index(within.startIndex, offsetBy: i - 1)]
                let right = within[within.index(within.startIndex, offsetBy: i + 1)...]

                return PairNumber(left: try parseNumber(left), right: try parseNumber(right))
            } else if c == "[" {
                bc += 1
            } else if c == "]" {
                bc -= 1
            }
        }

        throw InvalidNumberString(number: String(s))
    }

    static func explodeOnceIfNecessaryXX(_ num: PairNumber) -> PairNumber {
        let s = "\(num)"
        var depth = 0
        var lastRegNumPos: Int? = nil
        var lastRegNum: Int? = nil
        for (i, c) in s.enumerated() {
            if c == "[" {
                depth += 1
            } else if c == "]" {
                depth -= 1
            } else if c != "," {
                if depth >= 5 {
                    let leftNum = Int(String(c))!
                    let prefix: String
                    if lastRegNum != nil {
                        let leftSum = lastRegNum! + leftNum
                        let beforeLastRegNum = s[s.startIndex...s.index(s.startIndex, offsetBy: lastRegNumPos! - 1)]
                        let afterLastRegNum = s[s.index(s.startIndex, offsetBy: lastRegNumPos! + 1)...s.index(s.startIndex, offsetBy: i - 2)]
                        prefix = beforeLastRegNum + String(leftSum) + afterLastRegNum
                    } else {
                        prefix = String(s[...s.index(s.startIndex, offsetBy: i - 2)])
                    }
                    let rightNum = Int(String(s[s.index(s.startIndex, offsetBy: i + 2)]))!
                    let rest = s[s.index(s.startIndex, offsetBy: i + 4)...]

                    var nextRegNumPos: Int? = nil
                    var nextRegNum: Int? = nil
                    for (j, c2) in rest.enumerated() {
                        if c2 != "[" && c2 != "]" && c2 != "," {
                            nextRegNum = Int(String(c2))!
                            nextRegNumPos = j
                            break
                        }
                    }
                    let suffix: String
                    if nextRegNum == nil {
                        suffix = String(rest)
                    } else {
                        let rightSum = nextRegNum! + rightNum
                        let beforeNextRegNum = rest[rest.startIndex...rest.index(rest.startIndex, offsetBy: nextRegNumPos! - 1)]
                        let afterNextRegNum = rest[rest.index(rest.startIndex, offsetBy: nextRegNumPos! + 1)...]
                        suffix = beforeNextRegNum + String(rightSum) + afterNextRegNum

                    }
                    let newNumStr = prefix + "0" + suffix
                    do {
                        return (try parseNumber(newNumStr) as? PairNumber)!
                    } catch {
                        fatalError("\(error)")
                    }
                } else {
                    lastRegNumPos = i
                    lastRegNum = Int(String(c))!
                }
            }
        }

        return PairNumber(left: RegularNumber(value: 4), right: RegularNumber(value: 2))
    }

    static func explodeOnceIfNecessary(_ num: String) -> String {
        let s = tokenize("\(num)")
        var depth = 0
        var lastRegNumPos: Int? = nil
        var lastRegNum: Int? = nil
        for (i, c) in s.enumerated() {
            if c == "[" {
                depth += 1
            } else if c == "]" {
                depth -= 1
            } else if c != "," {
                if depth >= 5 {
                    let leftNum = Int(String(c))!
                    let prefix: Array<String>
                    if lastRegNum != nil {
                        let leftSum = lastRegNum! + leftNum
                        let beforeLastRegNum = s[s.startIndex...s.index(s.startIndex, offsetBy: lastRegNumPos! - 1)]
                        let afterLastRegNum = s[s.index(s.startIndex, offsetBy: lastRegNumPos! + 1)...s.index(s.startIndex, offsetBy: i - 2)]
                        prefix = beforeLastRegNum + [String(leftSum)] + afterLastRegNum
                    } else {
                        prefix = Array<String>(s[...s.index(s.startIndex, offsetBy: i - 2)])
                    }
                    let rightNum = Int(String(s[s.index(s.startIndex, offsetBy: i + 2)]))!
                    let rest = s[s.index(s.startIndex, offsetBy: i + 4)...]

                    var nextRegNumPos: Int? = nil
                    var nextRegNum: Int? = nil
                    for (j, c2) in rest.enumerated() {
                        if c2 != "[" && c2 != "]" && c2 != "," {
                            nextRegNum = Int(String(c2))!
                            nextRegNumPos = j
                            break
                        }
                    }
                    let suffix: Array<String>
                    if nextRegNum == nil {
                        suffix = Array<String>(rest)
                    } else {
                        let rightSum = nextRegNum! + rightNum
                        let beforeNextRegNum = rest[rest.startIndex...rest.index(rest.startIndex, offsetBy: nextRegNumPos! - 1)]
                        let afterNextRegNum = rest[rest.index(rest.startIndex, offsetBy: nextRegNumPos! + 1)...]
                        suffix = beforeNextRegNum + [String(rightSum)] + afterNextRegNum

                    }
                    let newNumStr = prefix + ["0"] + suffix
                    return newNumStr.joined()
                } else {
                    lastRegNumPos = i
                    lastRegNum = Int(String(c))!
                }
            }
        }
        return num

        //return PairNumber(left: RegularNumber(value: 4), right: RegularNumber(value: 2))
    }

    static func splitOnceIfNecessaryX(_ num: PairNumber) -> PairNumber {
        let s = "\(num)"
        for i in 1...s.count - 1 {
            let c1 = s[s.index(s.startIndex, offsetBy: i - 1)]
            let c2 = s[s.index(s.startIndex, offsetBy: i)]
            if c1 == "[" || c1 == "]" || c1 == "," || c2 == "[" || c2 == "]" || c2 == "," {
                continue
            }
            let number = Int(String(c1) + String(c2))!
            let prefix = s[s.startIndex...s.index(s.startIndex, offsetBy: i - 2)]
            let suffix = s[s.index(s.startIndex, offsetBy: i + 1)...]
            let left = number / 2
            let right = (number / 2) + (number % 2)
            do {
                let ret = prefix + "[" + String(left) + "," + String(right) + "]" + suffix
                return (try parseNumber(ret) as? PairNumber)!
            } catch {
                fatalError("\(error)")
            }
        }
        return num
    }

    static func splitOnceIfNecessary(_ num: String) -> String {
        let s = tokenize("\(num)")
        for i in 1...s.count - 1 {
            let c1 = s[s.index(s.startIndex, offsetBy: i - 1)]
            if c1 == "[" || c1 == "]" || c1 == "," {
                continue
            }
            let number = Int(String(c1))!
            if number < 9 {
                continue
            }
            let prefix = s[s.startIndex...s.index(s.startIndex, offsetBy: i - 2)]
            let suffix = s[s.index(s.startIndex, offsetBy: i)...]
            let left = number / 2
            let right = (number / 2) + (number % 2)
            let newNumber = ["[", String(left), ",", String(right), "]"]
            return (prefix + newNumber + suffix).joined()
        }
        return num
    }

    internal static func tokenize(_ s: String) -> Array<String> {
        var out = Array<String>()
        var buf = ""
        for c in s {
            if c == "[" {
                out.append(String(c))
            } else if c == "]" {
                if buf != "" {
                    out.append(buf)
                }
                buf = ""
                out.append(String(c))
            } else if c == "," {
                if buf != "" {
                    out.append(buf)
                }
                buf = ""
                out.append(String(c))
            } else {
                buf.append(c)
            }
        }
        return out
    }

    internal class Number: Equatable {
        func equalTo(rhs: Number) -> Bool {
            false
        }

        static func ==(lhs: Number, rhs: Number) -> Bool {
            lhs.equalTo(rhs: rhs)
        }

        static func +(lhs: Number, rhs: Number) -> Number {
            PairNumber(left: lhs, right: rhs)
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

        static func ==(lhs: RegularNumber, rhs: RegularNumber) -> Bool {
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

        static func ==(lhs: PairNumber, rhs: PairNumber) -> Bool {
            lhs.left.equalTo(rhs: rhs.left) && lhs.right.equalTo(rhs: rhs.right)
        }
    }

    private struct InvalidNumberString: Error {
        let number: String
    }
}
