import Foundation

class Day16: Day {
    private let input: String

    init(_ input: String) {
        self.input = input.trimmingCharacters(in: .whitespacesAndNewlines)
    }

    func part1() -> Int {
        readInputPacketVersionSumAndValue().0
    }

    func part2() -> Int {
        readInputPacketVersionSumAndValue().1
    }

    private func readInputPacketVersionSumAndValue() -> (Int, Int) {
        let buf = Buffer(Day16.allHexToBin(input))
        // Assumption: outer packet is always an operator packet
        return readPacketVersionSumAndValue(buf)
    }

    private func readPacketVersionSumAndValue(_ buf: Buffer) -> (Int, Int) {
        let version = binToInt(buf.advance(by: 3))
        let typeId = binToInt(buf.advance(by: 3))

        if typeId == 4 /* literal value */ {
            var numberBits = ""
            var prefix = "1"
            while prefix != "0" {
                prefix = buf.advance(by: 1)
                numberBits += buf.advance(by: 4)
            }
            return (version, binToInt(numberBits))
        }

        var subValues = Array<Int>()
        var versionNumbersSum = version
        let lengthTypeId = buf.advance(by: 1)
        if lengthTypeId == "0" {
            let totalLengthInBits = binToInt(buf.advance(by: 15))
            let subBuf = Buffer(buf.advance(by: totalLengthInBits))
            while subBuf.hasMore() {
                let (subVersion, subValue) = readPacketVersionSumAndValue(subBuf)
                versionNumbersSum += subVersion
                subValues.append(subValue)
            }
        } else {
            let numberOfSubPackets = binToInt(buf.advance(by: 11))
            for _ in 1...numberOfSubPackets {
                let (subVersion, subValue) = readPacketVersionSumAndValue(buf)
                versionNumbersSum += subVersion
                subValues.append(subValue)
            }
        }
        let value: Int
        if typeId == 0 /* sum */ {
            value = subValues.reduce(0, +)
        } else if typeId == 1 /* product */ {
            value = subValues.reduce(1, *)
        } else if typeId == 2 /* minimum */ {
            value = subValues.min()!
        } else if typeId == 3 /* maximum */ {
            value = subValues.max()!
        } else if typeId == 5 /* greater than */ {
            value = subValues[0] > subValues[1] ? 1 : 0
        } else if typeId == 6 /* less than */ {
            value = subValues[0] < subValues[1] ? 1 : 0
        } else if typeId == 7 /* equal to */ {
            value = subValues[0] == subValues[1] ? 1 : 0
        } else {
            fatalError("invalid typeId \(typeId)")
        }
        return (versionNumbersSum, value)
    }

    private func binToInt(_ bin: String) -> Int {
        Int(bin, radix: 2)!
    }

    internal static func allHexToBin(_ hex: String) -> String {
        hex.map({ char in
            let s = String(Int(String(char), radix: 16)!, radix: 2)
            return s.count < 4 ? String(repeating: "0", count: 4 - s.count) + s : s
        }).joined()
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
