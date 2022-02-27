use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let used = build_grid(s)?
        .iter()
        .flat_map(|row| row.chars())
        .filter(|c| *c == '1')
        .count();
    return Ok(format!("{}", used));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let grid = build_grid(s)?;

    let mut groups: HashMap<(i32, i32), u32> = HashMap::new();
    let mut group_counter = 0;
    for (y, row) in grid.iter().enumerate() {
        for (x, c) in row.chars().enumerate() {
            if c == '0' {
                continue;
            }
            let coords = (x as i32, y as i32);
            if groups.contains_key(&coords) {
                continue;
            }

            let mut next: Vec<(i32, i32)> = Vec::new();
            next.push(coords);

            while !next.is_empty() {
                let (nx, ny) = next.pop().unwrap();
                let nc = &grid[ny as usize].chars().nth(nx as usize).unwrap();
                if *nc != '1' {
                    continue;
                }

                groups.insert((nx, ny), group_counter);
                for o in [(nx + 1, ny), (nx - 1, ny), (nx, ny + 1), (nx, ny - 1)] {
                    if o.0 >= 0 && o.1 >= 0 && o.0 <= 127 && o.1 <= 127 && !groups.contains_key(&o)
                    {
                        next.push(o);
                    }
                }
            }
            group_counter += 1;
        }
    }

    return Ok(format!("{}", group_counter));
}

fn build_grid(key: &str) -> Result<Vec<String>, Box<dyn Error>> {
    let mut grid: Vec<String> = Vec::new();
    for row in 0..128 {
        let mut lengths: Vec<usize> = Vec::new();
        format!("{}-{}", key, row)
            .chars()
            .for_each(|c| lengths.push(c as usize));
        [17, 31, 73, 47, 23].iter().for_each(|n| lengths.push(*n));

        let list = full_knot_hash(lengths, 256, 64);

        let mut row = String::new();

        for block in 0..16 {
            let mut x: i32 = -1;
            for k in 0..16 {
                let n = list[block * 16 + k] as i32;
                if x == -1 {
                    x = n;
                } else {
                    x = x ^ n;
                }
            }

            for c in format!("{:0>2x}", x).chars() {
                row += format!(
                    "{:0>4b}",
                    u16::from_str_radix(format!("{}", c).as_str(), 16)?
                )
                .as_str();
            }
        }

        grid.push(row);
    }
    return Ok(grid);
}

fn full_knot_hash(lengths: Vec<usize>, list_size: usize, rounds: u32) -> Vec<usize> {
    let mut list: Vec<usize> = Vec::new();
    for i in 0..list_size {
        list.push(i);
    }

    let mut pos = 0;
    let mut skip = 0;
    for _ in 0..rounds {
        for l in &lengths {
            let end_pos = pos + l - 1;

            for k in 0..l / 2 {
                let li = (pos + k) % list.len();
                let ri = (end_pos - k) % list.len();

                let left = list[li];
                let right = list[ri];

                list[li] = right;
                list[ri] = left;
            }

            pos += l + skip;
            skip += 1;
        }
    }
    return list;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("flqrgnkx").unwrap(), "8108");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("flqrgnkx").unwrap(), "1242");
    }
}
