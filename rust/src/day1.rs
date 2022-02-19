use std::error::Error;

pub fn part1(s: &str) -> Result<i32, Box<dyn Error>> {
    return calc_sum(s, 1);
}

pub fn part2(s: &str) -> Result<i32, Box<dyn Error>> {
    let l = s.len();
    return calc_sum(s, l / 2);
}

fn calc_sum(s: &str, x: usize) -> Result<i32, Box<dyn Error>> {
    let chars: Vec<char> = s.chars().collect();
    let len = chars.len();
    let mut sum: i32 = 0;
    for (i, c) in chars.iter().enumerate() {
        let next_pos = (i + x) % len;
        let next = chars[next_pos];
        if *c == next {
            sum += c.to_digit(10).ok_or("Invalid digit")? as i32
        }
    }
    return Ok(sum);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("1122").unwrap(), 3);
        assert_eq!(part1("1111").unwrap(), 4);
        assert_eq!(part1("1234").unwrap(), 0);
        assert_eq!(part1("91212129").unwrap(), 9);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("1212").unwrap(), 6);
        assert_eq!(part2("1221").unwrap(), 0);
        assert_eq!(part2("123425").unwrap(), 4);
        assert_eq!(part2("123123").unwrap(), 12);
        assert_eq!(part2("12131415").unwrap(), 4);
    }
}
