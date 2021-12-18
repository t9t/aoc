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
            case .Open:
                bc += 1
            case .Close:
                bc -= 1
            case .Comma:
                if bc == 0 {
                    let left = within[...within.index(within.startIndex, offsetBy: i - 1)]
                    let right = within[within.index(within.startIndex, offsetBy: i + 1)...]

                    let leftMagnitude = determineMagnitude(Array<Token>(left))
                    let rightMagnitude = determineMagnitude(Array<Token>(right))
                    return 3 * leftMagnitude + 2 * rightMagnitude
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
            case .Open:
                depth += 1
            case .Close:
                depth -= 1
            case .Comma:
                continue
            case let .Number(number):
                if depth >= 5 {
                    let leftNum = number
                    var out = Array<Token>()
                    if lastRegNum != nil {
                        out.appendAll(s[s.startIndex...s.index(s.startIndex, offsetBy: lastRegNumPos! - 1)])
                        out.append(Token.Number(n: lastRegNum! + leftNum))
                        out.appendAll(s[s.index(s.startIndex, offsetBy: lastRegNumPos! + 1)...s.index(s.startIndex, offsetBy: i - 2)])
                    } else {
                        out.appendAll(s[...s.index(s.startIndex, offsetBy: i - 2)])
                    }

                    out.append(Token.Number(n: 0))

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

                        if nextRegNum == nil {
                            out.appendAll(rest)
                        } else {
                            out.appendAll(rest[rest.startIndex...rest.index(rest.startIndex, offsetBy: nextRegNumPos! - 1)])
                            out.append(Token.Number(n: nextRegNum! + rightNum))
                            out.appendAll(rest[rest.index(rest.startIndex, offsetBy: nextRegNumPos! + 1)...])
                        }

                        return (out, true)
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
                let newNumber = [Token.Open, Token.Number(n: left), Token.Comma, Token.Number(n: right), Token.Close]
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
        [Token.Open] + left + [Token.Comma] + right + [Token.Close]
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
                out.append(Token.Open)
            } else if c == "]" {
                if buf != "" {
                    out.append(Token.Number(n: Int(buf)!))
                }
                buf = ""
                out.append(Token.Close)
            } else if c == "," {
                if buf != "" {
                    out.append(Token.Number(n: Int(buf)!))
                }
                buf = ""
                out.append(Token.Comma)
            } else {
                buf.append(c)
            }
        }
        return out
    }

    internal enum Token {
        case Open
        case Close
        case Comma
        case Number(n: Int)
    }
}

extension Array {
    mutating func appendAll(_ other: Array<Element>) {
        for e in other {
            self.append(e)
        }
    }

    mutating func appendAll(_ other: ArraySlice<Element>) {
        for e in other {
            self.append(e)
        }
    }
}
