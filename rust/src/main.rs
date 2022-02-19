fn main() {
    let s = "1122";
    println!("sum: {}", calc_sum2(s))
}

fn calc_sum(s: &str) -> u32 {
    let chars: Vec<char> = s.chars().collect();
    let mut sum = 0;
    for (i, c) in chars.iter().enumerate() {
        let next: char = if i == chars.len() - 1 {
            chars[0]
        } else {
            chars[i + 1]
        };
        if *c == next {
            sum += c.to_digit(10).unwrap()
        }
    }
    return sum;
}

fn calc_sum2(s: &str) -> u32 {
    let chars: Vec<char> = s.chars().collect();
    let len = chars.len();
    let half = len / 2;
    let mut sum = 0;
    for (i, c) in chars.iter().enumerate() {
        let mut next_pos = i + half;
        if next_pos >= len {
            next_pos -= len;
        }
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
