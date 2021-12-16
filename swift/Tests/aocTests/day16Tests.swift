import XCTest
import class Foundation.Bundle

@testable import aoc

final class day16Tests: XCTestCase {
    func testAllHexToBin_D2FE28() throws {
        XCTAssertEqual(Day16.allHexToBin("D2FE28"), "110100101111111000101000")
    }

    func testAllHexToBin_38006F45291200() throws {
        XCTAssertEqual(Day16.allHexToBin("38006F45291200"), "00111000000000000110111101000101001010010001001000000000")
    }

    func testPart1() throws {
        let testCases = [
            "D2FE28": 6,
            "38006F45291200": 1 + 6 + 2,
            "EE00D40C823060": 7 + 2 + 4 + 1,
            "8A004A801A8002F478": 16,
            "620080001611562C8802118E34": 12,
            "C0015000016115A2E0802F182340": 23,
            "A0016C880162017C3686B18A3D4780": 31,
        ]
        for (input, expected) in testCases {
            let n = Day16(input).part1()
            XCTAssertEqual(n, expected)
        }
    }


    func testPart2() throws {
        let testCases = [
            "C200B40A82": 3,
            "04005AC33890": 54,
            "880086C3E88112": 7,
            "CE00C43D881120": 9,
            "D8005AC2A8F0": 1,
            "F600BC2D8F": 0,
            "9C005AC2F8F0": 0,
            "9C0141080250320F1802104A08": 1,
        ]
        for (input, expected) in testCases {
            let n = Day16(input).part2()
            XCTAssertEqual(n, expected)
        }
    }

}
