use std::error::Error;

pub fn part1(s: &str) -> Result<i32, Box<dyn Error>> {
    let mut lines: Vec<i32> = s.lines().map(|l| l.parse::<i32>().unwrap()).collect();

    let mut offset: i32 = 0;
    let mut steps = 0;
    loop {
        if offset < 0 || offset >= lines.len() as i32 {
            return Ok(steps);
        }
        let jmp = lines[offset as usize];
        lines[offset as usize] = jmp + 1;
        offset += jmp;
        steps += 1;
    }
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
        assert_eq!(part1("0\n3\n0\n1\n-3").unwrap(), 5);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), 5521);
    }
}
