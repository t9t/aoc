use std::collections::HashMap;
use std::collections::VecDeque;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut regs: HashMap<char, i64> = HashMap::new();

    let lines = s.lines().collect::<Vec<&str>>();
    let mut pos: i64 = 0;
    let mut last_frequency = 0;
    loop {
        if pos < 0 || pos >= lines.len() as i64 {
            Err("Jumped outside of range")?
        }
        let line = lines[pos as usize];
        let mut split = line.split(" ");
        let op = split.next().unwrap();
        let reg = split.next().unwrap().chars().next().unwrap();
        let x = *regs.get(&reg).unwrap_or(&0);
        let values = split.next().unwrap_or("0");
        let value = values.parse::<i64>();
        let val = if value.is_ok() {
            value?
        } else {
            *regs.get(&values.chars().next().unwrap()).unwrap_or(&0)
        };

        match op {
            "snd" => {
                last_frequency = x;
            }
            "set" => {
                regs.insert(reg, val);
            }
            "add" => {
                regs.insert(reg, x + val);
            }
            "mul" => {
                regs.insert(reg, x * val);
            }
            "mod" => {
                regs.insert(reg, x % val);
            }
            "rcv" => {
                if x != 0 {
                    return Ok(format!("{}", last_frequency));
                }
            }
            "jgz" => {
                if x > 0 {
                    pos += val;
                    continue;
                }
            }
            _ => Err(format!("Invalid instruction: {}", line))?,
        }
        pos += 1;
    }
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let lines = s.lines().collect::<Vec<&str>>();
    let mut regs0: HashMap<char, i64> = HashMap::new();
    regs0.insert('p', 0);
    let mut regs1: HashMap<char, i64> = HashMap::new();
    regs1.insert('p', 1);

    let mut from0to1: VecDeque<i64> = VecDeque::new();
    let mut from1to0: VecDeque<i64> = VecDeque::new();

    let mut pos0: i64 = 0;
    let mut pos1: i64 = 0;

    fn get_val(regs: &HashMap<char, i64>, reg_or_value: &str) -> i64 {
        let as_int = reg_or_value.parse::<i64>();
        if as_int.is_ok() {
            return as_int.unwrap();
        }
        let c = reg_or_value.chars().next().unwrap();
        let v = regs.get(&c);
        let i = v.unwrap_or(&0);
        return *i;
    }

    let mut one_sent = 0;

    loop {
        // Program 0
        loop {
            if pos0 < 0 || pos0 >= lines.len() as i64 {
                Err("Jumped outside of range")?
            }
            let line = lines[pos0 as usize];
            let mut split = line.split(" ");
            let op = split.next().unwrap();
            let reg_or_val = split.next().unwrap();
            let reg = reg_or_val.chars().next().unwrap();
            let x = get_val(&regs0, reg_or_val);
            let val = get_val(&regs0, split.next().unwrap_or("0"));

            match op {
                "snd" => {
                    from0to1.push_back(x);
                }
                "set" => {
                    regs0.insert(reg, val);
                }
                "add" => {
                    regs0.insert(reg, x + val);
                }
                "mul" => {
                    regs0.insert(reg, x * val);
                }
                "mod" => {
                    regs0.insert(reg, x % val);
                }
                "rcv" => {
                    if from1to0.is_empty() {
                        break;
                    }
                    let rcv = from1to0.pop_front().unwrap();
                    regs0.insert(reg, rcv);
                }
                "jgz" => {
                    if x > 0 {
                        pos0 += val;
                        continue;
                    }
                }
                _ => Err(format!("Invalid instruction: {}", line))?,
            }
            pos0 += 1;
        }

        // Program 1
        loop {
            if pos1 < 0 || pos1 >= lines.len() as i64 {
                Err("Jumped outside of range")?
            }
            let line = lines[pos1 as usize];
            let mut split = line.split(" ");
            let op = split.next().unwrap();
            let reg_or_val = split.next().unwrap();
            let reg = reg_or_val.chars().next().unwrap();
            let x = get_val(&regs1, reg_or_val);
            let val = get_val(&regs1, split.next().unwrap_or("0"));

            match op {
                "snd" => {
                    from1to0.push_back(x);
                    one_sent += 1;
                }
                "set" => {
                    regs1.insert(reg, val);
                }
                "add" => {
                    regs1.insert(reg, x + val);
                }
                "mul" => {
                    regs1.insert(reg, x * val);
                }
                "mod" => {
                    regs1.insert(reg, x % val);
                }
                "rcv" => {
                    if from0to1.is_empty() {
                        break;
                    }
                    let rcv = from0to1.pop_front().unwrap();
                    regs1.insert(reg, rcv);
                }
                "jgz" => {
                    if x > 0 {
                        pos1 += val;
                        continue;
                    }
                }
                _ => Err(format!("Invalid instruction: {}", line))?,
            }
            pos1 += 1;
        }

        if from0to1.is_empty() && from1to0.is_empty() {
            break;
        }
    }
    return Ok(format!("{}", one_sent));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2";
        assert_eq!(part1(input).unwrap(), "4");
    }
}
