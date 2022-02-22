use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", count_valid(s, are_equal)));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", count_valid(s, are_equal_or_anagrams)));
}

fn count_valid(s: &str, cmp: fn(&str, &str) -> bool) -> i32 {
    let lines = s.lines();
    let mut valid_count = 0;
    for line in lines {
        if is_valid(line, cmp) {
            valid_count += 1;
        }
    }
    return valid_count;
}

fn is_valid(s: &str, cmp: fn(&str, &str) -> bool) -> bool {
    let words: Vec<&str> = s.split(" ").collect();
    for i in 0..words.len() {
        let left = words[i];
        for j in i + 1..words.len() {
            let right = words[j];
            if cmp(left, right) {
                return false;
            }
        }
    }
    return true;
}

fn are_equal(left: &str, right: &str) -> bool {
    return left == right;
}

fn are_equal_or_anagrams(left: &str, right: &str) -> bool {
    return left == right || are_anagrams(left, right);
}

fn are_anagrams(left: &str, right: &str) -> bool {
    if left.len() != right.len() {
        return false;
    }
    for c in left.chars() {
        if left.matches(c).count() != right.matches(c).count() {
            return false;
        }
    }
    return true;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_are_anagrams() {
        assert_eq!(are_anagrams("abcde", "fghij"), false);
        assert_eq!(are_anagrams("abcde", "ecdab"), true);
        assert_eq!(are_anagrams("ab", "abc"), false);
        assert_eq!(are_anagrams("abc", "abd"), false);
        assert_eq!(are_anagrams("oiii", "ioii"), true);
        assert_eq!(are_anagrams("ooii", "oooi"), false);
    }

    #[test]
    fn test_is_valid() {
        assert_eq!(is_valid("aa bb cc dd ee", are_equal), true);
        assert_eq!(is_valid("aa bb cc dd aa", are_equal), false);
        assert_eq!(is_valid("aa bb cc dd aaa", are_equal), true);
        assert_eq!(is_valid("abcde fghij", are_equal_or_anagrams), true);
        assert_eq!(is_valid("abcde xyz ecdab", are_equal_or_anagrams), false);
        assert_eq!(
            is_valid("a ab abc abd abf abj", are_equal_or_anagrams),
            true
        );
        assert_eq!(
            is_valid("iiii oiii ooii oooi oooo", are_equal_or_anagrams),
            true
        );
        assert_eq!(
            is_valid("oiii ioii iioi iiio", are_equal_or_anagrams),
            false
        );
    }

    #[test]
    fn test_part1() {
        let input = "aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa";
        assert_eq!(part1(input).unwrap(), "2");
    }

    #[test]
    fn test_part2() {
        let input = "abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii ooii oooi oooo\noiii ioii iioi iiio";
        assert_eq!(part2(input).unwrap(), "3");
    }
}
