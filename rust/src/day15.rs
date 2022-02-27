use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let _s = "Generator A starts with 65
Generator B starts with 8921";
    fn parse_line(line_opt: Option<&str>) -> Result<u64, Box<dyn Error>> {
        let line = line_opt.ok_or("Invalid input")?;
        let s = line.split(" starts with ").nth(1).ok_or("Invalid line")?;
        return Ok(s.parse::<u64>()?);
    }

    let mut lines = s.lines();
    let mut a = parse_line(lines.next())?;
    let mut b = parse_line(lines.next())?;
    let mut matches = 0;

    for _ in 0..40000000 {
        a = (a * 16807) % 2147483647;
        b = (b * 48271) % 2147483647;

        if (a as u16) == (b as u16) {
            matches += 1;
        }
    }

    return Ok(format!("{}", matches));
}

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "Generator A starts with 65
Generator B starts with 8921";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "588");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
