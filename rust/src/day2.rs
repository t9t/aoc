pub fn part1(s: &str) -> i32 {
    let mut sum = 0;
    for line in s.lines() {
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

pub fn part2(s: &str) -> i32 {
    let mut sum = 0;
    for line in s.lines() {
        let nums: Vec<i32> = line
            .split("\t")
            .map(|x| x.parse::<i32>().unwrap())
            .collect();
        for (i, n) in nums.iter().enumerate() {
            let left = *n;
            for j in i + 1..nums.len() {
                let right = nums[j];
                let h = if left > right { left } else { right };
                let l = if left > right { right } else { left };
                if h % l == 0 {
                    sum += h / l;
                }
            }
        }
    }
    return sum;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8";
        assert_eq!(part1(input), 18);
    }

    #[test]
    fn test_part2() {
        let input = "5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5";
        assert_eq!(part2(input), 9);
    }
}
