use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let (max_sum, _, _) = parse_and_find(s)?;
    return Ok(format!("{}", max_sum));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let (_, _, max_str) = parse_and_find(s)?;
    return Ok(format!("{}", max_str));
}

fn parse_and_find(s: &str) -> Result<(u32, u32, u32), Box<dyn Error>> {
    let mut components = Vec::new();

    for line in s.lines() {
        let mut split = line.split("/");
        let left = split.next().unwrap().parse::<u32>()?;
        let right = split.next().unwrap().parse::<u32>()?;
        components.push((left, right));
    }
    return Ok(find_maxes(0, 0, 0, &components));
}

fn find_maxes(port: u32, len: u32, strength: u32, others: &Vec<(u32, u32)>) -> (u32, u32, u32) {
    let mut max_sum = 0;
    let mut max_len = len;
    let mut max_str: u32 = strength;

    for (i, &(other_left, other_right)) in others.iter().enumerate() {
        if other_left != port && other_right != port {
            continue;
        }

        let mut next_others = others.clone();
        next_others.remove(i);
        let next_port = if other_left == port {
            other_right
        } else {
            other_left
        };

        let next_strength = strength + other_left + other_right;
        let (sum, next_max_len, next_max_str) =
            find_maxes(next_port, len + 1, next_strength, &next_others);
        max_sum = max_sum.max(sum + other_left + other_right);
        if next_max_len > max_len {
            max_len = next_max_len;
            max_str = next_max_str;
        } else if next_max_len == max_len && next_max_str > max_str {
            max_str = next_max_str;
        }
    }
    return (max_sum, max_len, max_str);
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "31");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "19");
    }
}
