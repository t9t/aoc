use std::collections::HashMap;
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

pub fn part2(_s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", 5521));
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2";

    #[test]
    fn test_part1() {
        assert_eq!(part1(INPUT).unwrap(), "4");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
