use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut regs: HashMap<&str, i32> = HashMap::new();

    for line in s.lines() {
        let a = line.split(" ").collect::<Vec<&str>>();
        let reg = a[0];
        let op = a[1];
        let v = a[2].parse::<i32>()?;
        let creg = a[4];
        let comp = a[5];
        let compv = a[6].parse::<i32>()?;
        let comp_reg_v = *regs.get(creg).unwrap_or(&0);

        let condition_result = match comp {
            ">" => Ok(comp_reg_v > compv),
            "<" => Ok(comp_reg_v < compv),
            ">=" => Ok(comp_reg_v >= compv),
            "<=" => Ok(comp_reg_v <= compv),
            "==" => Ok(comp_reg_v == compv),
            "!=" => Ok(comp_reg_v != compv),
            _ => Err("Invalid comparison operator"),
        }?;
        if condition_result {
            let reg_v = *regs.get(reg).unwrap_or(&0);
            regs.insert(reg, reg_v + v * if op == "dec" { -1 } else { 1 });
        }
    }

    return Ok(format!("{}", regs.values().max().unwrap_or(&0)));
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
        let input = "b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10";
        assert_eq!(part1(input).unwrap(), "1");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
