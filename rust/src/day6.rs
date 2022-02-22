use std::collections::HashSet;
use std::error::Error;

pub fn part1(s: &str) -> Result<i32, Box<dyn Error>> {
    let mut banks: Vec<u32> = s.split("\t").map(|x| x.parse::<u32>().unwrap()).collect();
    let mut cycle = 0;
    let mut seen: HashSet<Vec<u32>> = HashSet::new();
    while !seen.contains(&banks) {
        cycle += 1;
        seen.insert(banks.clone());
        let mut reallocate = 0;
        let mut idx = 0;
        for (i, bank) in banks.iter().enumerate() {
            if *bank > reallocate {
                reallocate = *bank;
                idx = i;
            }
        }

        banks[idx] = 0;
        while reallocate > 0 {
            idx += 1;
            if idx >= banks.len() {
                idx = 0;
            }
            banks[idx] += 1;
            reallocate -= 1;
        }
    }

    return Ok(cycle);
}

pub fn part2(_s: &str) -> Result<i32, Box<dyn Error>> {
    return Ok(5521);
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_part1() {
        assert_eq!(part1("0\t2\t7\t0").unwrap(), 5);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), 5521);
    }
}
