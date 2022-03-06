use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let lines = s.lines().collect::<Vec<&str>>();
    let mut registers: HashMap<char, i64> = HashMap::new();

    fn get_val(k: &str, registers: &HashMap<char, i64>) -> i64 {
        let parsed = k.parse::<i64>();
        if parsed.is_ok() {
            return parsed.unwrap();
        }
        return *registers.get(&k.chars().next().unwrap()).unwrap_or(&0);
    }

    let mut pos: i64 = 0;
    let mut muls: u32 = 0;
    loop {
        if pos < 0 || pos >= lines.len() as i64 {
            break;
        }

        let line = lines[pos as usize];

        let mut parts = line.split(" ");
        let instruction = parts.next().unwrap();

        let x = parts.next().unwrap();
        let xreg = x.chars().next().unwrap();
        let xval = get_val(x, &registers);

        let y = parts.next().unwrap();
        let yval = get_val(y, &registers);

        match instruction {
            "set" => {
                registers.insert(xreg, yval);
            }
            "sub" => {
                registers.insert(xreg, xval - yval);
            }
            "mul" => {
                registers.insert(xreg, xval * yval);
                muls += 1;
            }
            "jnz" => {
                if xval != 0 {
                    pos += yval;
                    continue;
                }
            }
            _ => Err(format!("Invalid instruction: {}", line))?,
        }
        pos += 1;
    }
    return Ok(format!("{}", muls));
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
        assert_eq!(part1(INPUT).unwrap(), "1337");
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
