
import re
from enum import Enum

from itertools import combinations
from collections import deque
from collections import Counter

regex_floor = r"(\w+)(?:-compatible)? (microchip|generator)"


def part1(input: str):
    floors = read_floors(input)
    return mininum_moves_to_top_floor(floors)


def part2(input: str):
    floors = read_floors(input)
    additional_items = [('elerium', 'generator'), ('elerium', 'microchip'),
                        ('dilithium', 'generator'), ('dilithium', 'microchip')]
    for item in additional_items:
        floors[0].add(item)
    return mininum_moves_to_top_floor(floors)


class ItemType(Enum):
    GENERATOR = "generator"
    CHIP = "microchip"


class Item:
    def __init__(self, type: ItemType, element: str) -> None:
        self.type = type
        self.element = element


class State:
    def __init__(self, moves: int, elevator: int, floors: list) -> None:
        self.moves = moves
        self.elevator = elevator
        self.floors = floors
        self.key = (self.elevator, tuple(tuple(Counter(type for _, type in floor).most_common())
                                         for floor in self.floors))

    def generate_nexts(self):
        this_floor = self.floors[self.elevator]
        doubles = list(combinations(this_floor, 2))  # Move 2 items at once
        singles = list(combinations(this_floor, 1))  # Move just a single item
        movables = set(doubles + singles)  # All possible item combinations we could move
        directions = list()
        if self.elevator > 0:
            if len(self.floors[self.elevator-1]) != 0:
                # No point going down to empty floors
                directions.append(-1)
        if self.elevator < len(self.floors)-1:
            directions.append(1)

        states = list()
        for items_to_move in movables:
            for direction in directions:
                move_to = self.elevator + direction

                unsafe = False
                new_floors = list()
                for i, floor in enumerate(self.floors):
                    cp = floor.copy()
                    if i == self.elevator:
                        for item in items_to_move:
                            cp.remove(item)
                    elif i == move_to:
                        for item in items_to_move:
                            cp.add(item)
                    if not is_safe(cp):
                        unsafe = True
                        break
                    new_floors.append(cp)

                if not unsafe:
                    states.append(State(self.moves + 1, move_to, new_floors))
        return states


def read_floors(input: str):
    floors = list()
    for line in input.strip().splitlines():
        floors.append(set(re.findall(regex_floor, line)))
    return floors


def is_safe(floor: list):
    # No items or just 1 item is always safe
    if len(floor) <= 1:
        return True

    types = set()
    for item in floor:
        types.add(item[1])
        if len(types) == 2:
            break

    # If we only have generators or only chips, it's safe
    if len(types) <= 1:
        return True

    chips = set([item[0] for item in floor if item[1] == "microchip"])
    generators = set([item[0] for item in floor if item[1] == "generator"])

    for chip in chips:
        if chip not in generators:
            # If we have a chip without generator, it will take radiation from another generator -> unsafe
            return False
    return True


def mininum_moves_to_top_floor(floors):
    initial = State(0, 0, floors)
    queue = deque([initial])
    history = set([initial])
    total_items = sum([len(floor) for floor in floors])
    top_floor = len(floors)-1

    while queue:
        state = queue.popleft()
        for next in state.generate_nexts():
            if len(next.floors[top_floor]) == total_items:
                return next.moves

            if next.key not in history:
                history.add(next.key)
                queue.append(next)
