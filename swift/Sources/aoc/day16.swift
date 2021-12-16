import Foundation

class Day16: Day {
    private let input: String

    init(_ input: String) {
        self.input = input.trimmingCharacters(in: .whitespacesAndNewlines)
    }

    func part1() -> Int {
        readAllVersionNumbersSum(Day16.allHexToBin(input))
    }


    func part2() -> Int {
        return 1337
    }

    private func readAllVersionNumbersSum(_ bin: String) -> Int {
        let buf = Buffer(bin)
        var versionNumbersSum = 0
        while buf.hasMore() {
            versionNumbersSum += readVersionOfOnePacket(buf)
        }
        return versionNumbersSum
    }

    private func readVersionOfOnePacket(_ buf: Buffer) -> Int {
        var versionNumbersSum = binToInt(buf.advance(by: 3))
        let typeId = binToInt(buf.advance(by: 3))

        if typeId == 4 /* literal value */ {
            var prefix = "1"
            while prefix != "0" {
                prefix = buf.advance(by: 1)
                let _ = buf.advance(by: 4) // Number bits
            }
        } else {
            let lengthTypeId = buf.advance(by: 1)
            if lengthTypeId == "0" {
                let totalLengthInBits = binToInt(buf.advance(by: 15))
                let subPackets = buf.advance(by: totalLengthInBits)
                versionNumbersSum += readAllVersionNumbersSum(subPackets)
            } else {
                let numberOfSubPackets = binToInt(buf.advance(by: 11))

                for _ in 1...numberOfSubPackets {
                    versionNumbersSum += readVersionOfOnePacket(buf)
                }
            }
        }
        return versionNumbersSum
    }

    private func binToInt(_ bin: String) -> Int {
        Int(bin, radix: 2)!
    }

    private static func hexCharToBin(_ char: Character) -> String {
        let s = String(Int(String(char), radix: 16)!, radix: 2)
        if s.count < 4 {
            return String(repeating: "0", count: 4 - s.count) + s
        }
        return s
    }

    internal static func allHexToBin(_ hex: String) -> String {
        hex.map(Day16.hexCharToBin).joined()
    }

    private class Buffer {
        private var remaining: String

        init(_ remaining: String) {
            self.remaining = remaining
        }

        func advance(by: Int) -> String {
            let value = String(remaining[...remaining.index(remaining.startIndex, offsetBy: by - 1)])
            remaining = String(remaining[remaining.index(remaining.startIndex, offsetBy: by)...])
            return value
        }

        func hasMore() -> Bool {
            for c in remaining {
                if c != "0" {
                    return true
                }
            }
            return false
        }
    }
}
