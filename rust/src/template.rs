pub fn part1(s: &str) -> i32 {
    return 1337;
}

pub fn part2(s: &str) -> i32 {
    return 5521;
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT), 1337);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT), 5521);
    }
}
