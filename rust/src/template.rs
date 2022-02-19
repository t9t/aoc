use std::error::Error;

pub fn part1(_s: &str) -> Result<i32, Box<dyn Error>> {
    return Ok(1337);
}

pub fn part2(_s: &str) -> Result<i32, Box<dyn Error>> {
    return Ok(5521);
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), 1337);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), 5521);
    }
}
