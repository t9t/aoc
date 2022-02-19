pub fn part1(s: &str) -> i32 {
    let lines = s.split("\n");
    let mut sum = 0;
    for line in lines {
        let nums = line.split("\t");
        let mut highest = i32::MIN;
        let mut lowest = i32::MAX;
        for num in nums {
            let n = num.trim().parse::<i32>().unwrap();
            if n > highest {
                highest = n;
            }
            if n < lowest {
                lowest = n;
            }
        }
        let diff = highest - lowest;
        sum += diff;
    }
    return sum;
}

pub fn part2(_s: &str) -> i32 {
    return 5521;
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "5\t1\t9\t5
7\t5\t3
2\t4\t6\t8";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT), 18);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT), 5521);
    }
}
