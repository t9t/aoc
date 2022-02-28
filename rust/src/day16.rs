use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(to_str(dance(start_programs(), &parse_moves(s)?)));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let moves = parse_moves(s)?;
    // Find repetition cycle
    let start_state = start_programs();
    let mut programs = start_programs();
    let mut i = 0;
    loop {
        programs = dance(programs, &moves);
        i += 1;
        if programs == start_state {
            break;
        }
    }

    // Run for one cycle (the "final" cycle)
    for _ in 0..1_000_000_000 % i {
        programs = dance(programs, &moves);
    }

    return Ok(to_str(programs));
}

#[derive(Debug, PartialEq)]
enum Move {
    Spin(usize),
    Exchange(usize, usize),
    Partner(char, char),
}

fn start_programs() -> Vec<char> {
    return vec![
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
    ];
}

fn parse_moves(s: &str) -> Result<Vec<Move>, Box<dyn Error>> {
    let mut moves: Vec<Move> = Vec::new();

    for dance_move in s.split(",") {
        let move_type = dance_move.chars().nth(0).unwrap();
        let move_data = &dance_move[1..];
        match move_type {
            's' => moves.push(Move::Spin(move_data.parse::<usize>()?)),
            'x' => {
                let mut swap_programs = move_data.split("/");
                let left_pos = swap_programs.next().unwrap().parse::<usize>()?;
                let right_pos = swap_programs.next().unwrap().parse::<usize>()?;
                moves.push(Move::Exchange(left_pos, right_pos))
            }
            'p' => {
                let mut swap_programs = move_data.split("/");
                let left = swap_programs.next().unwrap().chars().next().unwrap();
                let right = swap_programs.next().unwrap().chars().next().unwrap();
                moves.push(Move::Partner(left, right));
            }
            _ => Err(format!("Unknown move {}", dance_move))?,
        }
    }

    return Ok(moves);
}

fn dance(mut programs: Vec<char>, moves: &Vec<Move>) -> Vec<char> {
    for dance_move in moves {
        match *dance_move {
            Move::Spin(size) => {
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
            Move::Exchange(left_pos, right_pos) => {
                let left = programs[left_pos];
                programs[left_pos] = programs[right_pos];
                programs[right_pos] = left;
            }
            Move::Partner(left, right) => {
                let left_pos = programs.iter().position(|&c| c == left).unwrap();
                let right_pos = programs.iter().position(|&c| c == right).unwrap();
                programs[left_pos] = right;
                programs[right_pos] = left;
            }
        }
    }

    return programs;
}

fn to_str(programs: Vec<char>) -> String {
    let mut out = String::new();
    for c in programs {
        out.push(c);
    }
    return out;
}

#[cfg(test)]
mod tests {
    use super::*;

    static INPUT: &str = "bla
bla";

    #[test]
    fn test_parse_moves() {
        assert_eq!(
            parse_moves("s1,x3/4,pe/b").unwrap(),
            vec![Move::Spin(1), Move::Exchange(3, 4), Move::Partner('e', 'b')]
        );
    }

    #[test]
    fn test_dance() {
        assert_eq!(
            dance(vec!['a', 'b', 'c', 'd', 'e'], &vec![Move::Spin(3)]),
            vec!['c', 'd', 'e', 'a', 'b']
        );
        assert_eq!(
            dance(vec!['a', 'b', 'c', 'd', 'e'], &vec![Move::Spin(1)]),
            vec!['e', 'a', 'b', 'c', 'd']
        );
        assert_eq!(
            dance(vec!['e', 'a', 'b', 'c', 'd'], &vec![Move::Exchange(3, 4)]),
            vec!['e', 'a', 'b', 'd', 'c']
        );
        assert_eq!(
            dance(
                vec!['e', 'a', 'b', 'd', 'c'],
                &vec![Move::Partner('e', 'b')]
            ),
            vec!['b', 'a', 'e', 'd', 'c']
        );
    }

    #[test]
    fn test_dance_multi() {
        let moves = vec![Move::Spin(1), Move::Exchange(3, 4), Move::Partner('e', 'b')];
        let programs = vec!['a', 'b', 'c', 'd', 'e'];
        assert_eq!(dance(programs, &moves), vec!['b', 'a', 'e', 'd', 'c']);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(INPUT).unwrap(), "5521");
    }
}
