use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let n = s.parse::<i32>()?;

    let mut next_opt: Option<i32> = None;
    for x in (1..1000).step_by(2) {
        let pow = x * x;
        if pow >= n {
            next_opt = Some(x);
            break;
        }
    }
    let next = next_opt.ok_or("No next power found")?;
    let steps = next / 2;
    let d = next * next - n;

    let nm1 = next - 1;
    let bottom_right = 0;
    let bottom_left = nm1;
    let top_left = nm1 * 2;
    let top_right = nm1 * 3;

    let mut tot = 0;
    if d == bottom_right {
        tot = steps * 2;
    } else if d == bottom_left {
        tot = steps * 2;
    } else if d == top_left {
        tot = steps * 2;
    } else if d == top_right {
        tot = steps * 2;
    } else if d > bottom_right && d < bottom_left {
        tot = steps + (d - steps).abs();
    } else if d > bottom_left && d < top_left {
        tot = steps + (d - bottom_left - steps).abs();
    } else if d > top_left && d < top_right {
        tot = steps + (d - top_left - steps).abs();
    } else if d > top_right {
        tot = steps + (d - top_right - steps).abs();
    }

    return Ok(format!("{}", tot));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let t = s.parse::<i32>()?;

    enum Direction {
        UP,
        LEFT,
        DOWN,
        RIGHT,
    }

    let mut n = 1;
    let mut x = 0;
    let mut y = 0;
    let mut next_turn_at = 2;
    let mut dx = 1;
    let mut dy = 0;
    let mut dir = Direction::RIGHT;
    let mut side_length = 1;
    let mut grid: HashMap<(i32, i32), i32> = HashMap::new();
    grid.insert((0, 0), 1);
    loop {
        n += 1;
        x += dx;
        y += dy;

        let mut sum = 0;
        for ox in (x - 1)..(x + 2) {
            for oy in (y - 1)..(y + 2) {
                if ox != x || oy != y {
                    sum += *grid.get(&(ox, oy)).unwrap_or(&0);
                }
            }
        }
        if sum > t {
            return Ok(format!("{}", sum));
        }
        grid.insert((x, y), sum);

        if n == next_turn_at {
            // "turn left"
            dir = match dir {
                Direction::UP => Direction::LEFT,
                Direction::LEFT => Direction::DOWN,
                Direction::DOWN => Direction::RIGHT,
                Direction::RIGHT => Direction::UP,
            };
            let delta = match dir {
                Direction::UP => (0, -1),
                Direction::LEFT => (-1, 0),
                Direction::DOWN => (0, 1),
                Direction::RIGHT => (1, 0),
            };
            dx = delta.0;
            dy = delta.1;
            if matches!(dir, Direction::RIGHT) || matches!(dir, Direction::LEFT) {
                side_length += 1;
            }
            next_turn_at += side_length;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("1").unwrap(), "0");
        assert_eq!(part1("12").unwrap(), "3");
        assert_eq!(part1("23").unwrap(), "2");
        assert_eq!(part1("1024").unwrap(), "31");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("3").unwrap(), "4");
        assert_eq!(part2("120").unwrap(), "122");
        assert_eq!(part2("350").unwrap(), "351");
        assert_eq!(part2("351").unwrap(), "362");
        assert_eq!(part2("800").unwrap(), "806");
    }
}
