use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(walk(s)?.0);
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", walk(s)?.1));
}

fn walk(s: &str) -> Result<(String, u32), Box<dyn Error>> {
    let grid = s.lines().collect::<Vec<&str>>();
    let height = grid.len() as i32;
    let width = grid[0].len() as i32;

    let mut x: i32;
    let mut y: i32 = 0;
    let mut dx: i32 = 0;
    let mut dy: i32 = 1;

    // Find starting position, always in row 0, going down
    let mut start_pos: Option<i32> = None;
    for (i, c) in grid[0].chars().enumerate() {
        if c == '|' {
            start_pos = Some(i as i32);
            break;
        }
    }
    x = start_pos.ok_or("No starting pos found")?;

    let mut letters: Vec<char> = Vec::new();
    let mut steps = 0;

    loop {
        x += dx;
        y += dy;
        steps += 1;

        if x < 0 || y < 0 || x >= width || y >= height {
            // Out of bounds, must has finished walking
            break;
        }

        let c = grid[y as usize].chars().nth(x as usize).unwrap();
        if c == '+' {
            let previous_x = x - dx;
            let previous_y = y - dy;
            for (new_dx, new_dy) in [(-1, 0), (1, 0), (0, -1), (0, 1)] {
                let new_x = x + new_dx;
                let new_y = y + new_dy;
                if new_x == previous_x && new_y == previous_y {
                    // We came from here
                } else if new_x < 0 || new_y < 0 || new_x >= width || new_y >= height {
                    // Out of bounds
                } else {
                    let next_c = grid[new_y as usize].chars().nth(new_x as usize).unwrap();
                    if next_c != ' ' {
                        dx = new_dx;
                        dy = new_dy;
                        break;
                    }
                }
            }
        } else if c >= 'A' && c <= 'Z' {
            letters.push(c);
        } else if c == ' ' {
            // Must have finished walking
            break;
        } else if c != '|' && c != '-' {
            Err(format!("Invalid char: {}", c))?;
            break;
        }
    }

    return Ok((letters.iter().collect(), steps));
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "     |          
     |  +--+    
     A  |  C    
 F---|--|-E---+ 
     |  |  |  D 
     +B-+  +--+ ";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "ABCDEF");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "38");
    }
}
