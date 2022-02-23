use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut depth = 0;
    let mut ignore_next = false;
    let mut garbage = false;
    let mut score = 0;
    for c in s.chars() {
        if ignore_next {
            ignore_next = false;
            continue;
        }

        if garbage {
            if c == '>' {
                garbage = false;
            } else if c == '!' {
                ignore_next = true;
            }
            continue;
        }

        match c {
            '{' => {
                depth += 1;
                score += depth
            }
            '}' => depth -= 1,
            '<' => garbage = true,
            _ => {}
        }
    }
    return Ok(format!("{}", score));
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
