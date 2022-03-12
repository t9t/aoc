use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut components = Vec::new();

    for line in s.lines() {
        let mut split = line.split("/");
        let left = split.next().unwrap().parse::<u32>()?;
        let right = split.next().unwrap().parse::<u32>()?;
        components.push((left, right));
    }

    let max_sum = find_max_sum(0, &components);

    return Ok(format!("{}", max_sum));
}

fn find_max_sum(port: u32, others: &Vec<(u32, u32)>) -> u32 {
    let mut max_sum = 0;
    for (i, &(other_left, other_right)) in others.iter().enumerate() {
        if other_left == port || other_right == port {
            let mut next_others = others.clone();
            next_others.remove(i);
            let next_port = if other_left == port {
                other_right
            } else {
                other_left
            };
            let sum = find_max_sum(next_port, &next_others) + other_left + other_right;
            max_sum = max_sum.max(sum);
        }
    }
    return max_sum;
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let mut components = Vec::new();

    for line in s.lines() {
        let mut split = line.split("/");
        let left = split.next().unwrap().parse::<u32>()?;
        let right = split.next().unwrap().parse::<u32>()?;
        components.push((left, right));
    }

    let (_, max_str) = max_len_strength(0, 0, 0, &components);
    return Ok(format!("{}", max_str));
}

fn max_len_strength(port: u32, len: u32, strength: u32, others: &Vec<(u32, u32)>) -> (u32, u32) {
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
        let (ml, ms) = max_len_strength(next_port, len + 1, next_strength, &next_others);
        if ml > max_len {
            max_len = ml;
            max_str = ms;
        } else if ml == max_len && ms > max_str {
            max_str = ms;
        }
    }
    return (max_len, max_str);
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
