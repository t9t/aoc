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
        let expanded = expand(inputImage, by: 3)
        let enhancedOnce = enhance(expanded)
        let enhancedTwice = enhance(enhancedOnce)

        #if false
        print("Expanded")
        printImage(expanded)
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
        var image = expand(inputImage, by: 5)

        for _ in 1...50 {
            image = enhance(image)
        }

        return countLitPixels(image)
    }

    private func expand(_ image: Image, by: Int) -> Image {
        var out = Image()
        for rowNum in -by...image.count - 1 + by {
            var row = Array<Bool>()
            for colNum in -by...image[0].count-1+by {
                let pixel : Bool
                if rowNum < 0 || colNum < 0 || rowNum > image.count-1 || colNum > image[0].count-1 {
                    pixel = false
                } else {
                    pixel = inputImage[rowNum][colNum]
                }
                row.append(pixel)
            }
            out.append(row)
        }
        return out
    }

    private func enhance(_ inputImage: Image) -> Image {
        // This "works" because we expand the image by 5 each time, so new sections at the edges are always fully
        // empty, simulating the situation for "infinity" (i.e. all image pixels which exist but are not in the array)
        let currentInfinity = inputImage[0][0]

        var outputImage = Image()
        let expansion = 1
        let maxY = inputImage.count - 1, maxX = inputImage[0].count - 1
        let newRowLength = inputImage[0].count + expansion

        func getNumberFor(x: Int, y: Int) -> Int {
            var bin = ""
            for dy in y - 1...y + 1 {
                for dx in x - 1...x + 1 {
                    let pixel: Bool
                    if dy < 0 || dx < 0 || dy > maxY || dx > maxX {
                        pixel = currentInfinity
                    } else {
                        pixel = inputImage[dy][dx]
                    }
                    bin += (pixel ? "1" : "0")
                }
            }
            return Int(bin, radix: 2)!
        }

        for rowNum in -1...inputImage.count + expansion {
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
