pub fn part1(s: &str) -> i32 {
    return calc_sum(s, 1);
}

pub fn part2(s: &str) -> i32 {
    let l = s.len();
    return calc_sum(s, l / 2);
}

fn calc_sum(s: &str, x: usize) -> i32 {
    let chars: Vec<char> = s.chars().collect();
    let len = chars.len();
    let mut sum: i32 = 0;
    for (i, c) in chars.iter().enumerate() {
        let next_pos = (i + x) % len;
        let next = chars[next_pos];
        if *c == next {
            sum += c.to_digit(10).unwrap() as i32
        }
    }
    return sum;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("1122"), 3);
        assert_eq!(part1("1111"), 4);
        assert_eq!(part1("1234"), 0);
        assert_eq!(part1("91212129"), 9);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2("1212"), 6);
        assert_eq!(part2("1221"), 0);
        assert_eq!(part2("123425"), 4);
        assert_eq!(part2("123123"), 12);
        assert_eq!(part2("12131415"), 4);
    }
}
