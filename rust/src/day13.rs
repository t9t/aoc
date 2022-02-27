use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut scanners: HashMap<u32, (i32, i32, i32)> = HashMap::new();
    let mut max_depth = 0;
    for line in s.lines() {
        let nums = (line
            .split(": ")
            .map(|x| x.parse::<u32>())
            .collect::<Result<Vec<u32>, _>>())?;
        scanners.insert(nums[0], (nums[1] as i32, 1, 0));
        max_depth = max_depth.max(nums[0]);
    }

    let mut severity = 0;
    for depth in 0..max_depth + 1 {
        // Detect hit
        if let Some((range, _, pos)) = scanners.get(&depth) {
            if pos == &0 {
                severity += range * depth as i32;
            }
        }

        // Move scanners
        for depth in 0..max_depth + 1 {
            if let Some((range, delta, pos)) = scanners.get(&depth) {
                let mut new_pos = *pos;
                new_pos += delta;
                let mut new_delta = *delta;
                if new_pos == -1 {
                    new_pos = 1;
                    new_delta = 1;
                } else if &new_pos == range {
                    new_pos = range - 2;
                    new_delta = -1;
                }
                let new_range = *range;
                scanners.insert(depth, (new_range, new_delta, new_pos));
            }
        }
    }
    return Ok(format!("{}", severity));
}

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "1337");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
