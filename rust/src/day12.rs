use std::collections::HashMap;
use std::collections::HashSet;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", group(s)?.get(&0).unwrap().len()));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", group(s)?.len()));
}

fn group(s: &str) -> Result<HashMap<u32, HashSet<u32>>, Box<dyn Error>> {
    let mut connections: HashMap<u32, Vec<u32>> = HashMap::new();
    let mut min_id = u32::MAX;
    let mut max_id = 0;
    for line in s.lines() {
        let split = line.split(" <-> ").collect::<Vec<&str>>();
        let left = split[0].parse::<u32>()?;
        let right = (split[1]
            .split(", ")
            .map(|x| x.parse::<u32>())
            .collect::<Result<Vec<u32>, _>>())?;
        connections.insert(left, right);
        min_id = min_id.min(left);
        max_id = max_id.max(left);
    }

    let mut grouped: HashSet<u32> = HashSet::new();
    let mut groups: HashMap<u32, HashSet<u32>> = HashMap::new();

    for f in min_id..max_id + 1 {
        if grouped.contains(&f) {
            continue;
        }
        grouped.insert(f);

        let mut group: HashSet<u32> = HashSet::new();
        let mut consider: Vec<u32> = Vec::new();

        consider.push(f);
        while !consider.is_empty() {
            let n = consider.pop().unwrap();
            group.insert(n);
            grouped.insert(n);
            let c = connections.get(&n).unwrap();
            for k in c {
                if !group.contains(k) {
                    consider.push(*k);
                }
            }
        }
        if !group.is_empty() {
            groups.insert(f, group);
        }
    }

    return Ok(groups);
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "6");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "2");
    }
}
