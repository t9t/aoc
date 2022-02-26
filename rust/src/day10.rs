use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", knot_hash(s, 256)?));
}

fn knot_hash(s: &str, list_size: u32) -> Result<u32, Box<dyn Error>> {
    let lengths = s
        .split(",")
        .map(|n| n.parse::<usize>())
        .collect::<Result<Vec<usize>, _>>()?;

    let mut list: Vec<u32> = Vec::new();
    for i in 0..list_size {
        list.push(i);
    }

    let mut pos = 0;
    let mut skip = 0;
    for l in lengths {
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

    return Ok(list[0] * list[1]);
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let mut lengths: Vec<usize> = Vec::new();
    s.chars().for_each(|c| lengths.push(c as usize));
    [17, 31, 73, 47, 23].iter().for_each(|n| lengths.push(*n));

    let list_size = 256;
    let mut list: Vec<usize> = Vec::new();
    for i in 0..list_size {
        list.push(i);
    }

    let mut pos = 0;
    let mut skip = 0;
    for _ in 0..64 {
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

    return Ok(hash);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_knot_hash() {
        assert_eq!(knot_hash("3,4,1,5", 5).unwrap(), 12);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("").unwrap(), "a2582a3a0e66e6e86e3812dcb672a272");
        assert_eq!(
            part2("AoC 2017").unwrap(),
            "33efeb34ea91902bb2f59c9920caa6cd"
        );
        assert_eq!(part2("1,2,3").unwrap(), "3efbe78a8d82f29979031a4aa0b16a9d");
        assert_eq!(part2("1,2,4").unwrap(), "63960835bcdc130f0b66d7ff4f6a5a8e");
    }
}
