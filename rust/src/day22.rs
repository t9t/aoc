use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", burst_and_count_infections1(s, 10_000)?));
}

fn burst_and_count_infections1(s: &str, count: u32) -> Result<u32, Box<dyn Error>> {
    let next_dir = |dir: Direction, state: &NodeState| match state {
        NodeState::Clean => dir.turn_left(),
        NodeState::Infected => dir.turn_right(),
        _ => panic!("Should never happen"),
    };

    let next_state = |state: &NodeState| match state {
        NodeState::Clean => NodeState::Infected,
        NodeState::Infected => NodeState::Clean,
        _ => panic!("Should never happen"),
    };

    return burst_and_count_infections(s, count, next_dir, next_state);
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", burst_and_count_infections2(s, 10_000_000)?));
}

fn burst_and_count_infections2(s: &str, count: u32) -> Result<u32, Box<dyn Error>> {
    let next_dir = |dir: Direction, state: &NodeState| match state {
        NodeState::Clean => dir.turn_left(),
        NodeState::Weakened => dir, // Does not turn
        NodeState::Infected => dir.turn_right(),
        NodeState::Flagged => dir.reverse(),
    };

    let next_state = |state: &NodeState| match state {
        NodeState::Clean => NodeState::Weakened,
        NodeState::Weakened => NodeState::Infected,
        NodeState::Infected => NodeState::Flagged,
        NodeState::Flagged => NodeState::Clean,
    };

    return burst_and_count_infections(s, count, next_dir, next_state);
}

fn burst_and_count_infections(
    s: &str,
    count: u32,
    next_dir: fn(Direction, &NodeState) -> Direction,
    next_state: fn(&NodeState) -> NodeState,
) -> Result<u32, Box<dyn Error>> {
    let mut nodes: HashMap<(i32, i32), NodeState> = HashMap::new();

    for (y, line) in s.lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            if c == '#' {
                nodes.insert((x as i32, y as i32), NodeState::Infected);
            }
        }
    }

    let mut infections = 0;
    let mut x = s.lines().next().unwrap().len() as i32 / 2;
    let mut y = s.lines().count() as i32 / 2;
    let mut dir = Direction::Up;

    for _ in 0..count {
        let state = nodes.get(&(x, y)).unwrap_or(&NodeState::Clean);
        dir = next_dir(dir, state);
        let new_state = next_state(state);
        if new_state == NodeState::Infected {
            infections += 1;
        }

        nodes.insert((x, y), new_state);
        let (dx, dy) = dir.delta();
        x += dx;
        y += dy;
    }

    return Ok(infections);
}

#[derive(Debug, PartialEq)]
enum NodeState {
    Clean,
    Weakened,
    Infected,
    Flagged,
}

#[derive(Debug, PartialEq)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    fn delta(&self) -> (i32, i32) {
        match *self {
            Direction::Up => (0, -1),
            Direction::Right => (1, 0),
            Direction::Down => (0, 1),
            Direction::Left => (-1, 0),
        }
    }

    fn turn_left(&self) -> Direction {
        match *self {
            Direction::Up => Direction::Left,
            Direction::Right => Direction::Up,
            Direction::Down => Direction::Right,
            Direction::Left => Direction::Down,
        }
    }

    fn turn_right(&self) -> Direction {
        match *self {
            Direction::Up => Direction::Right,
            Direction::Right => Direction::Down,
            Direction::Down => Direction::Left,
            Direction::Left => Direction::Up,
        }
    }

    fn reverse(&self) -> Direction {
        match *self {
            Direction::Up => Direction::Down,
            Direction::Right => Direction::Left,
            Direction::Down => Direction::Up,
            Direction::Left => Direction::Right,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "..#
#..
...";

    #[test]
    fn test_burst_and_count_infections() {
        assert_eq!(burst_and_count_infections1(INPUT, 7).unwrap(), 5);
        assert_eq!(burst_and_count_infections1(INPUT, 70).unwrap(), 41);
        assert_eq!(burst_and_count_infections1(INPUT, 10_000).unwrap(), 5587);
    }

    #[test]
    fn test_burst_and_count_infections2() {
        assert_eq!(burst_and_count_infections2(INPUT, 100).unwrap(), 26);
        assert_eq!(burst_and_count_infections2(INPUT, 10_000).unwrap(), 2608);
    }
}
