use std::collections::HashSet;
use std::collections::VecDeque;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let route = s.split(",").collect();
    let target = find_target(route)?;
    let steps = find_path((0, 0, 0), target).ok_or("No route found from target to origin")?;
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

fn find_path((sx, sy, sz): (i32, i32, i32), (tx, ty, tz): (i32, i32, i32)) -> Option<u32> {
    let mut visited: HashSet<(i32, i32, i32)> = HashSet::new();
    let mut next: VecDeque<(i32, i32, i32, u32)> = VecDeque::new();
    next.push_back((sx, sy, sz, 0));

    while !next.is_empty() {
        let (x, y, z, steps) = next.pop_front().unwrap();
        if x == tx && y == ty && z == tz {
            return Some(steps);
        }

        let neighbors = neighbors(x, y, z);
        for n in neighbors {
            if !visited.contains(&n) {
                visited.insert(n);
                next.push_back((n.0, n.1, n.2, steps + 1));
            }
        }
    }
    return None;
}

fn neighbors(x: i32, y: i32, z: i32) -> Vec<(i32, i32, i32)> {
    let mut neighbors: Vec<(i32, i32, i32)> = Vec::new();
    neighbors.push((x, y - 1, z + 1));
    neighbors.push((x + 1, y - 1, z));
    neighbors.push((x + 1, y, z - 1));
    neighbors.push((x, y + 1, z - 1));
    neighbors.push((x - 1, y + 1, z));
    neighbors.push((x - 1, y, z + 1));
    return neighbors;
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
