use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let (scanners, max_depth) = parse_scanners(s)?;
    return Ok(format!("{}", severity(&scanners, max_depth, 0).0));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let (scanners, max_depth) = parse_scanners(s)?;
    let mut delay = 0;
    loop {
        if severity(&scanners, max_depth, delay).1 == 0 {
            return Ok(format!("{}", delay));
        }
        delay += 1;
    }
}

fn parse_scanners(s: &str) -> Result<(HashMap<u32, u32>, u32), Box<dyn Error>> {
    let mut scanners: HashMap<u32, u32> = HashMap::new();
    let mut max_depth = 0;
    for line in s.lines() {
        let nums = (line
            .split(": ")
            .map(|x| x.parse::<u32>())
            .collect::<Result<Vec<u32>, _>>())?;
        scanners.insert(nums[0], nums[1]);
        max_depth = max_depth.max(nums[0]);
    }
    return Ok((scanners, max_depth));
}

fn severity(scanners: &HashMap<u32, u32>, max_depth: u32, delay: u32) -> (u32, u32) {
    let mut severity = 0;
    let mut hit_count = 0;
    for depth in 0..max_depth + 1 {
        let t = delay + depth;
        if let Some(range) = scanners.get(&depth) {
            if t % ((range - 1) * 2) == 0 {
                severity += range * depth;
                hit_count += 1;
            }
        }
    }
    return (severity, hit_count);
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "0: 3
1: 2
4: 4
6: 4";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "24");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "10");
    }
}
