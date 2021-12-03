import Foundation

let inputDir = getInputDir()
let inputFile = getInputDir() + "/2021/3.txt"
let input = try String(contentsOfFile: inputFile)

print(Day3(input).part2())

func getInputDir() -> String {
    // When running in AppCode it uses a .build/debug/ directory as working directory, so "../input" does not work
    if let inputDir = ProcessInfo.processInfo.environment["INPUT_DIR"] {
        return inputDir
    }
    // Assuming aoc/swift directory here
    return "../input"
}