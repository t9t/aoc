use std::collections::HashSet;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let banks = parse(s)?;
    return Ok(format!("{}", reallocation_cycles(banks).0));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let mut banks = parse(s)?;
    banks = reallocation_cycles(banks).1;
    return Ok(format!("{}", reallocation_cycles(banks).0));
}

fn parse(s: &str) -> Result<Vec<u32>, std::num::ParseIntError> {
    return s
        .split("\t")
        .map(|x| x.parse())
        .collect::<Result<Vec<u32>, _>>();
}

fn reallocation_cycles(mut banks: Vec<u32>) -> (i32, Vec<u32>) {
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

    return (cycle, banks);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("0\t2\t7\t0").unwrap(), "5");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("0\t2\t7\t0").unwrap(), "4");
    }
}
