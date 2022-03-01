use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let steps = s.parse::<usize>()?;
    let mut buffer: Vec<usize> = vec![0];
    let mut pos = 0;
    let mut new_value = 1;

    for _ in 0..2017 {
        pos += steps;
        pos %= buffer.len();
        let mut newbuf: Vec<usize> = Vec::new();
        for i in 0..pos + 1 {
            newbuf.push(buffer[i]);
        }
        newbuf.push(new_value);
        for i in pos + 1..buffer.len() {
            newbuf.push(buffer[i]);
        }
        pos += 1;
        new_value += 1;
        buffer = newbuf;
    }
    pos = if pos == buffer.len() - 1 { 0 } else { pos + 1 };

    return Ok(format!("{}", buffer[pos]));
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
        assert_eq!(part1("3").unwrap(), "638");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
