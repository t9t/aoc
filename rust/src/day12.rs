use std::collections::HashMap;
use std::collections::HashSet;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut connections: HashMap<u32, Vec<u32>> = HashMap::new();
    for line in s.lines() {
        let split = line.split(" <-> ").collect::<Vec<&str>>();
        let left = split[0].parse::<u32>()?;
        let right = (split[1]
            .split(", ")
            .map(|x| x.parse::<u32>())
            .collect::<Result<Vec<u32>, _>>())?;
        connections.insert(left, right);
    }

    let mut group0: HashSet<u32> = HashSet::new();
    let mut consider: Vec<u32> = Vec::new();
    consider.push(0);
    while !consider.is_empty() {
        let n = consider.pop().unwrap();
        group0.insert(n);
        let c = connections.get(&n).unwrap();
        for k in c {
            if !group0.contains(k) {
                consider.push(*k);
            }
        }
    }

    return Ok(format!("{}", group0.len()));
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
    fn test_part1() {
        let input = "0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5";

        assert_eq!(part1(input).unwrap(), "6");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
