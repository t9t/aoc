use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", count(s)?.0));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", count(s)?.1));
}

fn count(s: &str) -> Result<(u32, u32), Box<dyn Error>> {
    let mut depth = 0;
    let mut ignore_next = false;
    let mut garbage = false;
    let mut score = 0;
    let mut total_garbage = 0;
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
            } else {
                total_garbage += 1;
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
    return Ok((score, total_garbage));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count() {
        assert_eq!(count("{}").unwrap(), (1, 0));
        assert_eq!(count("{{{}}}").unwrap(), (6, 0));
        assert_eq!(count("{{},{}}").unwrap(), (5, 0));
        assert_eq!(count("{{{},{},{{}}}}").unwrap(), (16, 0));
        assert_eq!(count("{<a>,<a>,<a>,<a>}").unwrap(), (1, 4));
        assert_eq!(count("{{<ab>},{<ab>},{<ab>},{<ab>}}").unwrap(), (9, 8));
        assert_eq!(count("{{<!!>},{<!!>},{<!!>},{<!!>}}").unwrap(), (9, 0));
        assert_eq!(count("{{<a!>},{<a!>},{<a!>},{<ab>}}").unwrap(), (3, 17));
    }
}
