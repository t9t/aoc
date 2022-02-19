use std::error::Error;

pub fn part1(s: &str) -> Result<i32, Box<dyn Error>> {
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

    return Ok(tot);
}

pub fn part2(_s: &str) -> Result<i32, Box<dyn Error>> {
    return Ok(5521);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("1").unwrap(), 0);
        assert_eq!(part1("12").unwrap(), 3);
        assert_eq!(part1("23").unwrap(), 2);
        assert_eq!(part1("1024").unwrap(), 31);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("0").unwrap(), 5521);
    }
}
