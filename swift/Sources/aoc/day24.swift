import Foundation

class Day24: Day {
    private let inputLines: Array<String>

    init(_ input: String) {
        inputLines = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map(String.init)
    }

    func part1() throws -> Int {
        let n = 13579246899999
        let ret1 = try runCode(inputNumber: n)
        let ret2 = try runSimplified(inputNumber: n)
        print("Running code: \(ret1); running simplified form: \(ret2); same? \(ret1 == ret2)")
        return ret1
    }

    func part2() -> Int {
        showDifferences()
        return 1337
    }

    private func showDifferences() {
        var v1s = Array<String>(), v2s = Array<String>(), v3s = Array<String>()

        for lineNum in 0...17 {
            var allSame = true, previous = ""
            for digitNum in 0...13 {
                let inputLineNum = digitNum * 18 + lineNum
                let line = inputLines[inputLineNum]
                if previous != "" && line != previous {
                    allSame = false
                }
                previous = line

                print("| \(line.padding(toLength: 9, withPad: " ", startingAt: 0)) ", terminator: "")

                if lineNum == 4 {
                    v1s.append(String(line.split(separator: " ")[2]))
                } else if lineNum == 5 {
                    v2s.append(String(line.split(separator: " ")[2]))
                } else if lineNum == 15 {
                    v3s.append(String(line.split(separator: " ")[2]))
                }
            }
            print("| \(allSame ? "✅️" : "❌ ") |")
        }
        print("v1s: \(v1s) (unique: \(Set(v1s)); \(Set(v1s).count))")
        print("v2s: \(v2s) (unique: \(Set(v2s)); \(Set(v2s).count))")
        print("v3s: \(v3s) (unique: \(Set(v3s)); \(Set(v3s).count))")
    }

    internal func runCode(inputNumber: Int) throws -> Int {
        try runCode(inputs: String(inputNumber).map({ Int(String($0))! }))
    }

    internal func runCode(inputs: Array<Int>) throws -> Int {
        var inputStack = Array<Int>(inputs.reversed())
        var vars = ["w": 0, "x": 0, "y": 0, "z": 0]
        for line in inputLines {
            let parts = line.split(separator: " ").map(String.init)
            let instruction = parts[0], targetVar = parts[1]

            if instruction == "inp" {
                vars[targetVar] = inputStack.removeLast()
                continue
            }

            let rightOperand = parts[2]
            let v = vars[rightOperand] != nil ? vars[rightOperand]! : Int(rightOperand)!
            let tv = vars[targetVar]!

            if instruction == "add" {
                vars[targetVar] = tv + v
            } else if instruction == "mul" {
                vars[targetVar] = tv * v
            } else if instruction == "div" {
                vars[targetVar] = tv / v
            } else if instruction == "mod" {
                vars[targetVar] = tv % v
            } else if instruction == "eql" {
                vars[targetVar] = tv == v ? 1 : 0
            } else {
                throw InvalidInstruction(instruction: instruction)
            }
        }
        return vars["z"]!
    }

    internal func runSimplified(inputNumber: Int) throws -> Int {
        try runSimplified(inputs: String(inputNumber).map({ Int(String($0))! }))
    }

    internal func runSimplified(inputs: Array<Int>) throws -> Int {
        var zz = Array<Int>()

        for (i, w) in inputs.enumerated() {
            func getV(_ lineNumber: Int) -> Int {
                Int(inputLines[i * 18 + lineNumber].split(separator: " ")[2])!
            }

            // Assumption: this is probably only valid for my own input
            // v1 = 1 or 26; v2 = correlates to v1, either positive (v1=1) or negative (v1=26); v3 = a positive value
            let v1 = getV(4), v2 = getV(5), v3 = getV(15)

            // Imagine base 10
            // * 10 = "add position", e.g. 133 -> 1330
            // +  7 = "fill in last position", e.g. 1330 -> 1337
            // % 10 = "get last position", e.g. 1337 -> 7
            // / 10 = "remove last position", e.g. 1337 -> 133
            // z is base 26; z is a stack, * followed by + is push, % followed by / is pop

            #if false
            // Most simplified form:
            let x = ((z % 26) + v2) == w ? 0 : 1
            z = ((z/v1) * ((25 * x) + 1)) + ((w + v3) * x)
            #endif

            let rem = zz.last ?? 0
            if rem + v2 == w {
                if v1 == 26 {
                    zz.removeLast()
                }
            } else {
                let u = w + v3
                if v1 == 26 {
                    zz[zz.count-1] = u
                } else {
                    zz.append(u)
                }
            }
        }

        var z = 0
        for v in zz {
            z *= 26
            z += v
        }
        return z
    }

    private struct InvalidInstruction: Error {
        let instruction: String
    }
}
