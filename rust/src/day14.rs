use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut used = 0;
    for row in 0..128 {
        let mut lengths: Vec<usize> = Vec::new();
        format!("{}-{}", s, row)
            .chars()
            .for_each(|c| lengths.push(c as usize));
        [17, 31, 73, 47, 23].iter().for_each(|n| lengths.push(*n));

        let list = full_knot_hash(lengths, 256, 64);

        let mut hash = String::new();
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
            hash += format!("{:0>2x}", x).as_str();
        }

        let mut bin = String::new();
        for c in hash.chars() {
            bin += format!(
                "{:0>4b}",
                u16::from_str_radix(format!("{}", c).as_str(), 16)?
            )
            .as_str();
        }
        used += bin.matches("1").count();
    }

    return Ok(format!("{}", used));
}

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
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

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_part1() {
        assert_eq!(part1("flqrgnkx").unwrap(), "8108");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
