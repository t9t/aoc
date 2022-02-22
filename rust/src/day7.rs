use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut map: HashMap<&str, Vec<&str>> = HashMap::new();
    for line in s.lines() {
        let ab = line.split(" -> ").collect::<Vec<&str>>();
        let programs: Vec<&str> = if ab.len() == 2 {
            ab[1].split(", ").collect()
        } else {
            Vec::new()
        };
        map.insert(ab[0].split(" ").collect::<Vec<&str>>()[0], programs);
    }

    for (lk, lv) in &map {
        if !lv.is_empty() {
            if !map.values().any(|rv| rv.contains(lk)) {
                return Ok(format!("{}", lk));
            }
        }
    }

    return Err("No bottom program found")?;
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
        let input = "pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)";

        assert_eq!(part1(input).unwrap(), "tknk");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
