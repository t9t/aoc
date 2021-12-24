import Foundation

class Day23: Day {
    private let inputLines: Array<Substring>

    init(_ input: String) {
        inputLines = input.trimmingCharacters(in: .whitespacesAndNewlines).split(separator: "\n")
    }

    func part1() -> Int {
        let a = Room(top: "B", bottom: "A")
        let b = Room(top: "C", bottom: "D")
        let c = Room(top: "B", bottom: "C")
        let d = Room(top: "D", bottom: "A")
        let hallway = [Int: Character]()
        let rooms = [a, b, c, d]

        let state = State(rooms: rooms, hallway: hallway)
        let minEnergyLevel = determineMinEnergyLevelToEnd(state, energyLevelOfTheStepGettingHere: 0)

        return minEnergyLevel!
    }

    func part2() -> Int {
        return 1337
    }

    private let targetRoomIndexForPod: [Character: Int] = ["A": 0, "B": 1, "C": 2, "D": 3]
    private let targetPodForRoomIndex: [Int: Character] = [0: "A", 1: "B", 2: "C", 3: "D"]
    private let hallwayPositionAboveRoomWithIndex: [Int: Int] = [0: 3, 1: 5, 2: 7, 3: 9]
    private let energyLevelPerStepForPod: [Character: Int] = ["A": 1, "B": 10, "C": 100, "D": 1000]

    private struct State: Equatable, Hashable {
        let rooms: [Room], hallway: [Int: Character]
    }

    private var cache = [State: Int?]()

    private func determineMinEnergyLevelToEnd(_ state: State, energyLevelOfTheStepGettingHere: Int) -> Int? {
        let rooms = state.rooms, hallway = state.hallway

        let a = rooms[0]
        let b = rooms[1]
        let c = rooms[2]
        let d = rooms[3]
        if a.allSpotsFilledWith("A") && b.allSpotsFilledWith("B") && c.allSpotsFilledWith("C") && d.allSpotsFilledWith("D") {
            return energyLevelOfTheStepGettingHere
        }

        if let outcome = cache[state] {
            return outcome == nil ? nil : outcome! + energyLevelOfTheStepGettingHere
        }

        var minEnergyLevelGettingToEndAfterThisStep: Int? = nil

        // Try to move everything from the hallway into rooms
        for (hallwayPos, pod) in hallway {
            let targetRoomIndex = targetRoomIndexForPod[pod]!
            let targetRoom = rooms[targetRoomIndex]
            if targetRoom.isFull {
                continue
            }

            if !targetRoom.isEmpty {
                if !targetRoom.allFilledSpotsAre(pod) {
                    // There are non-matching pods in the room, cannot move in
                    continue
                }
            }

            let targetRoomHallwayPosition = hallwayPositionAboveRoomWithIndex[targetRoomIndex]!
            var pathObstructed = false
            for inBetweenPos in min(hallwayPos, targetRoomHallwayPosition)...max(hallwayPos, targetRoomHallwayPosition) {
                if inBetweenPos == hallwayPos {
                    // TODO: nicer
                    continue
                }
                if hallway[inBetweenPos] != nil {
                    pathObstructed = true
                    break
                }
            }

            if pathObstructed {
                continue
            }

            var modifiedHallway = hallway
            modifiedHallway.removeValue(forKey: hallwayPos) // Clear this position of the hallway
            let modifiedTargetRoom = targetRoom.push(pod)
            var modifiedRooms = rooms
            modifiedRooms[targetRoomIndex] = modifiedTargetRoom

            let horizontalSteps = abs(hallwayPos - targetRoomHallwayPosition)
            let verticalSteps = targetRoom.numberOfEmptySpots
            let totalSteps = horizontalSteps + verticalSteps
            let thisMoveEnergyLevel = totalSteps * energyLevelPerStepForPod[pod]!

            let newState = State(rooms: modifiedRooms, hallway: modifiedHallway)
            if let energyLevelUntilEnd = determineMinEnergyLevelToEnd(newState, energyLevelOfTheStepGettingHere: thisMoveEnergyLevel) {
                if minEnergyLevelGettingToEndAfterThisStep == nil {
                    minEnergyLevelGettingToEndAfterThisStep = energyLevelUntilEnd
                } else {
                    minEnergyLevelGettingToEndAfterThisStep = min(minEnergyLevelGettingToEndAfterThisStep!, energyLevelUntilEnd)
                }
            }
        }

        if hallway.count == (11 - 4) {
            // Hallway is full, cannot move anything out of rooms
            // TODO: impl
        }

        // Try to move anything from rooms into the hallway
        for (roomIndex, room) in rooms.enumerated() {
            if room.isEmpty {
                // Room is empty, nothing to move
                continue
            }
            let targetPodForThisRoom = targetPodForRoomIndex[roomIndex]!
            if room.allFilledSpotsAre(targetPodForThisRoom) {
                // Everything in the room is the target pod, nothing to move
                continue
            }

            let roomHallwayPosition = hallwayPositionAboveRoomWithIndex[roomIndex]!
            for hallwayPos in [1, 2, 4, 6, 8, 10, 11] {
                if hallway[hallwayPos] != nil {
                    // Position occupied, cannot move into it
                    continue
                }

                var pathObstructed = false
                for inBetweenPos in min(hallwayPos, roomHallwayPosition)...max(hallwayPos, roomHallwayPosition) {
                    if hallway[inBetweenPos] != nil {
                        pathObstructed = true
                        break
                    }
                }

                if pathObstructed {
                    continue
                }

                // Open candidate position in hallway, move in here
                let (podToMove, modifiedRoom) = room.pop()
                var modifiedRooms = rooms
                modifiedRooms[roomIndex] = modifiedRoom

                var modifiedHallway = hallway
                modifiedHallway[hallwayPos] = podToMove

                let horizontalSteps = abs(hallwayPos - roomHallwayPosition)
                let verticalSteps = modifiedRoom.numberOfEmptySpots
                let totalSteps = horizontalSteps + verticalSteps
                let thisMoveEnergyLevel = totalSteps * energyLevelPerStepForPod[podToMove]!

                let newState = State(rooms: modifiedRooms, hallway: modifiedHallway)

                if let energyLevelUntilEnd = determineMinEnergyLevelToEnd(newState, energyLevelOfTheStepGettingHere: thisMoveEnergyLevel) {
                    if minEnergyLevelGettingToEndAfterThisStep == nil {
                        minEnergyLevelGettingToEndAfterThisStep = energyLevelUntilEnd
                    } else {
                        minEnergyLevelGettingToEndAfterThisStep = min(minEnergyLevelGettingToEndAfterThisStep!, energyLevelUntilEnd)
                    }
                }
            }
        }

        cache[state] = minEnergyLevelGettingToEndAfterThisStep
        if minEnergyLevelGettingToEndAfterThisStep == nil {
            return nil
        }

        let totalMinEnergyToGetToEndIncludingThis = minEnergyLevelGettingToEndAfterThisStep! + energyLevelOfTheStepGettingHere
        return totalMinEnergyToGetToEndIncludingThis
    }

    private struct Room: Hashable, Equatable, CustomStringConvertible {
        let top: Character?, bottom: Character?

        var isEmpty: Bool {
            get {
                top == nil && bottom == nil
            }
        }

        var isFull: Bool {
            get {
                top != nil && bottom != nil
            }
        }

        var numberOfEmptySpots: Int {
            get {
                isEmpty ? 2 : (top == nil ? 1 : 0)
            }
        }

        func allSpotsFilledWith(_ c: Character) -> Bool {
            top == c && bottom == c
        }

        func allFilledSpotsAre(_ c: Character) -> Bool {
            if isEmpty {
                return false
            } else if top == nil {
                return bottom == c
            }
            return allSpotsFilledWith(c)
        }

        func pop() -> (Character, Room) {
            if isEmpty {
                fatalError("pop() on empty room")
            }
            if top != nil {
                return (top!, Room(top: nil, bottom: bottom))
            }
            return (bottom!, Room(top: nil, bottom: nil))
        }

        func push(_ c: Character) -> Room {
            if isFull {
                fatalError("push() on full room")
            }

            if bottom == nil {
                return Room(top: nil, bottom: c)
            }
            return Room(top: c, bottom: bottom)
        }

        var description: String {
            "Room(top: \(top ?? "-"), bottom: \(bottom ?? "-"))"
        }
    }
}
