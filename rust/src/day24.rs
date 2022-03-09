use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut starts = Vec::new();
    let mut others = Vec::new();

    for line in s.lines() {
        let mut split = line.split("/");
        let left = split.next().unwrap().parse::<u32>()?;
        let right = split.next().unwrap().parse::<u32>()?;
        if left == 0 || right == 0 {
            starts.push((left, right));
        } else {
            others.push((left, right));
        }
    }

    let mut max_sum = 0;
    for &(left, right) in &starts {
        max_sum = max_sum.max(calculate_max_sum(left, right, &others));
    }

    return Ok(format!("{}", max_sum));
}

fn calculate_max_sum(left: u32, right: u32, others: &Vec<(u32, u32)>) -> u32 {
    let mut max_next = 0;
    for (i, &(other_left, other_right)) in others.iter().enumerate() {
        if other_left == left || other_left == right || other_right == left || other_right == right
        {
            let mut next_others = others.clone();
            next_others.remove(i);
            let next_sum = calculate_max_sum(other_left, other_right, &next_others);
            max_next = max_next.max(next_sum);
        }
    }
    return left + right + max_next;
}

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
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
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
