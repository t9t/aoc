import Foundation

class Day20: Day {
    private typealias Image = Array<Array<Bool>>

    private let algorithm: Array<Bool>
    private let inputImage: Image

    init(_ input: String) {
        let inputParts = input.components(separatedBy: "\n\n")
        algorithm = inputParts[0].trimmingCharacters(in: .whitespacesAndNewlines).map({ $0 == "#" })
        inputImage = inputParts[1].trimmingCharacters(in: .whitespacesAndNewlines)
                .split(separator: "\n").map({ row in row.map({ $0 == "#" }) })
    }

    func part1() -> Int {
        let enhancedOnce = enhance(inputImage)
        let enhancedTwice = enhance(enhancedOnce)

        #if false
        print("Input image")
        printImage(inputImage)
        print("Enhanced once")
        printImage(enhancedOnce)
        print("Enhanced twice")
        printImage(enhancedTwice)
        #endif

        return countLitPixels(enhancedTwice)
    }

    func part2() -> Int {
        return 1337
    }

    private func enhance(_ inputImage: Image) -> Image {
        var outputImage = Image()
        let maxY = inputImage.count - 1, maxX = inputImage[0].count - 1
        let newRowLength = inputImage[0].count + 1

        func getNumberFor(x: Int, y: Int) -> Int {
            var bin = ""
            for dy in y - 1...y + 1 {
                for dx in x - 1...x + 1 {
                    let pixel: Bool
                    if dy < 0 || dx < 0 || dy > maxY || dx > maxX {
                        pixel = false
                    } else {
                        pixel = inputImage[dy][dx]
                    }
                    bin += (pixel ? "1" : "0")
                }
            }
            return Int(bin, radix: 2)!
        }

        for rowNum in -1...inputImage.count + 1 {
            var outputRow = Array<Bool>()
            for colNum in -1...newRowLength {
                let num = getNumberFor(x: colNum, y: rowNum)
                outputRow.append(algorithm[algorithm.index(algorithm.startIndex, offsetBy: num)])
            }
            outputImage.append(outputRow)
        }
        return outputImage
    }

    private func countLitPixels(_ image: Image) -> Int {
        image.map({ $0.filter({ $0 }).count }).reduce(0, +)
    }

    private func printImage(_ image: Image) {
        for row in image {
            for pixel in row {
                print(pixel ? "#" : ".", terminator: "")
            }
            print("")
        }
    }
}
