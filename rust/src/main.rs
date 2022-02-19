fn main() {
    let s = "1122";
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
    println!("sum: {}", sum)
}
