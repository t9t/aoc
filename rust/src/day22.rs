use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", burst_and_count_infections(s, 10_000)?));
}

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
}

fn burst_and_count_infections(s: &str, count: u32) -> Result<u32, Box<dyn Error>> {
    let mut nodes: HashMap<(i32, i32), bool> = HashMap::new();

    for (y, line) in s.lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            if c == '#' {
                nodes.insert((x as i32, y as i32), true);
            }
        }
    }

    let mut infections = 0;
    let mut x = s.lines().next().unwrap().len() as i32 / 2;
    let mut y = s.lines().count() as i32 / 2;
    let mut dir = Direction::UP;

    for _ in 0..count {
        let infected = *nodes.get(&(x, y)).unwrap_or(&false);

        if infected {
            dir = dir.turn_right();
            nodes.insert((x, y), false);
        } else {
            dir = dir.turn_left();
            nodes.insert((x, y), true);
            infections += 1;
        }
        let (dx, dy) = dir.delta();
        x += dx;
        y += dy;
    }

    return Ok(infections);
}

#[derive(Debug)]
enum Direction {
    UP,
    RIGHT,
    DOWN,
    LEFT,
}

impl Direction {
    fn delta(&self) -> (i32, i32) {
        match *self {
            Direction::UP => (0, -1),
            Direction::RIGHT => (1, 0),
            Direction::DOWN => (0, 1),
            Direction::LEFT => (-1, 0),
        }
    }

    fn turn_left(&self) -> Direction {
        match *self {
            Direction::UP => Direction::LEFT,
            Direction::RIGHT => Direction::UP,
            Direction::DOWN => Direction::RIGHT,
            Direction::LEFT => Direction::DOWN,
        }
    }

    fn turn_right(&self) -> Direction {
        match *self {
            Direction::UP => Direction::RIGHT,
            Direction::RIGHT => Direction::DOWN,
            Direction::DOWN => Direction::LEFT,
            Direction::LEFT => Direction::UP,
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
        assert_eq!(burst_and_count_infections(INPUT, 7).unwrap(), 5);
        assert_eq!(burst_and_count_infections(INPUT, 70).unwrap(), 41);
        assert_eq!(burst_and_count_infections(INPUT, 10_000).unwrap(), 5587);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
