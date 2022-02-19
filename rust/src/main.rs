use std::fs;
use std::io;

fn main() {
    let s = read_input(2017, 1).unwrap();
    println!("sum: {}", calc_sum(&s.trim()))
}

fn read_input(year: u16, day: u8) -> io::Result<String> {
    return fs::read_to_string(format!("../input/{}/{}.txt", year, day));
}

fn calc_sum(s: &str) -> u32 {
    return calc_sum_x(s, 1);
}

fn calc_sum2(s: &str) -> u32 {
    let l = s.len();
    return calc_sum_x(s, l / 2);
}

fn calc_sum_x(s: &str, x: usize) -> u32 {
    let chars: Vec<char> = s.chars().collect();
    let len = chars.len();
    let mut sum = 0;
    for (i, c) in chars.iter().enumerate() {
        let next_pos = (i + x) % len;
        let next = chars[next_pos];
        if *c == next {
            sum += c.to_digit(10).unwrap()
        }
    }
    return sum;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_calc_sum() {
        assert_eq!(calc_sum("1122"), 3);
        assert_eq!(calc_sum("1111"), 4);
        assert_eq!(calc_sum("1234"), 0);
        assert_eq!(calc_sum("91212129"), 9);
    }

    #[test]
    fn test_calc_sum2() {
        assert_eq!(calc_sum2("1212"), 6);
        assert_eq!(calc_sum2("1221"), 0);
        assert_eq!(calc_sum2("123425"), 4);
        assert_eq!(calc_sum2("123123"), 12);
        assert_eq!(calc_sum2("12131415"), 4);
    }
}
