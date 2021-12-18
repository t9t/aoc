import Foundation

class Day18: Day {
    private let input: Array<Array<Token>>

    init(_ input: String) throws {
        self.input = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map(String.init).map(Day18.tokenize)
    }

    func part1() -> Int {
        Day18.determineMagnitude(Day18.sum(input))
    }

    func part2() -> Int {
        var largestMagnitude = 0
        for (i, line1) in input.enumerated() {
            for (j, line2) in input.enumerated() {
                if i == j {
                    continue
                }
                largestMagnitude = max(largestMagnitude, Day18.determineMagnitude(Day18.addAndReduce(line1, line2)))
                largestMagnitude = max(largestMagnitude, Day18.determineMagnitude(Day18.addAndReduce(line2, line1)))
            }
        }
        return largestMagnitude
    }

    internal static func determineMagnitude(_ s: Array<Token>) -> Int {
        if s.count == 1 {
            if case .Number(let n) = s.first! {
                return n
            } else {
                fatalError()
            }
        }

        let within = s[s.index(after: s.startIndex)...s.index(s.endIndex, offsetBy: -2)]
        var bc = 0
        for (i, t) in within.enumerated() {
            switch t {
            case let .Char(c):
                if c == "," && bc == 0 {
                    let left = within[...within.index(within.startIndex, offsetBy: i - 1)]
                    let right = within[within.index(within.startIndex, offsetBy: i + 1)...]

                    let leftMagnitude = determineMagnitude(Array<Token>(left))
                    let rightMagnitude = determineMagnitude(Array<Token>(right))
                    return 3 * leftMagnitude + 2 * rightMagnitude
                } else if c == "[" {
                    bc += 1
                } else if c == "]" {
                    bc -= 1
                }
            case .Number(_):
                break
            }
        }
        fatalError()
    }

    static func explodeOnceIfNecessary(_ num: Array<Token>) -> (Array<Token>, Bool) {
        let s = num
        var depth = 0
        var lastRegNumPos: Int? = nil
        var lastRegNum: Int? = nil
        for (i, t) in s.enumerated() {
            switch t {
            case let .Char(c):
                if c == "[" {
                    depth += 1
                } else if c == "]" {
                    depth -= 1
                } else {
                    continue
                }
            case let .Number(number):
                if depth >= 5 {
                    let leftNum = number
                    let prefix: Array<Token>
                    if lastRegNum != nil {
                        let leftSum = lastRegNum! + leftNum
                        let beforeLastRegNum = Array<Token>(s[s.startIndex...s.index(s.startIndex, offsetBy: lastRegNumPos! - 1)])
                        let afterLastRegNum = Array<Token>(s[s.index(s.startIndex, offsetBy: lastRegNumPos! + 1)...s.index(s.startIndex, offsetBy: i - 2)])
                        prefix = beforeLastRegNum + [Token.Number(n: leftSum)] + afterLastRegNum
                    } else {
                        prefix = Array<Token>(s[...s.index(s.startIndex, offsetBy: i - 2)])
                    }
                    let element: Token = s[s.index(s.startIndex, offsetBy: i + 2)]
                    if case .Number(let rightNum) = element {
                        let rest = s[s.index(s.startIndex, offsetBy: i + 4)...]

                        var nextRegNumPos: Int? = nil
                        var nextRegNum: Int? = nil
                        forRest: for (j, c2) in rest.enumerated() {
                            switch c2 {
                            case let .Number(n):
                                nextRegNum = n
                                nextRegNumPos = j
                                break forRest
                            default:
                                continue
                            }
                        }
                        let suffix: Array<Token>
                        if nextRegNum == nil {
                            suffix = Array<Token>(rest)
                        } else {
                            let rightSum = nextRegNum! + rightNum
                            let beforeNextRegNum = rest[rest.startIndex...rest.index(rest.startIndex, offsetBy: nextRegNumPos! - 1)]
                            let afterNextRegNum = rest[rest.index(rest.startIndex, offsetBy: nextRegNumPos! + 1)...]
                            suffix = beforeNextRegNum + [Token.Number(n: rightSum)] + afterNextRegNum
                        }
                        let newNumStr = prefix + [Token.Number(n: 0)] + suffix
                        return (newNumStr, true)
                    } else {
                        fatalError()
                    }
                } else {
                    lastRegNumPos = i
                    lastRegNum = number
                }
            }
        }
        return (num, false)
    }

    static func splitOnceIfNecessary(_ num: Array<Token>) -> (Array<Token>, Bool) {
        let s = num
        for i in 1...s.count - 1 {
            let c1 = s[s.index(s.startIndex, offsetBy: i - 1)]
            switch c1 {
            case let .Number(number):
                if number < 10 {
                    continue
                }
                let prefix = s[s.startIndex...s.index(s.startIndex, offsetBy: i - 2)]
                let suffix = s[s.index(s.startIndex, offsetBy: i)...]
                let left = number / 2
                let right = (number / 2) + (number % 2)
                let newNumber = [Token.Char(c: "["), Token.Number(n: left), Token.Char(c: ","), Token.Number(n: right), Token.Char(c: "]")]
                return (prefix + newNumber + suffix, true)
            default:
                continue
            }
        }
        return (num, false)
    }

    internal static func reduce(_ s: Array<Token>) -> Array<Token> {
        var out = s
        while true {
            let (exploded, didExplode) = explodeOnceIfNecessary(out)
            if didExplode {
                out = exploded
                continue
            }

            let (split, didSplit) = splitOnceIfNecessary(out)
            if didSplit {
                out = split
            } else {
                break
            }
        }
        return out
    }

    internal static func add(_ left: Array<Token>, _ right: Array<Token>) -> Array<Token> {
        [Token.Char(c: "[")] + left + [Token.Char(c: ",")] + right + [Token.Char(c: "]")]
    }

    internal static func addAndReduce(_ left: Array<Token>, _ right: Array<Token>) -> Array<Token> {
        reduce(add(left, right))
    }

    internal static func sum(_ numbers: Array<Array<Token>>) -> Array<Token> {
        var sum = numbers[0]
        for i in 1...numbers.count - 1 {
            sum = addAndReduce(sum, numbers[i])
        }
        return sum
    }

    internal static func tokenize(_ s: String) -> Array<Token> {
        var out = Array<Token>()
        var buf = ""
        for c in s {
            if c == "[" {
                out.append(Token.Char(c: c))
            } else if c == "]" {
                if buf != "" {
                    out.append(Token.Number(n: Int(buf)!))
                }
                buf = ""
                out.append(Token.Char(c: c))
            } else if c == "," {
                if buf != "" {
                    out.append(Token.Number(n: Int(buf)!))
                }
                buf = ""
                out.append(Token.Char(c: c))
            } else {
                buf.append(c)
            }
        }
        return out
    }

    internal enum Token {
        case Char(c: Character)
        case Number(n: Int)
    }
}
