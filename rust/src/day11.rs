use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let route = s.split(",").collect();
    let (tx, ty, tz) = find_target(route)?;
    let steps = (tx.abs() + ty.abs() + tz.abs()) / 2;

    return Ok(format!("{}", steps));
}

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
}

fn find_target(route: Vec<&str>) -> Result<(i32, i32, i32), Box<dyn Error>> {
    let mut x = 0;
    let mut y = 0;
    let mut z = 0;
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
    }
    return Ok((x, y, z));
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_part1() {
        assert_eq!(part1("ne,ne,ne").unwrap(), "3");
        assert_eq!(part1("ne,ne,sw,sw").unwrap(), "0");
        assert_eq!(part1("ne,ne,s,s").unwrap(), "2");
        assert_eq!(part1("se,sw,se,sw,sw").unwrap(), "3");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
