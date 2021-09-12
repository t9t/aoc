
import re
from enum import Enum

from itertools import combinations
from itertools import chain
from collections import deque
from collections import Counter


def part1(input: str) -> int:
    floors = read_floors(input)
    return find_min_number_of_moves_to_reach_the_top_floor(floors)


def part2(input: str) -> int:
    floors = read_floors(input)
    first = floors[0]
    first.add((type_generator, "elerium"))
    first.add((type_microchip, "elerium"))
    first.add((type_generator, "dilithium"))
    first.add((type_microchip, "dilithium"))
    return find_min_number_of_moves_to_reach_the_top_floor(floors)


type_microchip = "microchip"
type_generator = "generator"
regex_items = r"(\w+)(?:-compatible)? (microchip|generator)"


def find_min_number_of_moves_to_reach_the_top_floor(floors: list) -> int:
    # Once the top floor has total_items, we have moved everything
    total_items = sum(len(floor) for floor in floors)

    start = (0, 0, floors)
    # Keep a history of visited states, so we don't re-visit them needlessly
    history = set()
    # Implement breath-first-search using a queue
    queue = deque([start])

    while len(queue) > 0:
        num_moves, level, floors = queue.popleft()

        next_states = generate_next_states(num_moves, level, floors)
        for state in next_states:
            key = state_key(state)
            if key in history:
                # Skip already visisted states
                continue

            new_floors = state[2]
            if len(new_floors[3]) == total_items:
                # Top floor contains all the items, we're done here
                return state[0]

            # This is a valid state, add it to the queue to continue our search from that state onwards
            queue.append(state)
            history.add(key)

    raise Exception("No solution found")


def read_floors(input: str) -> list:
    floors = list()
    for line in input.strip().splitlines():
        floor = set()
        matches = re.finditer(regex_items, line)
        for match in matches:
            element, type = match.groups()
            floor.add((type, element))

        floors.append(floor)
    return floors


def generate_next_states(num_moves: int, current_level: int, floors: list) -> list:
    current_floor = floors[current_level]
    single_items = combinations(current_floor, 1)
    double_items = combinations(current_floor, 2)
    # We can either move only a single item or two items at the same time
    move_options = list(chain(single_items, double_items))

    next_states = list()
    for direction in (-1, 1):  # either 1 down or 1 up
        other_level = current_level + direction
        # if direction == -1 and len(floors[other_level]) == 0:
        # Don't bring stuff down to an empty floor
        #    continue
        if other_level < 0 or other_level >= len(floors):
            continue

        for option in move_options:
            # Make sure to copy floors to not accidentally modify references to stored lists
            new_floors = [floor.copy() for floor in floors]
            # Move items in the move option from this level to the other level
            for item in option:
                new_floors[current_level].remove(item)
                new_floors[other_level].add(item)

            # If either of the new floor configuration is unsafe, don't consider it
            if is_safe(new_floors[current_level]) and is_safe(new_floors[other_level]):
                new_state = (num_moves+1, other_level, new_floors)
                next_states.append(new_state)

    return next_states


def is_safe(floor: set) -> bool:
    if len(floor) <= 1:
        # If there are 0 or 1 items, it's safe anyway
        return True

    types = set()
    for item in floor:
        types.add(item[0])

    if len(types) <= 1:
        # If there is just 1 item type, it is always safe
        return True

    # When there is a mix, each microchip should be in a generator, otherwise another generator will break it
    chips, generators = items_by_type(floor)
    for element in chips:
        if element not in generators:
            return False
    return True


def state_key(state) -> tuple:
    return (state[1], floors_key(state[2]))


def floors_key(floors: set) -> str:
    return ",".join(str(floor_key(floor)) for floor in floors)


def floor_key(floor: list) -> tuple:
    # Any floor with the same number of each item type represents an identical state
    chips, generators = items_by_type(floor)
    return (len(chips), len(generators))


def items_by_type(floor):
    chips = [item[1] for item in floor if item[0] == type_microchip]
    generators = [item[1] for item in floor if item[0] == type_generator]
    return chips, generators
