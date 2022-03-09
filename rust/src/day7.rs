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

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let mut weights: HashMap<&str, u32> = HashMap::new();
    let mut map: HashMap<&str, Vec<&str>> = HashMap::new();
    for line in s.lines() {
        let ab = line.split(" -> ").collect::<Vec<&str>>();
        let programs: Vec<&str> = if ab.len() == 2 {
            ab[1].split(", ").collect()
        } else {
            Vec::new()
        };
        let name = ab[0].split(" ").next().unwrap();
        let weight = ab[0].split("(").nth(1).unwrap().replace(")", "");
        map.insert(name, programs);
        weights.insert(name, weight.parse::<u32>()?);
    }

    let mut cache: HashMap<String, u32> = HashMap::new();
    fn get_weight(
        name: &str,
        programs: &HashMap<&str, Vec<&str>>,
        weights: &HashMap<&str, u32>,
        cache: &mut HashMap<String, u32>,
    ) -> u32 {
        let cached = cache.get(name);
        if cached.is_some() {
            return *cached.unwrap();
        }
        let mut weight = *weights.get(name).unwrap();
        let subprograms = programs.get(name).unwrap();
        for subprogram in subprograms {
            weight += get_weight(subprogram, programs, weights, cache);
        }

        cache.insert(String::from(name), weight);
        return weight;
    }

    let mut min_discrep = u32::MAX;
    for (_, subprograms) in &map {
        if subprograms.len() < 3 {
            // I've no idea how to find out which the "right" weight is when there are only 2 subprograms
            continue;
        }
        let mut weight_map: HashMap<u32, Vec<&str>> = HashMap::new();
        for subprogram in subprograms {
            let sw = get_weight(subprogram, &map, &weights, &mut cache);
            let c = weight_map.get_mut(&sw);
            if c.is_some() {
                c.unwrap().push(subprogram);
            } else {
                let mut x: Vec<&str> = Vec::new();
                x.push(subprogram);
                weight_map.insert(sw, x);
            }
        }
        if weight_map.len() != 1 {
            let common = weight_map.iter().max_by(|l, r| l.1.len().cmp(&r.1.len()));
            let discrep = weight_map.iter().min_by(|l, r| l.1.len().cmp(&r.1.len()));
            let mut yay = *common.unwrap().0;
            let ohno = discrep.unwrap().1[0];
            let kek = map.get(ohno).unwrap();
            for bur in kek {
                let zw = get_weight(bur, &map, &weights, &mut cache);
                yay -= zw;
            }
            min_discrep = min_discrep.min(yay);
        }
    }

    return Ok(format!("{}", min_discrep));
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
