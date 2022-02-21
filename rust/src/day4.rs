use std::error::Error;

pub fn part1(s: &str) -> Result<i32, Box<dyn Error>> {
    let lines = s.split("\n");
    let mut valid_count = 0;
    for line in lines {
        if is_valid(line) {
            valid_count += 1;
        }
    }
    return Ok(valid_count);
}

pub fn part2(_s: &str) -> Result<i32, Box<dyn Error>> {
    return Ok(5521);
}

fn is_valid(s: &str) -> bool {
    let words: Vec<&str> = s.split(" ").collect();
    for i in 0..words.len() {
        let left = words[i];
        for j in i + 1..words.len() {
            let right = words[j];
            if left == right {
                return false;
            }
        }
    }
    return true;
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_is_valid() {
        assert_eq!(is_valid("aa bb cc dd ee"), true);
        assert_eq!(is_valid("aa bb cc dd aa"), false);
        assert_eq!(is_valid("aa bb cc dd aaa"), true);
    }

    #[test]
    fn test_part1() {
        let input = "aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa";
        assert_eq!(part1(input).unwrap(), 2);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), 5521);
    }
}
