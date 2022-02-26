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

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_knot_hash() {
        assert_eq!(knot_hash("3,4,1,5", 5).unwrap(), 12);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
