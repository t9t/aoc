use std::collections::HashMap;
use std::collections::VecDeque;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let lines = s.lines().collect::<Vec<&str>>();
    let mut program = Program {
        pos: 0,
        regs: HashMap::new(),
        inbox: VecDeque::new(),
    };
    let mut outbox: VecDeque<i64> = VecDeque::new();
    program.run(&mut outbox, &lines)?;
    return Ok(format!("{}", outbox.pop_back().ok_or("No value sent")?));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let lines = s.lines().collect::<Vec<&str>>();

    let mut program0 = Program {
        pos: 0,
        regs: HashMap::new(),
        inbox: VecDeque::new(),
    };
    program0.regs.insert('p', 0);

    let mut program1 = Program {
        pos: 0,
        regs: HashMap::new(),
        inbox: VecDeque::new(),
    };
    program1.regs.insert('p', 1);

    let mut one_sent = 0;
    loop {
        let _ = program0.run(&mut program1.inbox, &lines)?;
        one_sent += program1.run(&mut program0.inbox, &lines)?;

        if program0.inbox.is_empty() && program1.inbox.is_empty() {
            return Ok(format!("{}", one_sent));
        }
    }
}

struct Program {
    pos: i64,
    regs: HashMap<char, i64>,
    inbox: VecDeque<i64>,
}

impl Program {
    fn run(
        &mut self,
        outbox: &mut VecDeque<i64>,
        lines: &Vec<&str>,
    ) -> Result<u32, Box<dyn Error>> {
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
        let mut sent = 0;
        loop {
            if self.pos < 0 || self.pos >= lines.len() as i64 {
                Err("Jumped outside of range")?
            }
            let line = lines[self.pos as usize];
            let mut split = line.split(" ");
            let op = split.next().unwrap();
            let reg_or_val = split.next().unwrap();
            let reg = reg_or_val.chars().next().unwrap();
            let x = get_val(&self.regs, reg_or_val);
            let val = get_val(&self.regs, split.next().unwrap_or("0"));

            match op {
                "snd" => {
                    outbox.push_back(x);
                    sent += 1;
                }
                "set" => {
                    self.regs.insert(reg, val);
                }
                "add" => {
                    self.regs.insert(reg, x + val);
                }
                "mul" => {
                    self.regs.insert(reg, x * val);
                }
                "mod" => {
                    self.regs.insert(reg, x % val);
                }
                "rcv" => {
                    if self.inbox.is_empty() {
                        break;
                    }
                    let rcv = self.inbox.pop_front().unwrap();
                    self.regs.insert(reg, rcv);
                }
                "jgz" => {
                    if x > 0 {
                        self.pos += val;
                        continue;
                    }
                }
                _ => Err(format!("Invalid instruction: {}", line))?,
            }
            self.pos += 1;
        }
        return Ok(sent);
    }
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
