fn main() {
    let s = "1122";
    println!("sum: {}", calc_sum(s))
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(calc_sum("1122"), 3);
        assert_eq!(calc_sum("1111"), 4);
        assert_eq!(calc_sum("1234"), 0);
        assert_eq!(calc_sum("91212129"), 9);
    }
}
