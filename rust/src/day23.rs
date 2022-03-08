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

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    // Find instructions that I think can be different in various inputs
    // Assumption: all other instructions are the same, and in the same place

    fn extract_num(line: &str, prefix: &str) -> Result<i32, std::num::ParseIntError> {
        return line.strip_prefix(prefix).unwrap().parse::<i32>();
    }

    let mut first_sub_b = true;
    let mut b: i32 = 0;
    let mut c: i32 = 0;
    let mut bd: i32 = 0;

    for line in s.lines() {
        if line.starts_with("set b") {
            b = extract_num(line, "set b ")?;
        } else if line.starts_with("mul b") {
            b *= extract_num(line, "mul b ")?;
        } else if line.starts_with("sub c") {
            c = b - extract_num(line, "sub c ")?;
        } else if line.starts_with("sub b") {
            if first_sub_b {
                b -= extract_num(line, "sub b ")?;
                first_sub_b = false;
            } else {
                bd = -extract_num(line, "sub b ")?;
            }
        }
    }
    let mut h = 0; // number of occurrences found
    let mut d: i32; // factor one
    let mut e: i32; // factor two
    let mut f: i32; // has a product been found

    loop {
        f = 1;
        d = 2;
        loop {
            e = 2;
            loop {
                let p = d * e;
                if p > b {
                    // d*e will never == b any more
                    break;
                }
                if p == b {
                    f = 0;
                    // break because it will never reset to 1
                    break;
                }
                e += 1;
                if e == b {
                    break;
                }
            }
            d += 1;
            if d == b {
                break;
            }
            if f == 0 {
                // break because it will never reset to 1
                break;
            }
        }
        if f == 0 {
            h += 1;
        }
        if (b - c) == 0 {
            break;
        }
        b += bd;
    }

    return Ok(format!("{}", h));
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
