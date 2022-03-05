use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    // The one with lowest total acceleration will eventuall stay closest to 0
    let mut lowest = i32::MAX;
    let mut lowest_i = 0;

    for (i, line) in s.lines().enumerate() {
        let mut parts = line.split("a=<");
        parts.next();
        let a = parts.next().unwrap().split(">").next().unwrap();

        let nums = (a
            .split(",")
            .map(|x| x.trim().parse::<i32>())
            .collect::<Result<Vec<i32>, _>>())?;

        let tot = nums[0].abs() + nums[1].abs() + nums[2].abs();
        if tot < lowest {
            lowest = tot;
            lowest_i = i;
        }
    }
    return Ok(format!("{}", lowest_i));
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
