use std::env;
use std::fs;
use std::io;

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);

    let day = args[1].parse::<u8>().unwrap();
    let part = args[2].parse::<u8>().unwrap();

    println!("Day: {}; part: {}", day, part);

    let input = read_input(2017, day).unwrap();
    let s = &input.trim();
    let output: u32;
    if part == 1 {
        output = day1::calc_sum(s);
    } else {
        output = day1::calc_sum2(s);
    }
    println!("Result: {}", output);
}

fn read_input(year: u16, day: u8) -> io::Result<String> {
    return fs::read_to_string(format!("../input/{}/{}.txt", year, day));
}

mod day1 {
    pub fn calc_sum(s: &str) -> u32 {
        return calc_sum_x(s, 1);
    }

    pub fn calc_sum2(s: &str) -> u32 {
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
}
