use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let instructions = parse_instructions(s)?;

    let line0 = s.lines().next().unwrap();
    let mut state = line0.chars().nth(line0.len() - 2).unwrap();

    let line1 = s.lines().nth(1).unwrap();
    let steps = line1.split(" ").nth(5).unwrap().parse::<i32>()?;

    let mut tape: HashMap<i32, u8> = HashMap::new();
    let mut pos: i32 = 0;

    for _ in 0..steps {
        let curr = *tape.get(&pos).unwrap_or(&0);
        let ins = instructions.get(&state).unwrap().get(&curr).unwrap();
        tape.insert(pos, ins.write);
        pos += ins.delta;
        state = ins.next;
    }

    return Ok(format!("{}", tape.values().filter(|&v| *v == 1).count()));
}

#[derive(Debug)]
struct Instruction {
    write: u8,
    delta: i32,
    next: char,
}

fn parse_instructions(s: &str) -> Result<HashMap<char, HashMap<u8, Instruction>>, Box<dyn Error>> {
    let mut instructions: HashMap<char, HashMap<u8, Instruction>> = HashMap::new();

    for mut desc in s.split("In state ").skip(1) {
        desc = desc.trim();
        let mut state_instructions: HashMap<u8, Instruction> = HashMap::new();

        for curval in desc.split("If the current value is ").skip(1) {
            let mut instruction = Instruction {
                write: 0,
                delta: 0,
                next: 'A',
            };
            for mut line in curval.split("- ").skip(1) {
                line = line.trim();
                line = &line[0..line.len() - 1];
                if line.starts_with("Write the value ") {
                    instruction.write = line
                        .strip_prefix("Write the value ")
                        .unwrap()
                        .parse::<u8>()?;
                } else if line.starts_with("Move one slot to the ") {
                    let dir = line.strip_prefix("Move one slot to the ").unwrap();
                    instruction.delta = if dir == "right" { 1 } else { -1 };
                } else if line.starts_with("Continue with state ") {
                    instruction.next = line
                        .strip_prefix("Continue with state ")
                        .unwrap()
                        .chars()
                        .next()
                        .unwrap();
                } else {
                    return Err(format!("Invalid line: {}", line))?;
                }
            }
            state_instructions.insert(curval[0..1].parse::<u8>()?, instruction);
        }

        instructions.insert(desc.chars().next().unwrap(), state_instructions);
    }

    return Ok(instructions);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "Begin in state A.
Perform a diagnostic checksum after 6 steps.
    
In state A:
If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.";
        assert_eq!(part1(input).unwrap(), "3");
    }
}
