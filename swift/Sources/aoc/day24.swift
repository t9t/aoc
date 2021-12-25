import Foundation

class Day24: Day {
    private let inputLines: Array<String>

    init(_ input: String) {
        inputLines = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n").map(String.init)
    }

    func part1() throws -> Int {
        try runCode(inputNumber: 13579246899999)
    }

    func part2() -> Int {
        1337
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

    private struct InvalidInstruction: Error {
        let instruction: String
    }
}
