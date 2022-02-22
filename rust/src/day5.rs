use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return steps_to_exit(s, |_| 1);
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return steps_to_exit(s, |offset| if offset >= 3 { -1 } else { 1 });
}

fn steps_to_exit(s: &str, d: fn(i32) -> i32) -> Result<String, Box<dyn Error>> {
    let mut lines: Vec<i32> = s.lines().map(|l| l.parse::<i32>().unwrap()).collect();

    let mut pos: i32 = 0;
    let mut steps = 0;
    loop {
        if pos < 0 || pos >= lines.len() as i32 {
            return Ok(format!("{}", steps));
        }
        let offset = lines[pos as usize];
        lines[pos as usize] += d(offset);
        pos += offset;
        steps += 1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("0\n3\n0\n1\n-3").unwrap(), "5");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("0\n3\n0\n1\n-3").unwrap(), "10");
    }
}
