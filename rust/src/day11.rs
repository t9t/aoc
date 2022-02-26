use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let (steps, _) = find_steps(s)?;

    return Ok(format!("{}", steps));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let (_, furthest) = find_steps(s)?;

    return Ok(format!("{}", furthest));
}

fn find_steps(s: &str) -> Result<(i32, i32), Box<dyn Error>> {
    let route: Vec<&str> = s.split(",").collect();

    let mut x = 0 as i32;
    let mut y = 0 as i32;
    let mut z = 0 as i32;
    let mut step_count = 0;
    let mut furthest = 0;
    for step in route {
        match step {
            "n" => {
                z += 1;
                y -= 1;
            }
            "ne" => {
                x += 1;
                y -= 1;
            }
            "se" => {
                x += 1;
                z -= 1;
            }
            "s" => {
                y += 1;
                z -= 1;
            }
            "sw" => {
                x -= 1;
                y += 1;
            }
            "nw" => {
                x -= 1;
                z += 1;
            }
            _ => Err(format!("Invalid direction {}", step))?,
        }
        step_count = (x.abs() + y.abs() + z.abs()) / 2;
        if step_count > furthest {
            furthest = step_count;
        }
    }
    return Ok((step_count, furthest));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("ne,ne,ne").unwrap(), "3");
        assert_eq!(part1("ne,ne,sw,sw").unwrap(), "0");
        assert_eq!(part1("ne,ne,s,s").unwrap(), "2");
        assert_eq!(part1("se,sw,se,sw,sw").unwrap(), "3");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("ne,ne,ne").unwrap(), "3");
        assert_eq!(part2("ne,ne,sw,sw").unwrap(), "2");
        assert_eq!(part2("ne,ne,s,s").unwrap(), "2");
        assert_eq!(part2("se,sw,se,sw,sw").unwrap(), "3");
    }
}
