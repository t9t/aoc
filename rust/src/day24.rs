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

    let all = build_all(0, &Vec::new(), &components);
    let mut max_len = 0;
    let mut max_str = 0;
    for bridge in all {
        let len = bridge.len();
        let strength = bridge.iter().map(|(l, r)| l + r).sum();

        if len > max_len {
            max_len = len;
            max_str = strength;
        } else if len == max_len && strength > max_str {
            max_str = strength;
        }
    }

    return Ok(format!("{}", max_str));
}

fn build_all(
    port: u32,
    so_far: &Vec<(u32, u32)>,
    others: &Vec<(u32, u32)>,
) -> Vec<Vec<(u32, u32)>> {
    let mut bridges: Vec<Vec<(u32, u32)>> = Vec::new();

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

        let mut next_bridge = so_far.clone();
        next_bridge.push((other_left, other_right));
        let all_next = build_all(next_port, &next_bridge, &next_others);
        bridges.push(next_bridge);
        for n in all_next {
            bridges.push(n);
        }
    }
    return bridges;
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
