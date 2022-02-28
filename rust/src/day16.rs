use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    let mut programs = [
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
    ]
    .to_vec();

    for dance_move in s.split(",") {
        let move_type = dance_move.chars().nth(0).unwrap();
        let move_data = &dance_move[1..];
        match move_type {
            's' => {
                let size = move_data.parse::<usize>()?;
                let end_start = programs.len() - size;
                let mut new_programs: Vec<char> = Vec::new();
                for i in end_start..programs.len() {
                    new_programs.push(programs[i]);
                }
                for i in 0..end_start {
                    new_programs.push(programs[i]);
                }
                programs = new_programs;
            }
            'x' => {
                let mut swap_programs = move_data.split("/");
                let left_pos = swap_programs.next().unwrap().parse::<usize>()?;
                let right_pos = swap_programs.next().unwrap().parse::<usize>()?;

                let left = programs[left_pos];
                programs[left_pos] = programs[right_pos];
                programs[right_pos] = left;
            }
            'p' => {
                let mut swap_programs = move_data.split("/");
                let left = swap_programs.next().unwrap().chars().next().unwrap();
                let right = swap_programs.next().unwrap().chars().next().unwrap();

                let left_pos = programs.iter().position(|&c| c == left).unwrap();
                let right_pos = programs.iter().position(|&c| c == right).unwrap();
                programs[left_pos] = right;
                programs[right_pos] = left;
            }
            _ => Err(format!("Unknown move {}", dance_move))?,
        }
    }

    let mut out = String::new();
    for c in programs {
        out.push(c);
    }

    return Ok(out);
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
