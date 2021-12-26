import Foundation

class Day24: Day {
    private let inputLines: Array<String>

    init(_ input: String) {
        inputLines = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map(String.init)
    }

    func part1() throws -> Int {
        try calculateHighestAndLowestModelNumber().0
    }

    func part2() throws -> Int {
        try calculateHighestAndLowestModelNumber().1
    }

    internal func calculateHighestAndLowestModelNumber() throws -> (Int, Int) {
        var stack = Array<(Int, Int)>()
        var highestDigits = Array<Int>(repeating: 0, count: 14)
        var lowestDigits = Array<Int>(repeating: 0, count: 14)

        for i in 0...13 {
            func getV(_ lineNumber: Int) -> Int {
                Int(inputLines[i * 18 + lineNumber].split(separator: " ")[2])!
            }

            // Assumption: line numbers will be the same in each input, with only the values on those lines being different
            let v1 = getV(4), v2 = getV(5), v3 = getV(15)
            if v1 == 1 {
                // push
                stack.append((i, v3))
            } else {
                let (otherI, otherV) = stack.removeLast()
                let diff = v2 + otherV
                if diff >= 0 {
                    highestDigits[i] = 9
                    highestDigits[otherI] = 9 - diff

                    lowestDigits[i] = 1 + diff
                    lowestDigits[otherI] = 1
                } else {
                    highestDigits[i] = 9 + diff
                    highestDigits[otherI] = 9

                    lowestDigits[i] = 1
                    lowestDigits[otherI] = 1 - diff
                }
            }
        }

        func toInt(_ arr: Array<Int>) throws -> Int {
            if arr.contains(0) {
                throw InvalidInput()
            }
            return Int(arr.map({ String($0) }).joined())!
        }

        return (try toInt(highestDigits), try toInt(lowestDigits))
    }

    private struct InvalidInput: Error {
    }

    // All code below is a remainder of the manual reverse engineering process, left in for posterity

    /// Check a number by running the input code & the simplified form, validating the algorithms
    private func checkNumber(_ n: Int) throws {
        let ret1 = try runCode(inputNumber: n)
        let ret2 = try runSimplified(inputNumber: n)
        print("Running code: \(ret1); running simplified form: \(ret2); same? \(ret1 == ret2)")
    }

    /// Visualize differences in the input code for processing each number
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

                // Assumption: line numbers will be the same in each input, with only the values on those lines being different
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
        var inputStack = Array<Int>(String(inputNumber).map({ Int(String($0))! }).reversed())
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

    private func runSimplified(inputNumber: Int) throws -> Int {
        let inputs = String(inputNumber).map({ Int(String($0))! })
        // Stack for processing (see comments below why)
        var zStack = Array<Int>()
        var z = 0

        for (i, w) in inputs.enumerated() {
            func getV(_ lineNumber: Int) -> Int {
                Int(inputLines[i * 18 + lineNumber].split(separator: " ")[2])!
            }

            // Assumption: this is probably only valid for my own input
            // v1 = 1 or 26; v2 = correlates to v1, either >0 (v1=1) or <=0 (v1=26); v3 = a positive value
            // When v1 = 1, v2 is always >=10 so can never be equal to the input
            let v1 = getV(4), v2 = getV(5), v3 = getV(15)

            // Imagine base 10
            // * 10 = "add position", e.g. 133 -> 1330
            // +  7 = "fill in last position", e.g. 1330 -> 1337
            // % 10 = "get last position", e.g. 1337 -> 7
            // / 10 = "remove last position", e.g. 1337 -> 133
            // z is base 26; z is a stack; * followed by + is push; % followed by / is pop
            // z = 0 means an empty stack, so the goal is to empty the stack
            // Since there are 7 pushes (v1=1), we need to make sure each other instruction (v1=26) is able to pop

            // Most simplified form of my input:
            let x = ((z % 26) + v2) == w ? 0 : 1
            z = ((z / v1) * ((25 * x) + 1)) + ((w + v3) * x)

            // Detailed overview of differences for my input & manually constructing any number based on push/pop rules:
            //|  i | 0         | 1         | 2         | 3         | 4         | 5         | 6         | 7         | 8         | 9         | 10        | 11        | 12        | 13        |
            //| v1 | div z 1   | div z 1   | div z 1   | div z 26  | div z 1   | div z 26  | div z 26  | div z 1   | div z 1   | div z 1   | div z 26  | div z 26  | div z 26  | div z 26  |
            //| v2 | add x 13  | add x 11  | add x 15  | add x -6  | add x 15  | add x -8  | add x -4  | add x 15  | add x 10  | add x 11  | add x -11 | add x 0   | add x -8  | add x -7  |
            //| v3 | add y 3   | add y 12  | add y 9   | add y 12  | add y 2   | add y 1   | add y 1   | add y 13  | add y 1   | add y 6   | add y 2   | add y 11  | add y 10  | add y 3   |
            //     | push      | push      | push      | ?         | push      | ?         | ?         | push      | push      | push      | ?         | ?         | ?         | ?         |
            //       1           3           5           7           9           2           4           6           8           9           9           9           9           9
            //     | 5: +8     | 1: +13    | 5: +14    | 14-6=8: - | 9: +11    | 11-8=3: - | 13-4=9: - | 4: +17    | 8: +9     | 9: +15    | 15-11=4:- | 9-0=9: -  | 17-8=9: - | 8-7=1: -  |
            //     | push -> 1 | push -> 2 | push -> 3 | pop[2]->2 | push -> 3 | pop[4]->2 | pop[1]->1 | push -> 2 | push -> 3 | push -> 4 | pop[9]->3 | pop[8]->2 | pop[7]->1 | pop[0]->0 |

            // pushed at -> popped at; what is the highest value we can push, to be able to pop?
            // 0 -> 13; push  +3; pop  -7; 9 -> 5 (diff -4)
            // 1 ->  6; push +12; pop  -4; 1 -> 9 (diff +8)
            // 2 ->  3; push  +9; pop  -6; 6 -> 9 (diff +3)
            // 4 ->  5; push  +2; pop  -8; 9 -> 3 (diff -6)
            // 7 -> 12; push +13; pop  -8; 4 -> 9 (diff +5)
            // 8 -> 11; push: +1; pop  -0; 8 -> 9 (diff +1)
            // 9 -> 10; push: +6; pop -11; 9 -> 4 (diff -5)
            // 91699394894995

            // pushed at -> popped at; what is the lowest value we can push, to be able to pop?
            // 0 -> 13; push  +3; pop  -7; 5 -> 1
            // 1 ->  6; push +12; pop  -4; 1 -> 9
            // 2 ->  3; push  +9; pop  -6; 1 -> 4
            // 4 ->  5; push  +2; pop  -8; 7 -> 1
            // 7 -> 12; push +13; pop  -8; 1 -> 6
            // 8 -> 11; push: +1; pop  -0; 1 -> 2
            // 9 -> 10; push: +6; pop -11; 6 -> 1
            // 51147191161261

            let u = w + v3
            if v1 == 1 {
                zStack.append(u)
            } else if zStack.last! + v2 == w { // v1 == 26 and v2 is always <=0 here
                zStack.removeLast()
            } else { // v1 == 26; this will never happen for a valid input (only pushes and pops)
                zStack[zStack.count - 1] = u
            }

            print("Step \(i); w: \(w); v1: \(v1); v2: \(v2); v3: \(v3); zStack: \(zStack)")
        }

        // Reconstruct "z" value from the stack
        var reZ = 0
        for v in zStack {
            reZ *= 26 // shift left
            reZ += v // add
        }

        // For example, 13579246899999 outputs 1422667258
        print("z: \(z); zzz: \(reZ); same? \(z == reZ)")
        return z
    }

    private struct InvalidInstruction: Error {
        let instruction: String
    }

}
