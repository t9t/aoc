use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let map = parse_programs(s)?;

    for (lk, (_, lv)) in &map {
        if !lv.is_empty() {
            if !map.values().any(|(_, rv)| rv.contains(lk)) {
                return Ok(format!("{}", lk));
            }
        }
    }

    return Err("No bottom program found")?;
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let programs = parse_programs(s)?;

    fn get_weight(name: &str, programs: &HashMap<&str, (u32, Vec<&str>)>) -> u32 {
        let (weight, subprograms) = programs.get(name).unwrap();
        let mut total_weight = *weight;
        for subprogram in subprograms {
            total_weight += get_weight(subprogram, programs);
        }
        return total_weight;
    }

    let mut min_discrepancy = u32::MAX;
    for (_, (_, subprograms)) in &programs {
        if subprograms.len() < 3 {
            // I've no idea how to find out which the "right" weight is when there are only 2 subprograms
            continue;
        }

        let mut weight_map = HashMap::new();
        for subprogram in subprograms {
            weight_map
                .entry(get_weight(subprogram, &programs))
                .or_insert_with(Vec::new)
                .push(subprogram);
        }

        if weight_map.len() == 1 {
            continue;
        }

        let common_weight = weight_map.iter().max_by(|l, r| l.1.len().cmp(&r.1.len()));
        let mut target = *common_weight.unwrap().0;
        let discrepancy = weight_map.iter().min_by(|l, r| l.1.len().cmp(&r.1.len()));
        let (_, subsubprograms) = programs.get(discrepancy.unwrap().1[0]).unwrap();
        for subsubprogram in subsubprograms {
            let zw = get_weight(subsubprogram, &programs);
            target -= zw;
        }
        min_discrepancy = min_discrepancy.min(target);
    }

    return Ok(format!("{}", min_discrepancy));
}

fn parse_programs(s: &str) -> Result<HashMap<&str, (u32, Vec<&str>)>, Box<dyn Error>> {
    let mut programs: HashMap<&str, (u32, Vec<&str>)> = HashMap::new();
    for line in s.lines() {
        let ab = line.split(" -> ").collect::<Vec<&str>>();
        let subprograms: Vec<&str> = if ab.len() == 2 {
            ab[1].split(", ").collect()
        } else {
            Vec::new()
        };
        let weight_str = ab[0].split("(").nth(1).unwrap().replace(")", "");
        let name = ab[0].split(" ").collect::<Vec<&str>>()[0];
        programs.insert(name, (weight_str.parse::<u32>()?, subprograms));
    }

    return Ok(programs);
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "pbga (66)
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

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "tknk");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "60");
    }
}
