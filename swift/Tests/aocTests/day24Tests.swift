import XCTest
import class Foundation.Bundle

@testable import aoc

final class day24Tests: XCTestCase {
    private let input = """
                        inp w
                        add z w
                        add z 10
                        add x z
                        mul x 3
                        add y x
                        div y 2
                        inp w
                        add w y
                        add w 1
                        mod x 9
                        div w 4
                        eql x w
                        mul z 0
                        add z x
                        """

    func testRunCode() throws {
        let n = try Day24(input).runCode(inputNumber: 42)
        XCTAssertEqual(n, 1)
    }

    func testNegativeNumbers() throws {
        let inputToZ = "inp w\nadd z w\n"
        XCTAssertEqual(try Day24(inputToZ + "add z -5").runCode(inputNumber: 3), -2)
        XCTAssertEqual(try Day24(inputToZ + "mul z -5").runCode(inputNumber: 3), -15)
        XCTAssertEqual(try Day24(inputToZ + "div z -2").runCode(inputNumber: 6), -3)
        XCTAssertEqual(try Day24(inputToZ + "mod z -5").runCode(inputNumber: 7), 2)
        XCTAssertEqual(try Day24(inputToZ + "add z -5\neql z -4").runCode(inputNumber: 1), 1)
    }
}
